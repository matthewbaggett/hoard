VERSION ?= $(shell git describe --tags --dirty --always | sed -e 's/^v//')

all: build
.PHONY: format build
format:
	go fmt ./...
pre-build:
	echo $(VERSION) > pkg/common/VERSION
	date -Iseconds > pkg/common/BUILD_TIME
build: format pre-build

	go build -ldflags "-X DataPond.version=$(VERSION)" -o bin/datapond datapond.go
	go build -ldflags "-X DataPond.version=$(VERSION)" -o bin/datalake datalake.go
	go build  -o bin/log log.go
	chmod +x bin/datapond bin/datalake bin/log
	./bin/log

clear:
	clear;

build-data-pond:
	echo "Building Data Pond"
	go build -ldflags "-X DataPond.version=$(VERSION)" -o bin/datapond datapond.go
	chmod +x bin/datapond
data-pond: clear build-data-pond
	echo "Running Data Pond"
	./bin/datapond \
		--http-bind=0.0.0.0 --http-port=10963

build-data-lake:
	echo "Building Data Lake"
	go build -ldflags "-X DataPond.version=$(VERSION)" -o bin/datalake datalake.go
	chmod +x bin/datalake
data-lake: clear build-data-lake
	echo "Running Data Lake"
	./bin/datalake