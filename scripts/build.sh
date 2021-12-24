#!/bin/bash

set -e

export CGO_ENABLED=0

TARGET_OS_NAMES=(linux-amd64 linux-arm64 darwin-amd64 darwin-arm64)

for os_name in ${TARGET_OS_NAMES[*]}; do

    tupleName=(${os_name//-/ })

    echo build $os_name ...

    GOOS=${tupleName[0]} \
        GOARCH=${tupleName[1]} \
        go build -o go-cli.${os_name}

    mkdir -p build
    mv go-cli.$os_name build/go-cli.$os_name

done
