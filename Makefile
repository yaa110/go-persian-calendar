.PHONY: test
test:
	@go test ./... -v -race -cover

.PHONY: lint
lint: tools
	@$(GOLANGCI_LINT) run -c .golangci.yaml

GOLANGCI_LINT ?= go tool -modfile=tools/go.mod golangci-lint

.PHONY: tools
tools:
	@go get -modfile=tools/go.mod -tool github.com/golangci/golangci-lint/v2/cmd/golangci-lint
