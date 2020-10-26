#! /bin/bash
# Copyright 2019 Kuei-chun Chen. All rights reserved.
die() { echo "$*" 1>&2 ; exit 1; }
REPO=$(basename "$(dirname "$(pwd)")")/$(basename "$(pwd)")
VERSION="v$(cat version)-$(date "+%Y%m%d")"
LDFLAGS="-X main.version=$VERSION -X main.repo=$REPO"
[[ "$(which go)" = "" ]] && die "go command not found"

mkdir -p dist
env GOOS=darwin GOARCH=amd64 go build -ldflags "$LDFLAGS" -o dist/matlas atlas.go
./dist/matlas -version
