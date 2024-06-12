.PHONY: build run test coverage clean fmt deps

# Variables
APP_NAME := goleagues
MAIN_PACKAGE := ./cmd/api/main.go
BUILD_DIR := ./build
TEST_COVERAGE_OUT := coverage.out
COVERAGE_HTML := coverage.html
EXT :=

# Detect OS
ifeq ($(OS),Windows_NT)
	EXT := .exe
endif

# Build the application
build:
	go build -o $(BUILD_DIR)/$(APP_NAME)$(EXT) $(MAIN_PACKAGE)

# Run the application
run:
	go run $(MAIN_PACKAGE)

# Run all tests with verbose output
test:
	go test -v ./...

# Run tests with coverage and generate coverage report in HTML
coverage:
	go test -v -coverprofile=$(TEST_COVERAGE_OUT) ./...
	go tool cover -html=$(TEST_COVERAGE_OUT) -o $(COVERAGE_HTML)

# Clean up binaries and coverage files
clean:
	go clean
	rm -r $(BUILD_DIR)
	rm -f $(TEST_COVERAGE_OUT) $(COVERAGE_HTML)

# Format the code
fmt:
	go fmt ./...

# Install dependencies
deps:
	go mod tidy
	go mod download
