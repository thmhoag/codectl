package repo

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/thmhoag/codectl/pkg/repomanager"
)

type Ctx interface {
	RepoManager() repomanager.Manager
	Config() *viper.Viper
	Log() *logrus.Entry
}

