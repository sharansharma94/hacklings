# Define the Go source files
SRC = $(wildcard *.go)

# Define the output binary name
OUT = ./bin/hacklings

# Default target to build the project
all: build

# Build the project
build:
	go build -o $(OUT) $(SRC)

# Clean the build artifacts
clean:
	rm -f $(OUT)

# Run the project
run: build
	./$(OUT)

# Run tests
test:
	go test ./...

# Watch for changes and rebuild
watch:
	reflex -r '\.go$$' make run

.PHONY: all build clean run test watch
