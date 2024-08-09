.PHONY: build, test, mocks

build:
	@echo "Building..."
	@go build -o avs-cli ./bin/avs-cli/

test:
	@echo "Testing..."
	@go test -v ./...

mocks:
	@echo "Generating mocks..."
	@mockery --recursive --output ./mocks --name BLSAvsSigner

