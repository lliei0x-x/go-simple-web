#!/bin/sh

# project name
PROJECT_NAME=go-simple-web

# Operating System to run binary program
export GOOS=linux 

# go module
export GO111MODULE=on
export GOPROXY=https://athens.azurefd.net

BASE_DIR=$(cd `dirname $0`;pwd -P) # go-simple-web

# set project version
GIT_COMMIT_VERSION=`date -u +%Y%m%d.%H%M%S`
if git status >/dev/null 2>&1 ; then 
    GIT_COMMIT_VERSION=r`git rev-list HEAD | wc -l | awk '{print $1}'`
fi

# golang build args
GO_LDFLAGS="-s -w"
GO_LDFLAGS="$GO_LDFLAGS -X main.commitVersion=$GIT_COMMIT_VERSION"
# complie project binary program
cd "$BASE_DIR"
go build -ldflags "$GO_LDFLAGS" -o $PROJECT_NAME