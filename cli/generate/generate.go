package generate

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/thmhoag/codectl/pkg/template"
	"strings"
)

type generateOpts struct {
	valuesFilePath   string
	overridesDirFlag string
	outputDirFlag    string
	cleanFlag        bool
}

func NewGenerateCmd(ctx Ctx) *cobra.Command {
	opts := &generateOpts{}

	log := ctx.Log()
	gen := ctx.Generator()

	cmd := &cobra.Command{
		Use:   "generate [template]",
		Aliases: []string{"g", "gen"},
		Short: "generate a template",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			templateName := args[0]
			templatePath := strings.ReplaceAll(templateName, ".", "/")

			gen.OverridesPath(opts.valuesFilePath).
				DestinationPath(opts.outputDirFlag).
				PathPrefix(templatePath).
				CleanDestination(opts.cleanFlag)

			templates, err := gen.GetTemplates()
			if err != nil {
				log.Fatal(err)
			}

			parameterValues := make(map[string]interface{})
			templateProps := templates[templateName]
			if templateProps.Parameters != nil && len(templateProps.Parameters) > 0 {

				for _, p := range templateProps.Parameters {
					parm := p
					parameterValues[parm.Name], err = promptForParameter(&parm)
					if err != nil {
						log.Fatal(err)
					}
				}
			}

			if err := gen.Generate(parameterValues); err != nil {
				log.Fatal(err)
			}
		},
	}

	cmd.PersistentFlags().StringVarP(&opts.valuesFilePath, "file", "f", "", "path to values file to use for the template")
	cmd.PersistentFlags().StringVar(&opts.overridesDirFlag, "overrides", "", "directory that will contain template file overrides")
	cmd.PersistentFlags().StringVarP(&opts.outputDirFlag, "output", "o", "", "directory to place the generated files")
	cmd.PersistentFlags().BoolVarP(&opts.cleanFlag, "clean", "c", false, "clean the destination prior to generating the files")

	return cmd
}

func promptForParameter(parm *template.Parameter) (string, error) {

	bold := promptui.Styler(promptui.FGBold)
	templates := &promptui.PromptTemplates{
		Prompt:  fmt.Sprintf("%s {{ .Prompt | bold }}%s ", bold(promptui.IconInitial), bold(":")),
		Valid:   fmt.Sprintf("%s {{ .Prompt | bold }}%s ", bold(promptui.IconGood), bold(":")),
		Invalid: fmt.Sprintf("%s {{ .Prompt | bold }}%s ", bold(promptui.IconBad), bold(":")),
		Success: "{{ printf \"%s:\" .Prompt | faint }} ",
	}

	prompt := promptui.Prompt{
		Label: parm,
		Templates: templates,
		Default: parm.Value,
		AllowEdit: true,
		Validate: func(s string) error {
			if !parm.Required {
				return nil
			}

			if s == "" {
				return errors.Errorf("parameter %s is required and was not provided", parm.Name)
			}

			return nil
		},
	}

	return prompt.Run()
}
