#!/bin/bash
## set version `git describe --tags $(git rev-list --tags --max-count=1)`
VERSION=`git describe --tags $(git rev-list --tags --max-count=1)`
## set build `git log -1 --pretty=format:%h`
BUILD=`git log -1 --pretty=format:%h`

## go build
go build -ldflags "-X main.version=${VERSION}@${BUILD}" 

## test version
./fstream -version
## test function
./fstream cat filestream.go