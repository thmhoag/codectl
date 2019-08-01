package clif

import (
	"bytes"
	"github.com/pkg/errors"
	"io"
	"os"
	"reflect"
	"sort"
	"strings"
	"text/tabwriter"
	"text/template"
)

// Formatter sets up the information required to print the output as desired
// and allows it to be printed
type Formatter interface {
	// Output sets the output writer to be used and
	// returns the Formatter so it can be chained
	//
	// Defaults to os.stdout if non is specified
	Output(io.Writer) Formatter

	// Write prints the formatted output to the specified io.Writer and
	// returns any errors
	Write(out ...interface{}) error
}

type formatter struct {
	format      Format
	output      io.Writer
	finalFormat string
	buffer      *bytes.Buffer
}

func New(s string) Formatter {
	return &formatter{
		format: Format(s),
		output: os.Stdout,
	}
}

func (f *formatter) Output(w io.Writer) Formatter {
	f.output = w
	return f
}

func (f *formatter) Write(out ...interface{}) error {

	processedOut := processOut(out...)
	f.preFormat()

	tmpl, err := f.parseFormat()
	if err != nil {
		return err
	}

	return f.write(tmpl, processedOut...)
}

func processOut(out ...interface{}) []interface{} {
	var processedOut []interface{}

	for _, v := range out {

		t := reflect.ValueOf(v)

		isSlice := t.Kind() == reflect.Slice
		isArray := t.Kind() == reflect.Array
		isMap := t.Kind() == reflect.Map

		if isSlice || isArray || isMap {
			convertedSlice := interfaceToSlice(v)
			processedOut = append(processedOut, processOut(convertedSlice...)...)
			continue
		}

		processedOut = append(processedOut, v)
	}

	return processedOut
}

func (f *formatter) preFormat() {
	f.finalFormat = f.format.FinalFormat()
}

func (f *formatter) parseFormat() (*template.Template, error) {
	tmpl, err := template.New("").Parse(f.format.FinalFormat())
	if err != nil {
		return tmpl, errors.Errorf("template parsing error: %v\n", err)
	}

	return tmpl, err
}

func (f *formatter) write(tmpl *template.Template, out ...interface{}) error {

	w := f.output
	if f.format.IsTable() {
		t := tabwriter.NewWriter(f.output, 20, 1, 3, ' ', 0)
		defer t.Flush()

		w = t
		writeHeaders(w, tmpl)
	}

	var rawOutput []string
	for _, prop := range out {

		buf := bytes.Buffer{}
		if err := tmpl.Execute(&buf, prop); err != nil {
			return err
		}

		buf.Write([]byte("\n"))
		rawOutput = append(rawOutput, buf.String())
	}

	sort.Strings(rawOutput)
	for _,s := range rawOutput {

		w.Write([]byte(s))
	}

	return nil
}

func writeHeaders(w io.Writer, tmpl *template.Template) {
	for _, n := range tmpl.Root.Nodes {
		replacer := strings.NewReplacer("{", "", "}", "", ".", "")
		headerString := strings.ToUpper(replacer.Replace(n.String()))
		w.Write([]byte(headerString))
	}

	w.Write([]byte("\n"))
}
