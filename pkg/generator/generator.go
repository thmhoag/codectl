package generator

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/gobuffalo/packd"
	"github.com/gobuffalo/packr/v2"
	tmpl "github.com/thmhoag/codectl/pkg/template"
	"gopkg.in/yaml.v2"
)

type Generator interface {

	// GetTemplates gets a map of the templates currently
	// available to the generator.
	//
	// Returns a map of the template properties and an error object
	GetTemplates() (map[string]*tmpl.Properties, error)

	// OverridesPath sets the path to use for template files that can override
	// files in the template being generated.
	//
	// Returns the Generator so it can be chained.
	OverridesPath(string) Generator

	// DestinationPath sets the path that will be used for the output of
	// the generated template.
	//
	// Returns the Generator so it can be chained.
	DestinationPath(string) Generator

	// PathPrefix sets the prefix to use for the path to the
	// embedded template folder
	//
	// Returns the Generator so it can be chained.
	PathPrefix(string) Generator

	// CleanDestination tells the generator whether or not to wipe the destination
	// before generating the template. Default is false.
	//
	// Returns the Generator so it can be chained.
	CleanDestination(bool) Generator

	// Generate processes the template and override files and generates the result
	// to the defined destination path.
	Generate(interface{}, *tmpl.Overrides) error
}

type generator struct {
	overridesPath    string
	destinationPath  string
	pathPrefix       string
	cleanDestination bool
	templates        *packr.Box
}

func (g *generator) GetTemplates() (map[string]*tmpl.Properties, error) {
	return tmpl.GetAll(g.templates, "")
}

func (g *generator) OverridesPath(path string) Generator {
	if path != "" {
		g.overridesPath = path
	}
	return g
}

func (g *generator) DestinationPath(path string) Generator {
	if path != "" {
		g.destinationPath = path
	}
	return g
}

func (g *generator) PathPrefix(prefix string) Generator {
	g.pathPrefix = prefix
	return g
}

func (g *generator) CleanDestination(clean bool) Generator {
	g.cleanDestination = clean
	return g
}

func (g *generator) Generate(parmsObject interface{}, overrides *tmpl.Overrides) error {
	if err := g.validate(); err != nil {
		return err
	}

	if err := cleanDestination(g); err != nil {
		return err
	}

	err := g.templates.WalkPrefix(g.pathPrefix, func(s string, f packd.File) error {
		i, err := f.FileInfo()
		if err != nil {
			return err
		}

		newFileName, err := processPath(g, i.Name(), parmsObject, overrides)
		if err != nil {
			return err
		}

		if strings.ToLower(filepath.Base(newFileName)) == strings.ToLower(tmpl.DefinitionFileName) {
			// this is the definition file, don't copy it
			return nil
		}

		isTemplateFile := filepath.Ext(i.Name()) == ".tpl"
		if isTemplateFile {
			newFileName = strings.TrimSuffix(newFileName, ".tpl")
		} else if filepath.Ext(i.Name()) == ".nogen" {
			newFileName = strings.TrimSuffix(newFileName, ".nogen")
		}

		newFilePath := filepath.Join(g.destinationPath, newFileName)

		newDir, _ := filepath.Split(newFilePath)
		if err := os.MkdirAll(newDir, os.ModePerm); err != nil {
			return err
		}

		overrideFilePath := filepath.Join(g.overridesPath, newFileName)
		if fileExists(overrideFilePath) {
			return fileCopy(overrideFilePath, newFilePath)
		}

		newFile, err := os.Create(newFilePath)
		if err != nil {
			return err
		}

		defer newFile.Close()

		if isTemplateFile {
			templateFileContents := f.String()
			tmpl, err := template.New("").Funcs(sprig.TxtFuncMap()).Funcs(createFuncMap()).Parse(templateFileContents)
			if err != nil {
				return err
			}

			if err := tmpl.Execute(newFile, parmsObject); err != nil {
				return err
			}

		} else {

			if _, err := newFile.Write([]byte(f.String())); err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

func (g *generator) validate() error {
	if g.overridesPath == "" {
		return errors.New("must provide a path for override files")
	}

	if g.destinationPath == "" {
		return errors.New("must provide a destination path")
	}

	if g.templates == nil {
		return errors.New("no template found")
	}

	return nil
}

func cleanDestination(g *generator) error {
	if !g.cleanDestination {
		return nil
	}

	files, err := filepath.Glob(fmt.Sprintf("%s/*", g.destinationPath))
	if err != nil {
		return err
	}

	for _, f := range files {

		if err := os.RemoveAll(f); err != nil {
			return err
		}
	}

	return nil
}

func processPath(g *generator, fileName string, parmsObject interface{}, overrides *tmpl.Overrides) (string, error) {
	newFileName := strings.Replace(fileName, g.pathPrefix, "", 1)
	newFileName = applyOverrides(newFileName, overrides)

	tmpl, err := template.New("").Funcs(sprig.TxtFuncMap()).Funcs(createFuncMap()).Parse(newFileName)
	if err != nil {
		return "", err
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, parmsObject); err != nil {
		return "", err
	}

	newFileName = buf.String()

	if runtime.GOOS == "windows" {
		newFileName = strings.ReplaceAll(newFileName, "/", `\`)
	}

	return newFileName, nil
}

func createFuncMap() template.FuncMap {
	return template.FuncMap{
		"toYaml": toYaml,
	}
}

func toYaml(i interface{}) (string, error) {
	bytes, err := yaml.Marshal(i)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// fileExists reports whether the named file or directory exists.
func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func fileCopy(src, dst string) error {

	if src == dst {
		return nil
	}

	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	return err
}

func applyOverrides(path string, overrides *tmpl.Overrides) string {
	for key, val := range overrides.Paths {
		if strings.HasPrefix(path, key) {
			path = strings.Replace(path, key, val, 1)
		}
	}

	return path
}