# Binary output location
BINARY = server/passdown

# Go build command
GO_BUILD = go build -o $(BINARY) ./server/src/main

# GitHub repository
GITHUB_REPO = git@github.com:unitedideas/passdown.git

# PID file for dev-client process
PID_FILE = client.pid

# Build binary
build:
	$(GO_BUILD)

# Run the application
run: build
	$(BINARY)

# Run the frontend
dev-client:
	cd client && npm start & echo $$! > $(PID_FILE)

# Kill the frontend process using the saved PID
kill-client:
	lsof -i :3000 -t | xargs kill -9 || true
	if [ -f $(PID_FILE) ]; then kill `cat $(PID_FILE)` || true && rm $(PID_FILE); fi

# Run both server and client
dev: dev-client run

# Clean up build artifacts and kill frontend process
clean: kill-client
	rm -f $(BINARY)

# Update code from GitHub
update:
	git pull origin main

# Run both server and client
all: clean update build dev