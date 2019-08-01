package template

import "github.com/spf13/cobra"

func NewTemplateCmd(ctx Ctx) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "template [command]",
		Short: "commands to manage templates",
	}

	cmd.AddCommand(newListCmd(ctx))

	return cmd
}
