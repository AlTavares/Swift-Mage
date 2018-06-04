#!/bin/sh

go get -u -d github.com/magefile/mage
cd $GOPATH/src/github.com/magefile/mage && go run bootstrap.go