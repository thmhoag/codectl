package config

import (
	"github.com/spf13/cobra"
)

func newSetCmd(ctx Ctx) *cobra.Command {

	cfg := ctx.Config()
	log := ctx.Log()

	cmd := &cobra.Command{
		Use:   "set [key] [value]",
		Short: "Set a configuration value",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {

			configKey := args[0]
			configVal := args[1]

			cfg.Set(configKey, configVal)
			if err := cfg.WriteConfig(); err != nil {
				log.WithError(err).Fatal("error writing configuration file")
			}
		},
	}

	return cmd
}