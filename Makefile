OUTPUT_DIR :=_output

# Constants
GOPATH := $(shell go env GOPATH)

# Version
revision:=$(shell git rev-parse --short HEAD)
build_time:=$(shell date +%D@%T)
version_stamp:=$(revision)-$(build_time)
# Set the linker flags so that the version will be included in the binaries:
import_path:=github.com/openshift-online/ocm-support-cli
ldflags:=-X $(import_path)/pkg/info.VersionStamp=$(version_stamp)

build: clean
	go build -o rosa-support -ldflags="$(ldflags)" . || exit 1

install: clean
	go build -o $(GOPATH)/bin/rosa-support -ldflags="$(ldflags)" ./cmd/rosa-support || exit 1

clean:
	rm -f rosa-support

cmds:
	for cmd in $$(ls cmd); do \
		go build -ldflags="$(ldflags)" "./cmd/$${cmd}" || exit 1; \
	done