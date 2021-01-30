package commands

import (
	"context"
	"errors"
	"strings"

	"github.com/spf13/cobra"
	"tableflip.dev/boa/test/bar"
	"tableflip.dev/boa/test/commands/options"
)

func addBar(topLevel *cobra.Command) {
	bo := &options.BarOptions{}

	cmd := &cobra.Command{
		Use:   "bar",
		Short: "Say hello!",
		Example: `
cli-base bar hello --name=example
`,
		ValidArgs: []string{},
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires a greeting")
			}
			bo.Message = strings.Join(args, " ")

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			s := bar.Bar{
				Message: bo.Message,
				Name:    bo.Name,
			}
			if oo.JSON {
				s.Output = "json"
			}
			err := s.Do(context.Background())
			return oo.HandleError(err)
		},
	}

	options.AddBarArgs(cmd, bo)
	options.AddOutputArg(cmd, oo)
	topLevel.AddCommand(cmd)
}
