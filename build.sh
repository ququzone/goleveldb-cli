#!/bin/sh

# build in docker
docker run --rm -it -v "$PWD":/go/src/github.com/ququzone/goleveldb-cli -w /go/src/github.com/ququzone/goleveldb-cli golang:1.10 go build .
