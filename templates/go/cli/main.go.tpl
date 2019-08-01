package main

import {{ (printf "%s/cmd" .ModuleName) | quote }}

var (
	version, commit, date string
)

func main() {

	cmd.Version = version
	cmd.Commit = commit
	cmd.Date = date

	cmd.Execute()
}