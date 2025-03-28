package v1_test

import (
	"testing"
	"time"

	v1 "code.vegaprotocol.io/vega/wallet/service/v1"
	"code.vegaprotocol.io/vega/wallet/service/v1/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

type testAuth struct {
	v1.Auth
	ctrl *gomock.Controller
}

func getTestAuth(t *testing.T) *testAuth {
	t.Helper()
	rsaKeys, err := v1.GenerateRSAKeys()
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	store := mocks.NewMockRSAStore(ctrl)
	store.EXPECT().GetRsaKeys().Return(rsaKeys, nil)

	tokenExpiry := 10 * time.Hour
	a, err := v1.NewAuth(zap.NewNop(), store, tokenExpiry)
	if err != nil {
		t.Fatal(err)
	}

	return &testAuth{
		Auth: a,
		ctrl: ctrl,
	}
}

func TestAuth(t *testing.T) {
	t.Run("verify a valid token", testVerifyValidToken)
	t.Run("verify an invalid token fail", testVerifyInvalidToken)
	t.Run("revoke a valid token", testRevokeValidToken)
	t.Run("revoke an invalid token fail", testRevokeInvalidToken)
}

func testVerifyValidToken(t *testing.T) {
	t.Parallel()
	auth := getTestAuth(t)
	w := "jeremy"

	// get a new session
	tok, err := auth.NewSession(w)
	assert.NoError(t, err)
	assert.NotEmpty(t, tok)

	wallet2, err := auth.VerifyToken(tok)
	assert.NoError(t, err)
	assert.Equal(t, w, wallet2)
}

func testVerifyInvalidToken(t *testing.T) {
	t.Parallel()
	auth := getTestAuth(t)
	tok := "that's not a token"

	w, err := auth.VerifyToken(tok)
	assert.EqualError(t, err, "couldn't parse JWT token: token is malformed: token contains an invalid number of segments")
	assert.Empty(t, w)
}

func testRevokeValidToken(t *testing.T) {
	t.Parallel()
	auth := getTestAuth(t)
	walletName := "jeremy"

	// get a new session
	tok, err := auth.NewSession(walletName)
	assert.NoError(t, err)
	assert.NotEmpty(t, tok)

	wallet2, err := auth.VerifyToken(tok)
	assert.NoError(t, err)
	assert.Equal(t, walletName, wallet2)

	// now we made sure the token exists, let's revoke and re-verify it
	name, err := auth.Revoke(tok)
	assert.NoError(t, err)
	assert.Equal(t, walletName, name)

	w, err := auth.VerifyToken(tok)
	assert.ErrorIs(t, err, v1.ErrSessionNotFound)
	assert.Empty(t, w)
}

func testRevokeInvalidToken(t *testing.T) {
	t.Parallel()
	auth := getTestAuth(t)
	tok := "hehehe that's not a toekn"

	name, err := auth.Revoke(tok)
	assert.EqualError(t, err, "couldn't parse JWT token: token is malformed: token contains an invalid number of segments")
	assert.Empty(t, name)
}
