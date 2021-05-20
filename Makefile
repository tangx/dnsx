
PKG = $(shell cat go.mod | grep "^module " | sed -e "s/module //g")
VERSION = v$(shell cat .version)
COMMIT_SHA ?= $(shell git describe --always)

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
GOBUILD=CGO_ENABLED=0 go build -a -ldflags "-X ${PKG}/version.Version=${VERSION}+sha.${COMMIT_SHA}"

WORKSPACE ?= dnsx

clean:
	rm -rf ./cmd/$(WORKSPACE)/out

upgrade:
	go get -u ./...

build:
	$(MAKE) build.dnsx GOOS=linux GOARCH=amd64
	$(MAKE) build.dnsx GOOS=linux GOARCH=arm64

build.dnsx:
	@echo "Building dnsx for $(GOOS)/$(GOARCH)"
	cd ./cmd/$(WORKSPACE) && $(GOBUILD) -o ./out/dnsx-$(GOOS)-$(GOARCH) && tar -czf ./out/dnsx-$(GOOS)-$(GOARCH).tar.gz -C ./out/ dnsx-$(GOOS)-$(GOARCH)

install: build.dnsx
	mv ./cmd/$(WORKSPACE)/out/dnsx-$(GOOS)-$(GOARCH) ${GOPATH}/bin/dnsx

release:
	git push
	git push origin ${VERSION}

dnsx.config:
	COMMIT_SHA=${COMMIT_SHA} dnsx config

dnsx.buildx:
	COMMIT_SHA=${COMMIT_SHA} dnsx buildx --with-builder --push

dnsx.deploy:
	COMMIT_SHA=${COMMIT_SHA} dnsx deploy

apply: install dnsx.buildx dnsx.deploy

debug: install
	COMMIT_SHA=${COMMIT_SHA} dnsx run env | grep PROJECT_

