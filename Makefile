BINARY_NAME=intelygenz_scraper

SRC_DIR=.
PKG_DIR=./pkg

GO_VERSION=$(shell go version)
GOPATH=$(shell go env GOPATH)
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)


GO_BUILD=go build
GO_TEST=go test
GO_COVER=go test -cover
GO_CLEAN=go clean

GO_MOD=go mod
GO_RUN=go run


.PHONY: clean

all: build

build:
	$(GO_BUILD) -o $(BINARY_NAME) -v $(SRC_DIR)/main.go

test:
	$(GO_TEST) ./...

coverage:
	$(GO_COVER) ./... -coverprofile=coverage.out


clean:
	$(GO_CLEAN)
	rm -f $(BINARY_NAME)

