APP=tradeve

.PHONY: ready
ready: lint test build

all: ready

build:
		go mod download
		go mod tidy
		go build -o bin/${APP} .

install: build
		cp bin/${APP} ${GOPATH}/bin/${APP}
		[[ -f $(shell which codesign) ]] && codesign -s - ${GOPATH}/bin/${APP}

.PHONY: lint
lint:
		golangci-lint run --fix ./... --timeout=120s

.PHONY: test
test:
		 go test ./...

test-cover:
		go install github.com/ory/go-acc@latest
		go-acc -o coverage.out ./... -- -v
		go tool cover -html=coverage.out -o coverage.html
		open coverage.html

.PHONY: help
help: ## Display available commands
		@grep -E '^[a-zA-Z_-]+:.*?## .*$$' Makefile | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'