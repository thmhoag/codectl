package repo

import (
	"github.com/spf13/cobra"
	"github.com/thmhoag/codectl/pkg/clif"
)

type listOpts struct {
	formatString string
}

func newListCmd(ctx Ctx) *cobra.Command {
	opts := &listOpts{}

	log := ctx.Log()
	repman := ctx.RepoManager()

	cmd := &cobra.Command{
		Use:   "ls",
		Short: "List template repos",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {

			repos := repman.GetRepos()
			if err := clif.New(opts.formatString).
				Output(cmd.OutOrStdout()).
				Write(repos); err != nil {
				log.Fatal(err)
			}
		},
	}

	cmd.Flags().StringVar(&opts.formatString, "format", "table {{.Name}} {{.URL}}", "provide a template to format the list")

	return cmd
}