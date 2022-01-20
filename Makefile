VERSION := dev-$(shell git describe --tags $(shell git rev-list --tags --max-count=1))
COMMIT := $(shell git rev-parse --short HEAD)
BUILDFLAGS = -ldflags "-X github.com/linuxsuren/cobra-extension/version.version=$(VERSION) \
	-X github.com/linuxsuren/cobra-extension/version.commit=$(COMMIT) \
	-X github.com/linuxsuren/cobra-extension/version.date=$(shell date +'%Y-%m-%d') -w -s"

build:
	CGO_ENABLE=0 go build $(BUILDFLAGS) -o bin/cgit

build-linux:
	CGO_ENABLE=0 GOOS=linux go build $(BUILDFLAGS) -o bin/linux/cgit
	upx bin/linux/cgit

goreleaser:
	goreleaser release --rm-dist --snapshot

copy: build
	sudo cp bin/cgit /usr/local/bin/cgit

# Install golang-lint via https://golangci-lint.run/usage/install/#local-installation
# or via 'hd install golangci-lint'
lint:
	golangci-lint run ./...
