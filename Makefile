.PHONY: test clean all default

default: build

build:
	go build -o ~/.local/bin/daily
