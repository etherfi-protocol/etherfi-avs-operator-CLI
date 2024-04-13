.PHONY: build, test, coverage, clean

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

clean:
	@echo "Cleaning..."
	@rm -rf dist
