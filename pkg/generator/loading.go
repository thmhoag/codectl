package generator

import (
	"github.com/gobuffalo/packr/v2"
	"os"
)

func LoadFromPackr(box *packr.Box) Generator {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return &generator{
		overridesPath: currentDir,
		destinationPath: currentDir,
		templates: box,
	}
}

func LoadFromPath(path string) Generator {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return &generator{
		overridesPath: currentDir,
		destinationPath: currentDir,
		templates: packr.New("templatesBox", path),
	}
}
