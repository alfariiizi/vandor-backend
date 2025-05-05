# Build the application
all: build test

build:
	@echo "Building..."
	
	
	@go build -o bin/main main.go

# Run the application
run:
	@go run main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f bin/*

# Live Reload
watch:
	./scripts/watch.sh $(MODULE)

.PHONY: all build run test clean watch
