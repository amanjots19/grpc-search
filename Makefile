# Project directories
PROTO_DIR := proto
CMD_DIR := cmd
BUILD_DIR := bin

# Protobuf files
PROTO_FILES := $(wildcard $(PROTO_DIR)/*.proto)

# Output directories for generated code
PROTO_GO_OUT := $(PROTO_DIR)
PROTO_GO_GRPC_OUT := $(PROTO_DIR)

# Generated Go files
PROTO_GO_FILES := $(PROTO_FILES:.proto=.pb.go)
PROTO_GO_GRPC_FILES := $(PROTO_FILES:.proto=_grpc.pb.go)

# Binary name
BINARY := $(BUILD_DIR)/server

# Main files
MAIN_FILES := $(CMD_DIR)/*.go

# Default target
all: build

# Create the build directory if it doesn't exist
$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

# Generate Go files from proto files
proto: $(PROTO_FILES)
	@echo "Generating Go files from proto files..."
	protoc --go_out=$(PROTO_GO_OUT) --go_opt=paths=source_relative \
	       --go-grpc_out=$(PROTO_GO_GRPC_OUT) --go-grpc_opt=paths=source_relative \
	       $(PROTO_FILES)

# Build the server binary
$(BINARY): proto $(MAIN_FILES) | $(BUILD_DIR)
	@echo "Building the binary..."
	go build -o $(BINARY) $(MAIN_FILES)

# Build target
build: $(BINARY)

# Clean up generated files and binaries
clean:
	@echo "Cleaning up..."
	rm -f $(PROTO_GO_OUT)/*.pb.go $(PROTO_GO_GRPC_OUT)/*_grpc.pb.go
	rm -f $(BINARY)
	rm -rf $(BUILD_DIR)

# Run the server
run: $(BINARY)
	@echo "Running the server..."
	./$(BINARY)

# Phony targets
.PHONY: all proto build clean run
