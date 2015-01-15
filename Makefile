NAME=wait-for-weave
HARDWARE=$(shell uname -m)
VERSION=0.0.1

build:
	go build -o stage/wait-for-weave wait-for-weave.go

image: build
	docker build -t binocarlos/wait-for-weave .

.PHONY: build image