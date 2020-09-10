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
## hello 파일 만들기
cat <<EOF > hello
hello, world!
EOF
## 테스트 fstream을 실행하여 cat 명령으로 hello 파일 읽기
./fstream  cat hello
## hello 파일 지우기
rm hello