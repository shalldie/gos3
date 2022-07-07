#!/bin/bash

set -e

export CGO_ENABLED=0

TARGET_OS_NAMES=(linux-amd64 linux-arm64 darwin-amd64 darwin-arm64)

for os_name in ${TARGET_OS_NAMES[*]}; do

    tupleName=(${os_name//-/ })

    echo build $os_name ...

    GOOS=${tupleName[0]} \
        GOARCH=${tupleName[1]} \
        go build -o gos3.${os_name} cmd/main.go

    mkdir -p build
    mv gos3.$os_name build/gos3.$os_name

done
