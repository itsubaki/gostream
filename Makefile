SHELL := /bin/bash

test:
	go version
	go test -v -cover $(shell go list ./...)
