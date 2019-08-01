#!/usr/bin/env bash

SCRIPTDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
REPODIR="$(dirname "$SCRIPTDIR")"
cd "$REPODIR"

# always fail script if a cmd fails
set -eo pipefail

command -v docker >/dev/null 2>&1 || { echo "docker is required but not installed, aborting..." >&2; exit 1; }

[ -n "$GITHUB_TOKEN" ] || { echo "GITHUB_TOKEN is required and not set, aborting..." >&2; exit 1; }

docker build . -t goreleaser-build-sdk

docker run --rm --privileged \
  -v $PWD:$PWD \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -w $PWD \
  -e GITHUB_TOKEN="$GITHUB_TOKEN" \
  goreleaser-build-sdk release --rm-dist "$@"