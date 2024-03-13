# Name of the output binary
BINARY_NAME := jp.sh

# Build the project
compile:
	go build -o build/$(BINARY_NAME) main.go

# Run the compiled binary
run: build
	build/$(BINARY_NAME)

dev:
	go run main.go

# Clean up generated files
clean:
	go clean
	rm -f build/$(BINARY_NAME)

