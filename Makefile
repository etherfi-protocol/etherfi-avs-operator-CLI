.PHONY: build, test, coverage, clean, mocks

build:
	@echo "Building..."
	@go build -o dist/etherfi ./cmd/etherfi/

test:
	@echo "Testing..."
	@go test -v ./...

coverage:
	@echo "Testing with coverage..."
	@go test -coverprofile=coverage.out ./...
	#@go tool cover -html=coverage.out

mocks:
	@echo "Generating mocks..."
	@mockery --recursive --output ./mocks --name BLSAvsSigner

clean:
	@echo "Cleaning..."
	@rm -rf dist
