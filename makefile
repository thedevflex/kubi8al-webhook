BINARY_NAME := kubi8al-webhook
IMAGE_NAME := ghcr.io/thedevflex/kubi8al-webhook
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
GO_FILES := $(shell find . -type f -name '*.go' -not -path "./vendor/*")

LDFLAGS := -ldflags="-X 'main.Version=$(VERSION)' -X 'main.Commit=$(shell git rev-parse HEAD 2>/dev/null || echo "dev")' -X 'main.BuildTime=$(shell date -u '+%Y-%m-%d_%H:%M:%S')'"

all: build

build: $(BINARY_NAME)

$(BINARY_NAME): $(GO_FILES)
	@echo "Building $(BINARY_NAME) $(VERSION)"
	CGO_ENABLED=0 go build -v -o $(BINARY_NAME) $(LDFLAGS) ./cmd/main.go

run: build
	@echo "Running $(BINARY_NAME)"
	./$(BINARY_NAME)

docker-build:
	@echo "Building Docker image $(IMAGE_NAME):$(VERSION)"
	docker build -t $(IMAGE_NAME):$(VERSION) .
	docker tag $(IMAGE_NAME):$(VERSION) $(IMAGE_NAME):latest

docker-run: docker-build
	@echo "Running Docker container"
	docker run -p 8080:8080 \
		-e EMMITER_API_ADDRESS=$(EMMITER_API_ADDRESS) \
		-e WEBHOOK_SECRET=$(WEBHOOK_SECRET) \
		$(IMAGE_NAME):$(VERSION)

test:
	@echo "Running tests..."
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

lint:
	@echo "Running linters..."
	golangci-lint run

fmt:
	@echo "Formatting code..."
	go fmt ./...

clean:
	@echo "Cleaning..."
	go clean
	rm -f $(BINARY_NAME)
	rm -f coverage.txt


deps:
	@echo "Installing dependencies..."
	go mod download
	go mod verify

update-deps:
	@echo "Updating dependencies..."
	go get -u ./...
	go mod tidy

help:
	@echo "\nUsage: make <target>\n"
	@echo "Available targets:"
	@echo "  all           - Build the application (default)"
	@echo "  build         - Build the application"
	@echo "  run           - Run the application locally"
	@echo "  docker-build  - Build the Docker image"
	@echo "  docker-run    - Run the application in Docker"
	@echo "  test          - Run tests"
	@echo "  lint          - Run linters"
	@echo "  fmt           - Format code"
	@echo "  clean         - Remove build artifacts"
	@echo "  generate      - Generate code (mocks, etc.)"
	@echo "  deps          - Install dependencies"
	@echo "  update-deps   - Update dependencies"
	@echo "  help          - Show this help message"

.PHONY: all build run docker-build docker-run test lint fmt clean generate deps update-deps help