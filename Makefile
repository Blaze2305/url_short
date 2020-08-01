OS = linux
GO_BUILD = go build
GOFLAGS = GOOS=${OS} CGO_ENABLED=0

build:
	${GOFLAGS} ${GO_BUILD} -tags netgo -a -v -installsuffix cgo -o bin/url_short cmd/main.go
.PHONY: go-build

vendor:
	go mod vendor
.PHONY: vendor

all: vendor build
.PHONY: all

default: all
.PHONY: default