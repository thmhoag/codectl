package template

import (
	"github.com/sirupsen/logrus"
	"github.com/thmhoag/codectl/pkg/generator"
)

type Ctx interface {
	Log() *logrus.Entry
	Generator() generator.Generator
}
