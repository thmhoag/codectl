#!/usr/bin/env bash

SCRIPTDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
REPODIR="$(dirname "$SCRIPTDIR")"
cd "$REPODIR"

# always fail script if a cmd fails
set -e

# fail script if piped command fails
set -o pipefail

# Required commands
command -v go >/dev/null 2>&1 || { echo "go is required but not installed, aborting..." >&2; exit 1; }

echo "Installing..."
go install -ldflags="-s -w -X main.version=v0.0.0-dev -X main.commit=none -X main.date=none"

echo "Complete!"