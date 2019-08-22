package template

import (
	"github.com/spf13/cobra"
	"github.com/thmhoag/clif"
)

type listOpts struct {
	formatString string
}

func newListCmd(ctx Ctx) *cobra.Command {
	opts := &listOpts{}

	gen := ctx.Generator()
	log := ctx.Log()

	cmd := &cobra.Command{
		Use:   "ls",
		Short: "list available templates",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {

			log.Trace("getting all available templates")
			templates, err := gen.GetTemplates()
			if err != nil {
				log.Fatal(err)
			}

			log.Tracef("got a total of %d templates", len(templates))

			log.Debugf("format string received: %s", opts.formatString)
			if err := clif.New(opts.formatString).
				Output(cmd.OutOrStdout()).
				Write(templates); err != nil {
				log.Fatal(err)
			}
		},
	}

	cmd.Flags().StringVar(&opts.formatString, "format", "table {{.Name}} {{.Description}}", "provide a template to format the list")

	return cmd
}
