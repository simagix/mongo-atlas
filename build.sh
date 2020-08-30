#! /bin/bash
# Copyright 2019 Kuei-chun Chen. All rights reserved.

DEP=`which dep`
if [ "$DEP" == "" ]; then
    echo "dep command not found"
    exit
fi

if [ -d vendor ]; then
    UPDATE="-update"
fi
export ver=$(cat version)
export version="v${ver}-$(date "+%Y%m%d")"
$DEP ensure $UPDATE
mkdir -p dist
env GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.version=$version" -o dist/matlas atlas.go
./dist/matlas -version
