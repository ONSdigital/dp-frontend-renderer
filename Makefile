BINPATH ?= build

BUILD_TIME=$(shell date +%s)
GIT_COMMIT=$(shell git rev-parse HEAD)
VERSION ?= $(shell git tag --points-at HEAD | grep ^v | head -n 1)
LDFLAGS=-ldflags "-w -s -X 'main.Version=${VERSION}' -X 'main.BuildTime=$(BUILD_TIME)' -X 'main.GitCommit=$(GIT_COMMIT)'"

build: generate-prod
	go build $(LDFLAGS) -tags 'production' -o $(BINPATH)/dp-frontend-renderer
	cp taxonomy-redirects.yml $(BINPATH)

debug: generate-debug
	go build $(LDFLAGS) -race -tags 'debug' -o build/dp-frontend-renderer
	HUMAN_LOG=1 DEBUG=1 ./build/dp-frontend-renderer

generate-debug:
	# build the dev version
	cd assets; go run github.com/jteeuwen/go-bindata/go-bindata -debug -o data.go -pkg assets templates/... locales/...
	{ echo "// +build debug\n"; cat assets/data.go; } > assets/debug.go.new
	mv assets/debug.go.new assets/data.go

generate-prod:
	# build the production version
	cd assets; go run github.com/jteeuwen/go-bindata/go-bindata -o data.go -pkg assets templates/... locales/...
	{ echo "// +build production\n"; cat assets/data.go; } > assets/data.go.new
	mv assets/data.go.new assets/data.go

test: generate-prod
	go test -race -cover -tags 'production' ./...

.PHONY: build debug generate
