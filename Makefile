build:
	go generate ./...
	go build -o build/dp-frontend-renderer

debug:
	cd assets; go-bindata -debug -o bindata.go -pkg assets templates/... lang/...
	go build -o build/dp-frontend-renderer
	HUMAN_LOG=1 DEBUG=1 ./build/dp-frontend-renderer

.PHONY: build debug
