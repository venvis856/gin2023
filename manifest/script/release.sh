#!/usr/bin/env bash

set -eu
set -o pipefail

echo -n "Build timestamp: "
TIMESTAMP=$(date +%s)
echo $TIMESTAMP

echo -n "Git commit: "
if [[ "$(git diff --stat)" != '' ]] || [[ -n "$(git status -s)" ]]; then
  echo "Changes detected, building a development version"
  COMMIT=$(git rev-parse HEAD)-dirty
  RELEASE=false
else
  echo "Building a release version"
  COMMIT=$(git rev-parse HEAD)
  RELEASE=true
fi
echo $COMMIT

echo -n "Tagged version: "
if git describe --tags --exact-match --match "v[0-9]*.[0-9]*.[0-9]*"; then
  VERSION=$(git describe --tags --exact-match --match "v[0-9]*.[0-9]*.[0-9]*")
  echo $VERSION
else
  VERSION=v0.0.0
  RELEASE=false
fi

echo Running "go $@"
exec go "$1" -ldflags \
  "-w -s -X gin/internal/app/version.buildTimestamp=$TIMESTAMP
   -X gin/internal/app/version.buildCommitHash=$COMMIT
   -X gin/internal/app/version.buildVersion=$VERSION
   -X gin/internal/app/version.buildRelease=$RELEASE" "${@:2}"
