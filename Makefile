build:
	go generate ./...
	go build -o build/dp-frontend-renderer

debug:
	cd assets; go-bindata -debug -o templates.go -pkg assets templates/...
	go build -o build/dp-frontend-renderer
	HUMAN_LOG=1 DEBUG=1 ./build/dp-frontend-renderer

.PHONY: build debug
