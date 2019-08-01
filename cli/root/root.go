package root

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"strings"
	nested "github.com/antonfisher/nested-logrus-formatter"
)

type rootOpts struct {
	loglevel    string
	showVersion bool
}

func NewRootCmd(ctx Ctx) *cobra.Command {
	opts := &rootOpts{}

	cfg := ctx.Config()

	cmd := &cobra.Command{
		Use:   "codectl [command]",
		Short: "a template generator",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			ctx.ConfigureLogger(setupLogger(cmd, cfg))
			ctx.EditLogger(setupTelemetry(ctx, cmd))

			ctx.Log().Tracef("%s beginning execution", cmd.Name())
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			ctx.Log().Tracef("%s ending execution", cmd.Name())
		},
	}

	cmd.SetVersionTemplate("{{.Version}}\n")
	cmd.Flags().BoolVarP(&opts.showVersion, "version", "v", false, "return the version of the executable")
	cmd.PersistentFlags().StringVar(&opts.loglevel, "loglevel", "warning", "logrus log level [panic, fatal, error, warning, info, debug, trace]")
	cfg.BindPFlags(cmd.PersistentFlags())

	return cmd
}

func setupLogger(cmd *cobra.Command, config *viper.Viper) func (*logrus.Logger) *logrus.Logger {
	return func (log *logrus.Logger) *logrus.Logger {
		log.SetOutput(cmd.OutOrStderr())
		log.SetFormatter(&nested.Formatter{
			HideKeys:    true,
			FieldsOrder: []string{"cmd", "calledAs", "args", "flags"},
		})
		if level, ok := logLevelLookup[strings.ToLower(config.GetString("loglevel"))]; ok {
			log.SetLevel(level)
			log.Tracef("log level set to %s", level)
		}

		return log
	}
}

func setupTelemetry(ctx Ctx, cmd *cobra.Command) func (*logrus.Entry) *logrus.Entry {

	logflags := make(map[string]string)
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		logflags[f.Name] = f.Value.String()
	})

	return func(e *logrus.Entry) *logrus.Entry {
		return e.WithFields(logrus.Fields{
			"cmd":        cmd.CommandPath(),
			"calledAs":   cmd.CalledAs(),
			"args":       cmd.Args,
			"flags":      logflags,
			"workingDir": ctx.WorkingDir(),
		})
	}
}