package repo

import (
	"github.com/spf13/cobra"
	"github.com/thmhoag/codectl/pkg/repomanager"
)

func newAddCmd(ctx Ctx) *cobra.Command {

	log := ctx.Log()
	repman := ctx.RepoManager()

	cmd := &cobra.Command{
		Use:   "add [name] [url]",
		Short: "Add a new template repo",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			repo := &repomanager.Repository{
				Name: args[0],
				URL: args[1],
			}

			if exists := repman.Exists(repo); exists {
				cmd.Printf("Repo \"%s\" is already in your repositories\n", repo.Name)
				return
			}

			if err := repman.AddRepo(repo); err != nil {
				log.WithError(err).Errorln("unable to add repo")
			}
		},
	}

	return cmd
}