
PKG = $(shell cat go.mod | grep "^module " | sed -e "s/module //g")
VERSION = v$(shell cat .version)
COMMIT_SHA ?= $(shell git describe --always)

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
GOBUILD=CGO_ENABLED=0 go build -a -ldflags "-X ${PKG}/version.Version=${VERSION}+sha.${COMMIT_SHA}"

clean:
	rm -rf ./out

upgrade:
	go get -u ./...

buildx:
	$(MAKE) build.dnsx GOOS=linux GOARCH=amd64
	$(MAKE) build.dnsx GOOS=linux GOARCH=arm64
	$(MAKE) build.dnsx GOOS=dawrin GOARCH=amd64
	$(MAKE) build.dnsx GOOS=dawwin GOARCH=arm64

build.dnsx:
	@echo "Building dnsx for $(GOOS)/$(GOARCH)"
	$(GOBUILD) -o ./out/dnsx-$(GOOS)-$(GOARCH) 

install: build.dnsx
	mv ./out/dnsx-$(GOOS)-$(GOARCH) /usr/local/bin/dnsx

release:
	git push
	git push origin ${VERSION}
