# Change these variables as necessary.
PACKAGE_PATH := ./bin/htmgenie
BINARY_NAME := htmgenie.exe

vet: ## Run go vet against code
	go vet ./...

tidy:
	go fmt ./...
	go mod tidy -v

build:
	go build -o=./bin/${BINARY_NAME} ${MAIN_PACKAGE_PATH}

test:
	go test ./... -coverprofile=coverage.txt -covermode=atomic
	go tool cover -func=coverage.txt
	go tool cover -html=coverage.txt -o cover.html

install:
	go install ./...
