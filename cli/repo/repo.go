package repo

import "github.com/spf13/cobra"

func NewRepoCmd(ctx Ctx) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "repo [command]",
		Short: "Manage template repositories",
	}

	cmd.AddCommand(newAddCmd(ctx))
	cmd.AddCommand(newListCmd(ctx))
	cmd.AddCommand(newRemoveCmd(ctx))
	cmd.AddCommand(newUpdateCmd(ctx))

	return cmd
}
