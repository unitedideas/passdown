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

# Run the frontend
dev-client:
	cd client && npm start &

# Run both server and client
dev: dev-client run

# Clean up build artifacts
clean:
	rm -f $(BINARY)

# Update code from GitHub
update:
	git pull origin main

# Run both server and client
all: clean update build dev
