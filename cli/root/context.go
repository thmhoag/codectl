package root

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Ctx interface {
	Config() *viper.Viper
	ConfigureLogger(setup func (l *logrus.Logger) *logrus.Logger)
	EditLogger(func (entry *logrus.Entry) *logrus.Entry)
	Log() *logrus.Entry
	WorkingDir() string
}
