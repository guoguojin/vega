package cmd

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"

	"code.vegaprotocol.io/vega/cmd/vegawallet/commands/cli"
	"code.vegaprotocol.io/vega/cmd/vegawallet/commands/flags"
	"code.vegaprotocol.io/vega/cmd/vegawallet/commands/printer"
	vglog "code.vegaprotocol.io/vega/libs/zap"
	api "code.vegaprotocol.io/vega/protos/vega/api/v1"
	commandspb "code.vegaprotocol.io/vega/protos/vega/commands/v1"
	coreversion "code.vegaprotocol.io/vega/version"
	"code.vegaprotocol.io/vega/wallet/network"
	"code.vegaprotocol.io/vega/wallet/node"
	"code.vegaprotocol.io/vega/wallet/version"
	"github.com/golang/protobuf/proto"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	sendTxLong = cli.LongDesc(`
		Send a transaction to a Vega node via the gRPC API. The command can be sent to
		any node of a registered network or to a specific node address.

		The transaction should be base64-encoded.
	`)

	sendTxExample = cli.Examples(`
		# Send a transaction to a registered network
		{{.Software}} tx send --network NETWORK BASE64_TRANSACTION

		# Send a transaction to a specific Vega node address
		{{.Software}} tx send --node-address ADDRESS BASE64_TRANSACTION

		# Send a transaction with a log level set to debug
		{{.Software}} tx send --network NETWORK --level debug BASE64_TRANSACTION

		# Send a transaction with a maximum of 10 retries
		{{.Software}} tx send --network NETWORK --retries 10 BASE64_TRANSACTION

		# Send a transaction without verifying network version compatibility
		{{.Software}} tx send --network NETWORK --retries 10 BASE64_TRANSACTION --no-version-check
	`)
)

type SendTxHandler func(io.Writer, *SendTxFlags, *RootFlags, *SendTxRequest) error

func NewCmdTxSend(w io.Writer, rf *RootFlags) *cobra.Command {
	return BuildCmdTxSend(w, SendTx, rf)
}

