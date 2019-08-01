package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Ctx interface {
	Config() *viper.Viper
	Log() *logrus.Entry
}
