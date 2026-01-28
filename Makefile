dev: # Run the grid application (development use only)
	@go run ./cmd/grid/main.go

build: # Build the binary of grid application
	go build -o grid ./cmd/grid/main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -o grid ./cmd/grid/main.go

build-window:
	GOOS=windows GOARCH=amd64 go build -o grid.exe ./cmd/grid/main.go

build-mac:
	GOOS=darwin GOARCH=arm64 go build -o grid ./cmd/grid/main.go
