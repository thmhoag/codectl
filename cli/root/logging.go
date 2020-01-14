package root

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"github.com/thmhoag/codectl/pkg/clog"
)

var (
	logLevelLookup = map[string]logrus.Level{
		"panic":   logrus.PanicLevel,
		"fatal":   logrus.FatalLevel,
		"error":   logrus.ErrorLevel,
		"warning": logrus.WarnLevel,
		"info":    logrus.InfoLevel,
		"debug":   logrus.DebugLevel,
		"trace":   logrus.TraceLevel,
	}

	logFormatterLookup = map[string]logrus.Formatter{
		"clean": getCleanLogFormatter(),
		"detailed": getDetailedLogFormatter(),
	}
)

func getCleanLogFormatter() logrus.Formatter {
	return &clog.Formatter{
		UseColors:     true,
		ShowFullLevel: true,
		TrimMessages:  true,
	}
}

func getDetailedLogFormatter() logrus.Formatter {
	return &nested.Formatter{
		HideKeys:    true,
		FieldsOrder: []string{"cmd", "calledAs", "args", "flags"},
	}
}