VERSION ?= $(shell git describe --tags --dirty --always | sed -e 's/^v//')

all: build
.PHONY: format build
format:
	go fmt ./...
build: format
	go build -ldflags "-X DataPond.version=$(VERSION)" -o bin/datapond datapond.go
	go build -ldflags "-X DataPond.version=$(VERSION)" -o bin/datalake datalake.go
	go build  -o bin/log log.go
	chmod +x bin/datapond bin/datalake bin/log
	./bin/log