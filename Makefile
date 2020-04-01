#!/usr/bin/make -f

test: fmt
	go test -timeout=1s -race -covermode=atomic -count=1 ./...

compile:
	go build ./...

fmt:
	go fmt ./...

build: test compile

.PHONY: test compile build
