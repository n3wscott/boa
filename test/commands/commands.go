package commands

import (
	"github.com/spf13/cobra"
	"tableflip.dev/boa"
	"tableflip.dev/boa/test/commands/options"
)

var (
	oo = &options.OutputOptions{}
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cli",
		Short: options.Wrap80("Interact via the command line."),
		RunE: func(cmd *cobra.Command, args []string) error {
			return boa.PromptNext(cmd, args)
		},
	}

	AddCommands(cmd)
	return cmd
}

func AddCommands(topLevel *cobra.Command) {
	addFoo(topLevel)
	addBar(topLevel)
}
