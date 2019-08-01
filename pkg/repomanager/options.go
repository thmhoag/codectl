package repomanager

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type ManagerOpts struct {
	CacheDir string
	PropName string
	Config   *viper.Viper
}

func (o *ManagerOpts) Validate() error {

	if o.CacheDir == "" {
		return errors.New("must have a repository cache dir")
	}

	if o.PropName == "" {
		o.PropName = "repositories"
	}

	if o.Config == nil {
		return errors.New("config may not be nil")
	}

	return nil
}