package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"log"
)

var (
	showVersionFlag bool

	Version string = "0.0.0"
	Commit  string
	Date    string

	logLevelLookup = map[string]logrus.Level{
		"panic":   logrus.PanicLevel,
		"fatal":   logrus.FatalLevel,
		"error":   logrus.ErrorLevel,
		"warning": logrus.WarnLevel,
		"info":    logrus.InfoLevel,
		"debug":   logrus.DebugLevel,
		"trace":   logrus.TraceLevel,
	}
)

var rootCmd = &cobra.Command{
	Use: {{ .ProjectName | quote }},
	Short: "a short description of the app",
	Long: `a longer description of the app`,
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("fatal error:\n%s", err)
	}
}

func init() {

	rootCmd.Version = Version
	rootCmd.SetVersionTemplate("{{ .Version }}\n")
	rootCmd.Flags().BoolVarP(&showVersionFlag, "version", "v", false, "return the version of the executable")
	rootCmd.PersistentFlags().String("loglevel", "warning", "logrus log level [panic, fatal, error, warning, info, debug, trace]")
}