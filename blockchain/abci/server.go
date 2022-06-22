// Copyright (c) 2022 Gobalsky Labs Limited
//
// Use of this software is governed by the Business Source License included
// in the LICENSE file and at https://www.mariadb.com/bsl11.
//
// Change Date: 18 months from the later of the date of the first publicly
// available Distribution of this version of the repository, and 25 June 2022.
//
// On the date above, in accordance with the Business Source License, use
// of this software will be governed by version 3 or later of the GNU General
// Public License.

package abci

import (
	"net"
	"strconv"

	"code.vegaprotocol.io/vega/blockchain"
	"code.vegaprotocol.io/vega/logging"

	"github.com/tendermint/tendermint/abci/server"
	"github.com/tendermint/tendermint/abci/types"
	tmlog "github.com/tendermint/tendermint/libs/log"
	"github.com/tendermint/tendermint/libs/service"
)

// Server is an abstraction over the abci server.
type Server struct {
	blockchain.Config
	log  *logging.Logger
	abci types.Application
	srv  service.Service
}

// NewServer instantiate a new server.
func NewServer(log *logging.Logger, config blockchain.Config, app types.Application) *Server {
	// setup logger
	log = log.Named("tm")
	log.SetLevel(config.Level.Get())

	return &Server{
		Config: config,
		log:    log,
		abci:   app,
		srv:    nil,
	}
}

// ReloadConf update the internal configuration.
func (s *Server) ReloadConf(cfg blockchain.Config) {
	s.log.Info("reloading configuration")
	if s.log.GetLevel() != cfg.Level.Get() {
		s.log.Info("updating log level",
			logging.String("old", s.log.GetLevel().String()),
			logging.String("new", cfg.Level.String()),
		)
		s.log.SetLevel(cfg.Level.Get())
	}

	// TODO(): not updating the the actual server for now, may need to look at this later
	// e.g restart the http server on another port or whatever
	s.Config = cfg
}

// Start configures and runs a new socket based ABCI tendermint blockchain
// server for the VEGA application.
func (s *Server) Start() error {
	addr := net.JoinHostPort(s.Tendermint.ServerAddr, strconv.Itoa(s.Tendermint.ServerPort))
	srv, err := server.NewServer(addr, "socket", s.abci)
	if err != nil {
		return err
	}
	srv.SetLogger(&abciLogger{s.log.Named("abci.socket-server")})

	s.log.Info("Starting abci-blockchain socket server",
		logging.String("addr", s.Tendermint.ServerAddr),
		logging.Int("port", s.Tendermint.ServerPort))

	if err := srv.Start(); err != nil {
		return err
	}

	s.srv = srv

	return nil
}

// Stop the abci server.
func (s *Server) Stop() {
	if s.srv != nil {
		s.log.Info("Stopping abci-blockchain socket server")
		if err := s.srv.Stop(); err != nil {
			s.log.Error("Failed to stop abci-blockchain socket server cleanly",
				logging.Error(err))
		}
	}
}

type abciLogger struct {
	*logging.Logger
}

func (l *abciLogger) Debug(msg string, keyvals ...interface{}) {
	l.Debugf("%v %v", msg, append([]interface{}{}, keyvals...))
}

func (l *abciLogger) Error(msg string, keyvals ...interface{}) {
	l.Errorf("%v %v", msg, append([]interface{}{}, keyvals...))
}

func (l *abciLogger) Info(msg string, keyvals ...interface{}) {
	l.Infof("%v %v", msg, append([]interface{}{}, keyvals...))
}

func (l *abciLogger) With(keyvals ...interface{}) tmlog.Logger {
	return l
}
