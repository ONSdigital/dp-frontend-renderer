BINPATH ?= build

build: generate
	go build -tags 'production' -o $(BINPATH)/dp-frontend-renderer
	cp taxonomy-redirects.yml $(BINPATH)

debug: generate
	go build -tags 'debug' -o build/dp-frontend-renderer
	HUMAN_LOG=1 DEBUG=1 ./build/dp-frontend-renderer

generate:
	# build the production version
	go generate ./...
	{ echo "// +build production"; cat assets/data.go; } > assets/data.go.new
	mv assets/data.go.new assets/data.go
	# build the dev version
	cd assets; go-bindata -debug -o debug.go -pkg assets templates/... locales/...
	{ echo "// +build debug"; cat assets/debug.go; } > assets/debug.go.new
	mv assets/debug.go.new assets/debug.go

test:
	go test -race -cover $(shell go list ./... | grep -v /vendor/) -tags 'production' ./...

.PHONY: build debug generate
