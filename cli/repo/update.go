package repo

import (
	"github.com/spf13/cobra"
)

func newUpdateCmd(ctx Ctx) *cobra.Command {

	log := ctx.Log()
	repman := ctx.RepoManager()

	cmd := &cobra.Command{
		Use:   "update [name] [url]",
		Short: "Update cached templates from remote repositories",
		Args:  cobra.RangeArgs(0, 2),
		Run: func(cmd *cobra.Command, args []string) {

			if err := repman.UpdateAll(); err != nil {
				log.WithError(err).Fatalln("unable to update repositories")
			}
		},
	}

	return cmd
}