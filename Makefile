.PHONY: build, test
build:
	@echo "Building..."
	@go build -o avs-cli ./cmd/avs-cli/

test:
	@echo "Testing..."
	@go test -v ./...
