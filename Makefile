OUTPUT_DIR :=_output

# Constants
GOPATH := $(shell go env GOPATH)

# Version
revision:=$(shell git rev-parse --short HEAD)
build_time:=$(shell date +%D@%T)
version_stamp:=$(revision)-$(build_time)
# Set the linker flags so that the version will be included in the binaries:
import_path:=github.com/openshift-online/rosa-support
ldflags:=-X $(import_path)/pkg/info.VersionStamp=$(version_stamp)

.PHONY: build
build: clean
	go build -o rosa-support -ldflags="$(ldflags)" . || exit 1

.PHONY: install
install: clean
	go build -o $(GOPATH)/bin/rosa-support -ldflags="$(ldflags)" . || exit 1

.PHONY: clean
clean:
	rm -f rosa-support

.PHONY: test
test:
	go test ./...

.PHONY: fmt
fmt:
	gofmt -s -l -w cmd pkg

.PHONY: lint
lint:
	golangci-lint run --timeout 5m0s