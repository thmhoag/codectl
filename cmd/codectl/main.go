package main

import (
	"os"

	"github.com/thmhoag/codectl/cli"
)

var (
	semver, commit, built string
)

func main() {

	cli.Semver = semver
	cli.Commit = commit
	cli.Built = built

	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}
