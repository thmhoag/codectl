#!/usr/bin/env bash

SCRIPTDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
REPODIR="$(dirname "$SCRIPTDIR")"
cd "$REPODIR/cmd/codectl"

# always fail script if a cmd fails
set -e

# fail script if piped command fails
set -o pipefail

# Required commands
command -v go >/dev/null 2>&1 || { echo "go is required but not installed, aborting..." >&2; exit 1; }
command -v git >/dev/null 2>&1 || { echo "git is required but not installed, aborting..." >&2; exit 1; }
command -v packr2 >/dev/null 2>&1 || { echo "packr2 is required but not installed, aborting..." >&2; exit 1; }

echo "Using packr2 $(packr2 version)"

BUILD_TIME="$(date +'%m/%d/%Y.%H:%M')"
COMMIT_HASH="$(git rev-parse --short HEAD)"

echo "Installing..."
packr2 install -ldflags="-s -w -X main.semver=v0.0.0-dev -X main.commit=$COMMIT_HASH -X main.built=$BUILD_TIME"

echo "Cleaning intermediary files..."
packr2 clean

echo "Complete!"