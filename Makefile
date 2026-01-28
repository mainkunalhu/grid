
dev: # Run the grid application (development use only)
	@go run ./cmd/grid/main.go

build:# Build the binary of grid application
	go build ./cmd/grid/main.go
