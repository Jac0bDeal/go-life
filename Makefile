all: clean lint test bin

clean:
	@echo "Cleaning bin/..."
	@rm -rf bin/*

dependencies:
	@echo "Installing project dependencies..."
	@go get -u golang.org/x/lint/golint

bin:
	@echo "Building go-life binary for use on local system..."
	@go build -o bin/go-life main.go

lint:
	@echo "Running linters..."
	@go vet ./...
	@golint ./...

test:
	@echo "Running tests..."
	@go test ./...
