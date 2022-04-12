package genesis

import (
	"context"

	"code.vegaprotocol.io/vega/config"
	"github.com/jessevdk/go-flags"
)

type Cmd struct {
	// Global options
	config.VegaHomeFlag
	config.PassphraseFlag

	// Subcommands
	Generate       generateCmd       `command:"generate" description:"Generates the genesis file"`
	Update         updateCmd         `command:"update" description:"Update the genesis file with the app_state, useful if the genesis generation is not done using \"vega genesis generate\""`
	LoadCheckpoint loadCheckpointCmd `command:"load_checkpoint" description:"Load the given checkpoint file in the genesis file"`
}

var genesisCmd Cmd

func Genesis(ctx context.Context, parser *flags.Parser) error {
	genesisCmd = Cmd{
		Generate: generateCmd{
			TmHome: "$HOME/.tendermint",
		},
		Update: updateCmd{
			TmHome: "$HOME/.tendermint",
		},
		LoadCheckpoint: loadCheckpointCmd{
			TmHome: "$HOME/.tendermint",
		},
	}

	desc := "Manage the genesis file"
	cmd, err := parser.AddCommand("genesis", desc, desc, &genesisCmd)
	if err != nil {
		return err
	}
	return initNewCmd(ctx, cmd)
}
