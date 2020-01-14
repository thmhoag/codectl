package clog

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

// Formatter - logrus formatter, implements logrus.Formatter
type Formatter struct {
	UseColors        bool     // show colors on log level
	UseMessageColors bool     // show message color as well as level
	ShowFullLevel    bool     // true to show full level [WARNING] instead [WARN]
	TrimMessages     bool     // true to trim whitespace on messages
	UseTitleLevel    bool     // always capitalize first letter of the level
}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	levelColor := getColorByLevel(entry.Level)

	b := &bytes.Buffer{}

	if f.UseColors {
		fmt.Fprintf(b, "\x1b[%dm", levelColor)
	}

	level := strings.ToLower(entry.Level.String())
	if !f.ShowFullLevel {
		level = level[:4]
	}

	if f.UseTitleLevel {
		level = strings.Title(level)
	}

	b.WriteString(level)
	b.WriteString(": ")

	if f.UseColors && !f.UseMessageColors {
		b.WriteString("\x1b[0m")
	}

	msg := entry.Message
	if f.TrimMessages {
		msg = strings.TrimSpace(msg)
	}

	b.WriteString(msg)
	b.WriteByte('\n')
	return b.Bytes(), nil
}

const (
	colorRed    = 31
	colorYellow = 33
	colorBlue   = 36
	colorGray   = 37
)

func getColorByLevel(level logrus.Level) int {
	switch level {
	case logrus.DebugLevel, logrus.TraceLevel:
		return colorGray
	case logrus.WarnLevel:
		return colorYellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		return colorRed
	default:
		return colorBlue
	}
}
