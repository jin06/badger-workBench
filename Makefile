APP_NAME = BadgerWorkbench
VERSION = 1.0.0
SRC = .
BIN_DIR = bin

# Build targets
.PHONY: all clean darwin windows linux help

all: darwin windows linux ## Build for all platforms

darwin: ## Build for macOS
	mkdir -p $(BIN_DIR)
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o $(BIN_DIR)/$(APP_NAME)-$(VERSION)-darwin-amd64 $(SRC)
	CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 go build -o $(BIN_DIR)/$(APP_NAME)-$(VERSION)-darwin-arm64 $(SRC)

windows: ## Build for Windowt st
	mkdir -p $(BIN_DIR)
	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -o $(BIN_DIR)/$(APP_NAME)-$(VERSION)-windows-amd64.exe $(SRC)
	CGO_ENABLED=1 GOOS=windows GOARCH=386 go build -o $(BIN_DIR)/$(APP_NAME)-$(VERSION)-windows-386.exe $(SRC)

linux: ## Build for Linux
	mkdir -p $(BIN_DIR)
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)/$(APP_NAME)-$(VERSION)-linux-amd64 $(SRC)
	CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -o $(BIN_DIR)/$(APP_NAME)-$(VERSION)-linux-arm64 $(SRC)

clean: ## Remove all built binaries
	rm -rf $(BIN_DIR)

help: ## Show help for make commands
	@echo "Usage:"
	@echo "  make all       - Build for all platforms (macOS, Windows, Linux)"
	@echo "  make darwin    - Build for macOS (amd64 and arm64)"
	@echo "  make windows   - Build for Windows (amd64 and 386)"
	@echo "  make linux     - Build for Linux (amd64 and arm64)"
	@echo "  make clean     - Remove all built binaries"
