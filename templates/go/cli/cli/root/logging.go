package root

import (
	"github.com/sirupsen/logrus"
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
)