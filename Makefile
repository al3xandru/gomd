LDFLAGS = -X main.buildSha=$(shell git rev-parse --short HEAD) -X main.buildDate=$(shell date +%Y%m%dT%H%M)

default: all

all: gomd

gomd: gomd_amd64 gomd_arm64
	@lipo -create -output $@ gomd_amd64 gomd_arm64
	@rm -vf gomd_amd64 gomd_arm64

gomd_arm64: gomd.go
	 @GOOS=darwin GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o $@ gomd.go

gomd_amd64: gomd.go
	@GOOS=darwin GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $@ gomd.go

.PHONY: install
install: gomd
	mv gomd $(HOME)/bin

.PHONY: clean
clean:
	@rm -vf gomd
	@rm -vf gomd_*
	@rm -vf output.html