func BuildCmdTxSend(w io.Writer, handler SendTxHandler, rf *RootFlags) *cobra.Command {
	f := &SendTxFlags{}

	cmd := &cobra.Command{
		Use:     "send",
		Short:   "Send a transaction to a Vega node",
		Long:    sendTxLong,
		Example: sendTxExample,
		RunE: func(_ *cobra.Command, args []string) error {
			if aLen := len(args); aLen == 0 {
				return flags.ArgMustBeSpecifiedError("transaction")
			} else if aLen > 1 {
				return flags.TooManyArgsError("transaction")
			}
			f.RawTx = args[0]

			req, err := f.Validate()
			if err != nil {
				return err
			}

			if err := handler(w, f, rf, req); err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&f.Network,
		"network", "n",
		"",
		"Network to which the command is sent",
	)
	cmd.Flags().StringVar(&f.NodeAddress,
		"node-address",
		"",
		"Vega node address to which the command is sent",
	)
	cmd.Flags().StringVar(&f.LogLevel,
		"level",
		zapcore.InfoLevel.String(),
		fmt.Sprintf("Set the log level: %v", SupportedLogLevels),
	)
	cmd.Flags().Uint64Var(&f.Retries,
		"retries",
		DefaultForwarderRetryCount,
		"Number of retries when contacting the Vega node",
	)
	cmd.Flags().BoolVar(&f.NoVersionCheck,
		"no-version-check",
		false,
		"Do not check for network version compatibility",
	)

	autoCompleteNetwork(cmd, rf.Home)
	autoCompleteLogLevel(cmd)
	return cmd
}

type SendTxFlags struct {
	Network        string
	NodeAddress    string
	Retries        uint64
	LogLevel       string
	RawTx          string
	NoVersionCheck bool
}

func (f *SendTxFlags) Validate() (*SendTxRequest, error) {
	req := &SendTxRequest{
		Retries: f.Retries,
	}

	if len(f.LogLevel) == 0 {
		return nil, flags.FlagMustBeSpecifiedError("level")
	}
	if err := ValidateLogLevel(f.LogLevel); err != nil {
		return nil, err
	}
	req.LogLevel = f.LogLevel

	if len(f.NodeAddress) == 0 && len(f.Network) == 0 {
		return nil, flags.OneOfFlagsMustBeSpecifiedError("network", "node-address")
	}
	if len(f.NodeAddress) != 0 && len(f.Network) != 0 {
		return nil, flags.FlagsMutuallyExclusiveError("network", "node-address")
	}
	req.NodeAddress = f.NodeAddress
	req.Network = f.Network

	if len(f.RawTx) == 0 {
		return nil, flags.ArgMustBeSpecifiedError("transaction")
	}
	decodedTx, err := base64.StdEncoding.DecodeString(f.RawTx)
	if err != nil {
		return nil, flags.MustBase64EncodedError("transaction")
	}
	tx := &commandspb.Transaction{}
	if err := proto.Unmarshal(decodedTx, tx); err != nil {
		return nil, fmt.Errorf("couldn't unmarshal transaction: %w", err)
	}
	req.Tx = tx

	return req, nil
}

type SendTxRequest struct {
	Network     string
	NodeAddress string
	Retries     uint64
	LogLevel    string
	Tx          *commandspb.Transaction
}

func SendTx(w io.Writer, f *SendTxFlags, rf *RootFlags, req *SendTxRequest) error {
	p := printer.NewInteractivePrinter(w)
	str := p.String()
	defer p.Print(str)

	if rf.Output == flags.InteractiveOutput && version.IsUnreleased() {
		str.CrossMark().DangerText("You are running an unreleased version of the Vega wallet (").DangerText(coreversion.Get()).DangerText(").").NextLine()
		str.Pad().DangerText("Use it at your own risk!").NextSection()
	}

	log, err := BuildLogger(rf.Output, req.LogLevel)
	if err != nil {
		return err
	}
	defer vglog.Sync(log)

	var hosts []string
	if len(req.Network) != 0 {
		hosts, err = getHostsFromNetwork(rf, req.Network)
		if err != nil {
			return err
		}
	} else {
		hosts = []string{req.NodeAddress}
	}

	if !f.NoVersionCheck {
		networkVersion, err := getNetworkVersion(hosts)
		if err != nil {
			return err
		}
		if networkVersion != coreversion.Get() {
			return fmt.Errorf("software is not compatible with this network: network is running version %s but wallet software has version %s", networkVersion, coreversion.Get())
		}
	}

	forwarder, err := node.NewForwarder(log.Named("forwarder"), network.GRPCConfig{
		Hosts:   hosts,
		Retries: req.Retries,
	})
	if err != nil {
		return fmt.Errorf("couldn't initialise the node forwarder: %w", err)
	}
	defer func() {
		if err = forwarder.Stop(); err != nil {
			log.Warn("Couldn't stop the forwarder", zap.Error(err))
		}
	}()

	if rf.Output == flags.InteractiveOutput {
		str.BlueArrow().InfoText("Logs").NextLine()
	}

	ctx, cancelFn := context.WithTimeout(context.Background(), ForwarderRequestTimeout)
	defer cancelFn()

	resp, err := forwarder.SendTx(ctx, req.Tx, api.SubmitTransactionRequest_TYPE_ASYNC, -1)
	if err != nil {
		log.Error("couldn't send transaction", zap.Error(err))
		return fmt.Errorf("couldn't send transaction: %w", err)
	}

	if !resp.Success {
		d, err := hex.DecodeString(resp.Data)
		if err != nil {
			log.Error("unable to decode resp error string")
		}
		log.Error("transaction failed", zap.String("err", string(d)), zap.Uint32("code", resp.Code))
		return fmt.Errorf("transaction failed: %s", resp.Data)
	}

	log.Info("transaction successfully sent", zap.String("hash", resp.TxHash))
	if rf.Output == flags.InteractiveOutput {
		str.NextLine().CheckMark().Text("Transaction sent: ").SuccessText(resp.TxHash).NextLine()
	}

	return nil
}
