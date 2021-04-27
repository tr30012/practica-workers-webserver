.PHONY: build

build: 
    go build -v ./cmd/webserver

.DEFAULT_GOAL := build
 
