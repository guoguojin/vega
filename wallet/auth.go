package wallet

import (
	"crypto/rsa"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"

	"code.vegaprotocol.io/vega/logging"
	"github.com/dgrijalva/jwt-go/v4"
	"golang.org/x/crypto/sha3"
)

const (
	//  7 days, needs to be in seconds for the token
	tokenExpiry = time.Hour * 24 * 7
)

var (
	ErrSessionNotFound = errors.New("session not found")
)

type auth struct {
	log *logging.Logger
	// sessionID -> wallet naem
	sessions map[string]string
	privKey  *rsa.PrivateKey
	pubKey   *rsa.PublicKey

	mu sync.Mutex
}

func newAuth(log *logging.Logger, rootPath string) (*auth, error) {
	// get rsa keys
	pubBuf, privBuf, err := readRsaKeys(rootPath)
	if err != nil {
		return nil, err
	}
	priv, err := jwt.ParseRSAPrivateKeyFromPEM(privBuf)
	if err != nil {
		fmt.Printf("bad private key\n")
		return nil, err
	}
	pub, err := jwt.ParseRSAPublicKeyFromPEM(pubBuf)
	if err != nil {
		fmt.Printf("bad public key\n")
		return nil, err
	}

	return &auth{
		sessions: map[string]string{},
		privKey:  priv,
		pubKey:   pub,
		log:      log,
	}, nil
}

type Claims struct {
	jwt.StandardClaims
	Session string
	Wallet  string
}

func (a *auth) NewSession(walletname string) (string, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	session := genSession()
	// Create the Claims
	claims := &Claims{
		Session: session,
		Wallet:  walletname,
		StandardClaims: jwt.StandardClaims{
			// these are seconds
			ExpiresAt: jwt.NewTime((float64)(time.Now().Add(tokenExpiry).Unix())),
			Issuer:    "vega wallet",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	ss, err := token.SignedString(a.privKey)
	if err != nil {
		a.log.Error("unable to sign token", logging.Error(err))
		return "", err
	}

	// all good up to now, insert the new session
	a.sessions[session] = walletname
	return ss, nil
}

// returns the walletname associated for this session
func (a *auth) VerifyToken(token string) (string, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	// first parse the token
	claims, err := a.parseToken(token)
	if err != nil {
		return "", err
	}

	wallet, ok := a.sessions[claims.Session]
	if !ok {
		return "", ErrSessionNotFound
	}

	return wallet, nil
}

func (a *auth) Revoke(token string) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	claims, err := a.parseToken(token)
	if err != nil {
		return err
	}

	// extract session from the token
	_, ok := a.sessions[claims.Session]
	if !ok {
		return ErrSessionNotFound
	}
	delete(a.sessions, claims.Session)
	return nil
}

func (a *auth) parseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return a.pubKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, err

}

// ExtractToken this is public for testing purposes
func ExtractToken(f func(string, http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if len(token) <= 0 {
			// invalid token, return an error here
			writeError(w, ErrInvalidOrMissingToken, http.StatusBadRequest)
			return
		}
		splitToken := strings.Split(token, "Bearer")
		if len(splitToken) != 2 || len(splitToken[1]) <= 0 {
			// invalid token, return an error here
			writeError(w, ErrInvalidOrMissingToken, http.StatusBadRequest)
			return
		}
		// then call the function
		f(strings.TrimSpace(splitToken[1]), w, r)
	}
}

func genSession() string {
	hasher := sha3.New256()
	hasher.Write([]byte(randSeq(10)))
	return hex.EncodeToString(hasher.Sum(nil))
}

var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}
