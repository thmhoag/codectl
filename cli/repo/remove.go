package repo

import (
	"github.com/spf13/cobra"
	"github.com/thmhoag/codectl/pkg/repomanager"
)

func newRemoveCmd(ctx Ctx) *cobra.Command {

	log := ctx.Log()
	repman := ctx.RepoManager()

	cmd := &cobra.Command{
		Use:   "rm [name]",
		Short: "Remove a template repo",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			repo := &repomanager.Repository{
				Name: args[0],
			}

			if exists := repman.Exists(repo); !exists {
				cmd.Printf("Repo \"%s\" does not exist in your repositories\n", repo.Name)
				return
			}

			if err := repman.RemoveRepo(repo); err != nil {
				log.WithError(err).Errorln("unable to remove repo")
			}
		},
	}

	return cmd
}