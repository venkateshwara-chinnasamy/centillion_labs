all: build

build:
	mkdir -p build
	go build -o build/pipeline

run: build
	./build/pipeline

unit-test:
	go test -v ./...

bench:
	go test -bench=. -benchmem ./...

coverage:
	go test -coverprofile=build/coverage.out ./...
	go tool cover -html=build/coverage.out -o build/coverage.html
	go tool cover -func=build/coverage.out

test: unit-test bench coverage

clean:
	rm -rf build

.PHONY: all build run unit-test bench coverage test clean