package template

import (
	"encoding/json"
	"github.com/gobuffalo/packd"
	"github.com/gobuffalo/packr/v2"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	DefinitionFileName = ".codectl.yaml"
)

// GetAll returns a map of template path to their definitions
func GetAll(box *packr.Box, dir string) (map[string]*Properties, error) {

	templates := map[string]*Properties{}
	err := box.WalkPrefix(dir, func(s string, file packd.File) error {

		fileInfo, err := file.FileInfo()
		if err != nil {
			log.Debugf("error getting file info from packr box:\n%s", err)
			return err
		}

		log.Tracef("checking if %s is at the root of a template", fileInfo.Name())
		if !isInTemplateRoot(fileInfo) {
			log.Tracef("file %s was not the root of a template, skipping", fileInfo.Name())
			return nil
		}

		props, err := unmarshalProperties(file)
		if err != nil {
			return err
		}

		if props.Name == "" {
			props.Name = strings.ReplaceAll(filepath.Dir(fileInfo.Name()), "/", ".")
			props.Name = strings.ReplaceAll(props.Name, `\`, ".")
		}

		if props.Overrides.Paths != nil {
			props.Overrides.Paths = prepOverridePaths(props.Overrides.Paths)
		}

		templates[props.Name] = props
		return nil
	})

	return templates, err
}

func isInTemplateRoot(file os.FileInfo) bool {

	return strings.ToLower(filepath.Base(file.Name())) == strings.ToLower(DefinitionFileName)
}

func unmarshalProperties(file packd.File) (*Properties, error) {

	fileBytes := []byte(file.String())

	props := &Properties{}
	// Try and load yaml
	if err := yaml.Unmarshal(fileBytes, props); err != nil {
		return nil, err
	}

	parsedSuccessfully := props.Name != "" || props.Description != ""
	if parsedSuccessfully {
		// we're done, send it back
		return props, nil
	}

	// Try and load json
	if err := json.Unmarshal(fileBytes, props); err != nil {
		return nil, err
	}

	return props, nil
}

func prepOverridePaths(overrides map[string]string) map[string]string {
	newPaths := make(map[string]string)

	for k, v := range overrides {
		newKey := k

		if !strings.HasPrefix(newKey, "/") {
			newKey = "/" + newKey
		}

		if runtime.GOOS == "windows" {
			newKey = strings.ReplaceAll(newKey, "/", `\`)
		}

		newPaths[newKey] = v
	}

	return newPaths
}