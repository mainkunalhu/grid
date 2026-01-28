
dev: # Run the grid application (development use only)
	@go run ./cmd/grid/main.go

build:# Build the binary of grid application
	go build ./cmd/grid/main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build ./cmd/grid/main.go -o grid
build-window:
	GOOS=windows GOARCH=arm64 go build ./cmd/grid/main.go -o grid.exe
build-mac:
	GOOS=darwin GOARCH=arm64 go build ./cmd/grid/main.go -o grid
