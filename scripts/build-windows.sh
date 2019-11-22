#!/usr/bin/env bash

SCRIPTDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
REPODIR="$(dirname "$SCRIPTDIR")"
cd "$REPODIR/cmd/codectl"

# always fail script if a cmd fails
set -eo pipefail

# Required commands
command -v go >/dev/null 2>&1 || { echo "go is required but not installed, aborting..." >&2; exit 1; }
command -v git >/dev/null 2>&1 || { echo "git is required but not installed, aborting..." >&2; exit 1; }

BUILD_TIME="$(date +'%m/%d/%Y.%H:%M')"
COMMIT_HASH="$(git rev-parse --short HEAD)"
GOLANG_VERSION=$(go version | awk '{print $3;}')

echo "Installing..."
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -X main.semver=v0.0.0-dev -X main.commit=$COMMIT_HASH -X main.built=$BUILD_TIME -X main.goversion=$GOLANG_VERSION" -o "$REPODIR/dist/codectl-windows-amd64/codectl.exe"

echo "Complete!"
