package config

import (
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

func newViewCmd(ctx Ctx) *cobra.Command {

	cfg := ctx.Config()
	log := ctx.Log()

	cmd := &cobra.Command{
		Use:   "view",
		Short: "Display the current configuration",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {

			configYaml, err := yaml.Marshal(cfg.AllSettings())
			if err != nil {
				log.Fatalf("Failed to convert the merge config to yaml\n%v\n", err)
			}

			cmd.Println(string(configYaml))
		},
	}

	return cmd
}