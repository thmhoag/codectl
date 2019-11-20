package generate

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/thmhoag/codectl/pkg/template"
	"runtime"
	"strings"
	"github.com/AlecAivazis/survey/v2"
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

			// TODO: Find more elegant way to handle this dirty hack
			if runtime.GOOS == "windows" {
				templatePath = strings.ReplaceAll(templateName, ".", `\`)
			}

			gen.OverridesPath(opts.overridesDirFlag).
				DestinationPath(opts.outputDirFlag).
				PathPrefix(templatePath).
				CleanDestination(opts.cleanFlag)

			templates, err := gen.GetTemplates()
			if err != nil {
				log.Fatal(err)
			}

			templateProps := templates[templateName]
			if templateProps == nil {
				fmt.Println("Invalid template name:", templateName)
				return
			}

			parameterValues := make(map[string]interface{})
			if opts.valuesFilePath != "" {
				parameterValues, err = template.LoadValuesFromFile(opts.valuesFilePath)
				if err != nil {
					log.WithError(err).Fatalf("unable to process values file at %s", opts.valuesFilePath)
				}
			}

			templateHasParms := templateProps.Parameters != nil && len(templateProps.Parameters) > 0
			parmsAlreadyProvided := parameterValues != nil && len(parameterValues) > 0
			if templateHasParms && !parmsAlreadyProvided {

				for _, p := range templateProps.Parameters {
					parm := p
					parameterValues[parm.Name], err = promptForParameter(&parm)
					if err != nil {
						 if err.Error() == "interrupt" {
							 return
						 }

						log.Fatal(err)
					}
				}
			}

			if err := gen.Generate(parameterValues); err != nil {
				log.WithError(err).Fatal("error processing template")
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

	prompt := &survey.Input{
		Renderer: survey.Renderer{},
		Message:  fmt.Sprintf("%s:", parm.Prompt),
		Default: parm.Value,
	}

	var response string
	surveyOpts := []survey.AskOpt{
		survey.WithIcons(func(icons *survey.IconSet){
			icons.Question.Text = ""
			icons.Question.Format = "default"
			icons.Error.Format = "red+hb"
		}),
	}

	if parm.Required {
		surveyOpts = append(surveyOpts, survey.WithValidator(survey.Required))
	}

	err := survey.AskOne(prompt, &response, surveyOpts...)

	return response, err
}
