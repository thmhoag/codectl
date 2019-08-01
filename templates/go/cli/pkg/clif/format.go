package clif

import "strings"

// Format keys used to specify certain kinds of output formats
const (
	TableFormatKey  = "table"
	RawFormatKey    = "raw"
	PrettyFormatKey = "pretty"
)

// Format is the format string rendered using the Context
type Format string

// IsTable returns true if the format is a table-type format
func (f Format) IsTable() bool {
	return strings.HasPrefix(string(f), TableFormatKey)
}

// Contains returns true if the format contains the substring
func (f Format) Contains(sub string) bool {
	return strings.Contains(string(f), sub)
}

func (f Format) FinalFormat() string {

	finalFormat := strings.Trim(string(f), " ")
	if f.IsTable() {
		finalFormat = finalFormat[len(TableFormatKey):]
		finalFormat = strings.Trim(finalFormat, " ")
		r := strings.NewReplacer(" ", "\t", `\t`, "\t", `\n`, "\n")
		finalFormat = r.Replace(finalFormat)
	}

	return finalFormat
}
