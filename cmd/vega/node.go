package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"code.vegaprotocol.io/vega/internal"
	"code.vegaprotocol.io/vega/internal/api/endpoints/gql"
	"code.vegaprotocol.io/vega/internal/api/endpoints/grpc"
	"code.vegaprotocol.io/vega/internal/api/endpoints/restproxy"
	"code.vegaprotocol.io/vega/internal/blockchain"
	"code.vegaprotocol.io/vega/internal/execution"
	"code.vegaprotocol.io/vega/internal/fsutil"
	"code.vegaprotocol.io/vega/internal/logging"
	"code.vegaprotocol.io/vega/internal/monitoring"
	"code.vegaprotocol.io/vega/internal/pprof"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NodeCommand use to implement 'node' command.
type NodeCommand struct {
	command

	configPath string
	withPPROF  bool
	Log        *logging.Logger
}

// Init initialises the node command.
func (l *NodeCommand) Init(c *Cli) {
	l.cli = c
	l.cmd = &cobra.Command{
		Use:   "node",
		Short: "Run a new Vega node",
		Long:  "Run a new Vega node as defined by config files",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return l.runNode(args)
		},
		Example: nodeExample(),
	}
	l.addFlags()
}

// addFlags adds flags for specific command.
func (l *NodeCommand) addFlags() {
	flagSet := l.cmd.Flags()
	flagSet.StringVarP(&l.configPath, "config", "C", "", "file path to search for vega config file(s)")
	flagSet.BoolVarP(&l.withPPROF, "with-pprof", "", false, "start the node with pprof support")
}

// runNode is the entry of node command.
func (l *NodeCommand) runNode(args []string) error {

	// context used in waitSig to exit from the application
	// not from the user inputs
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Use configPath from args
	configPath := l.configPath
	if configPath == "" {
		// Use configPath from ENV
		configPath = envConfigPath()
		if configPath == "" {
			// Default directory ($HOME/.vega)
			configPath = fsutil.DefaultVegaDir()
		}
	}

	l.Log.Info("Config path", logging.String("config-path", configPath))

	// VEGA config (holds all package level configs)
	conf, err := internal.NewConfigFromFile(l.Log, configPath)
	if err != nil {
		// We revert to default configs if there are any errors in read/parse process
		l.Log.Error("Error reading config from file, using defaults", logging.Error(err))
		defaultConf, err := internal.NewDefaultConfig(l.Log, fsutil.DefaultVegaDir())
		if err != nil {
			return err
		}
		conf = defaultConf
	} else {
		conf.ListenForChanges()
	}

	var pproffhandlr *pprof.Pprofhandler
	if flagProvided("--with-pprof") || conf.Pprof.Enabled {
		l.Log.Info("vega is starting with pprof profile, this is not a recommended setting for production")
		var err error
		pproffhandlr, err = pprof.New(conf.Pprof)
		if err != nil {
			return err
		}
		defer pproffhandlr.Stop()
	}

	resolver, err := internal.NewResolver(conf, cancel)

	// Statistics provider
	stats := internal.NewStats(l.Log, l.cli.version, l.cli.versionHash)

	// Resolve services for injection to servers/execution engine
	orderService, err := resolver.ResolveOrderService()
	if err != nil {
		return err
	}
	tradeService, err := resolver.ResolveTradeService()
	if err != nil {
		return err
	}
	candleService, err := resolver.ResolveCandleService()
	if err != nil {
		return err
	}
	marketService, err := resolver.ResolveMarketService()
	if err != nil {
		return err
	}
	partyService, err := resolver.ResolvePartyService()
	if err != nil {
		return err
	}
	timeService, err := resolver.ResolveTimeService()
	if err != nil {
		return err
	}
	orderStore, err := resolver.ResolveOrderStore()
	if err != nil {
		return err
	}
	tradeStore, err := resolver.ResolveTradeStore()
	if err != nil {
		return err
	}
	candleStore, err := resolver.ResolveCandleStore()
	if err != nil {
		return err
	}
	marketStore, err := resolver.ResolveMarketStore()
	if err != nil {
		return err
	}
	partyStore, err := resolver.ResolvePartyStore()
	if err != nil {
		return err
	}

	client, err := resolver.ResolveBlockchainClient()
	if err != nil {
		return err
	}

	// Execution engine (broker operation at runtime etc)
	executionEngine := execution.NewEngine(
		conf.Execution,
		timeService,
		orderStore,
		tradeStore,
		candleStore,
		marketStore,
		partyStore,
	)

	// ABCI<>blockchain server
	bcService := blockchain.NewService(conf.Blockchain, stats.Blockchain, executionEngine, timeService)
	bcProcessor := blockchain.NewProcessor(conf.Blockchain, bcService)
	bcApp := blockchain.NewApplication(conf.Blockchain, stats.Blockchain, bcProcessor, bcService, timeService, cancel)
	socketServer := blockchain.NewServer(conf.Blockchain, stats.Blockchain, bcApp)
	err = socketServer.Start()
	if err != nil {
		return errors.Wrap(err, "ABCI socket server error")
	}

	statusChecker := monitoring.NewStatusChecker(l.Log, client, 500*time.Millisecond)
	statusChecker.OnChainDisconnect(cancel)

	// gRPC server
	grpcServer := grpc.NewGRPCServer(
		conf.API,
		stats,
		client,
		timeService,
		marketService,
		partyService,
		orderService,
		tradeService,
		candleService,
		statusChecker,
	)
	go grpcServer.Start()

	// REST<>gRPC (gRPC proxy) server
	restServer := restproxy.NewRestProxyServer(conf.API)
	go restServer.Start()

	// GraphQL server
	graphServer := gql.NewGraphQLServer(
		conf.API,
		orderService,
		tradeService,
		candleService,
		marketService,
		partyService,
		timeService,
		statusChecker,
	)
	go graphServer.Start()

	waitSig(ctx, l.Log)

	// Clean up and close resources
	l.Log.Info("Closing REST proxy server", logging.Error(restServer.Stop()))
	l.Log.Info("Closing GRPC server", logging.Error(grpcServer.Stop()))
	l.Log.Info("Closing GraphQL server", logging.Error(graphServer.Stop()))
	l.Log.Info("Closing blockchain server", logging.Error(socketServer.Stop()))
	l.Log.Info("Closing stores", logging.Error(resolver.CloseStores()))
	statusChecker.Stop()

	return nil
}

// nodeExample shows examples for node command, and is used in auto-generated cli docs.
func nodeExample() string {
	return `$ vega node
VEGA started successfully`
}

// envConfigPath attempts to look at ENV variable VEGA_CONFIG for the config.toml path
func envConfigPath() string {
	err := viper.BindEnv("config")
	if err == nil {
		return viper.GetString("config")
	}
	return ""
}

// waitSig will wait for a sigterm or sigint interrupt.
func waitSig(ctx context.Context, log *logging.Logger) {
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	select {
	case sig := <-gracefulStop:
		log.Info("caught signal", logging.String("name", fmt.Sprintf("%+v", sig)))
	case <-ctx.Done():
		// nothing to do
	}
}

func flagProvided(flag string) bool {
	for _, v := range os.Args[1:] {
		if v == flag {
			return true
		}
	}

	return false
}
