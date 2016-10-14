build:
	go generate ./...
	go build -o build/dp-frontend-renderer

.PHONY: build
