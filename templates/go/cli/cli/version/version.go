package version

import (
	"fmt"
	"github.com/spf13/cobra"
	"reflect"
	"text/tabwriter"
)

func NewVersionCmd(ctx Ctx) *cobra.Command {

	version := ctx.CurrentVersion()

	cmd := &cobra.Command{
		Use:   "version",
		Short: "displays current version information",
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {

			t := tabwriter.NewWriter(cmd.OutOrStdout(), 10, 1, 3, ' ', 0)
			defer t.Flush()

			e := reflect.ValueOf(version).Elem()
			for i := 0; i < e.NumField(); i++ {
				fieldName := e.Type().Field(i).Name
				fmt.Fprintf(t, "%v:\t", fieldName)
				fieldValue := e.Field(i).Interface()
				fmt.Fprintf(t, "%v\n", fieldValue)
			}
		},
	}

	return cmd
}
