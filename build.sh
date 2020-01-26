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

$DEP ensure $UPDATE

export ver="0.2.0"
export version="v${ver}-$(date "+%Y%m%d")"
# env GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$version" -o matlas-linux-x64 atlas.go
env GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.version=$version" -o matlas atlas.go
