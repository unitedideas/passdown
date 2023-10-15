# Binary output location
BINARY = server/passdown

# Go build command
GO_BUILD = go build -o $(BINARY) ./server/src/main

# GitHub repository
GITHUB_REPO = git@github.com:unitedideas/passdown.git

# Build binary
build:
	$(GO_BUILD)

# Run the application
run: build
	$(BINARY)

# Clean up build artifacts
clean:
	rm -f $(BINARY)

# Update code from GitHub
update:
	git pull origin main


# Clean, update from GitHub, build, then run
all: clean update build run
