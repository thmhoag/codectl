package config

import "github.com/spf13/cobra"

func NewConfigCmd(ctx Ctx) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "config [command]",
		Short: "Manage the config",
	}

	cmd.AddCommand(newViewCmd(ctx))
	cmd.AddCommand(newSetCmd(ctx))

	return cmd
}