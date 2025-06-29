# Build the application
all: build test

build:
	@echo "Building..."
	@go build -o bin/main cmd/app/main.go

# Run the application
run:
	@go run cmd/app/main.go

run-dev:
	@echo "Running in development mode..."
	@$(MAKE) run-entgo cmd=generate
	@./scripts/watch.sh all

run-entgo:
	@echo "Running Entgo..."
	@./scripts/ent-tools.sh ${cmd} ${name}

run-atlas-status:
	@echo "Running Atlas status..."
	@./scripts/atlas-tools.sh status

run-atlas-diff:
	@echo "Running Atlas migrate..."
	@./scripts/atlas-tools.sh diff ${name}

run-atlas-migrate:
	@echo "Running Atlas migrate..."
	@./scripts/atlas-tools.sh apply 

generate-usecase:
	@echo "Generating usecase: ${name}"
	@go run ./cmd/usecase-generator/main.go ${name}

generate-service:
	@echo "Generating service: ${name} inside of ${group}"
	@go run ./cmd/service-generator/main.go ${group} ${name}

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v

test-usecase:
	@echo "Testing Usecase..."
	@go test ./internal/core/usecase -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f bin/*

.PHONY: all build run run-dev test clean watch generate-usecase generate-service test-usecase
