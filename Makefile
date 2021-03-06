NAME=wait-for-weave
HARDWARE=$(shell uname -m)
VERSION=0.0.2

deps:
	go get

build:
	#go build -o stage/wait-for-weave wait-for-weave.go
	CGO_ENABLED=0 go build -a -installsuffix cgo -o stage/wait-for-weave wait-for-weave.go

image: build
	docker build -t binocarlos/wait-for-weave .

.PHONY: build image