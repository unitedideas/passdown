# Binary output location
BINARY = server/passdown

# Go build command
GO_BUILD = go build -o $(BINARY) ./server/src/main

# Build binary
build:
	$(GO_BUILD)

# Run the application
run: build
	$(BINARY)

# Clean up build artifacts
clean:
	rm -f $(BINARY)

# Clean, build, then run
all: clean build run
