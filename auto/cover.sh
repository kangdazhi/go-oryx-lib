#!/usr/bin/env bash

set -e
echo "mode: atomic" > coverage.txt

function coverage() {
    go test $0 -race -coverprofile=tmp.txt -covermode=atomic
    ret=$?; if [[ $ret -eq 0 ]]; then
        cat tmp.txt >> coverage.txt
        rm -f tmp.txt
    fi
}

coverage github.com/ossrs/go-oryx-lib/aac
coverage github.com/ossrs/go-oryx-lib/amf0
coverage github.com/ossrs/go-oryx-lib/asprocess
coverage github.com/ossrs/go-oryx-lib/avc
coverage github.com/ossrs/go-oryx-lib/flv
coverage github.com/ossrs/go-oryx-lib/http
coverage github.com/ossrs/go-oryx-lib/https
coverage github.com/ossrs/go-oryx-lib/json
coverage github.com/ossrs/go-oryx-lib/kxps
coverage github.com/ossrs/go-oryx-lib/logger
coverage github.com/ossrs/go-oryx-lib/options
coverage github.com/ossrs/go-oryx-lib/rtmp

