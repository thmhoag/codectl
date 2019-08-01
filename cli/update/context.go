package update

import (
	"github.com/sirupsen/logrus"
	"github.com/thmhoag/codectl/cli/version"
)

type Ctx interface {
	CurrentVersion() *version.Properties
	Log() *logrus.Entry
}
