# Go Concurrent Pipeline

A concurrent data processing pipeline implemented in Go that processes integers with error handling and context cancellation.

## Features
- Concurrent processing using goroutines and channels
- Squaring numbers with simulated delay
- Filtering numbers > 1000
- Error handling for negative numbers
- Context cancellation support
- Unit tests and benchmarks

## Requirements
- Go 1.19 or higher
- Make

## Build & Run
```bash
# Build
make build

# Run
make run
```

## Testing & Reports
```bash
# Unit tests
make unit-test
# Outputs test results to console

# Benchmarks
make bench
# Generates:
- benchmark_report.txt (basic benchmark results)

# Code coverage
make coverage
# Generates in build/:
- coverage.out (raw coverage data)
- coverage.html (visual coverage report)
- Function coverage report in console

# All tests
make test
# Runs unit tests, benchmarks, and generates all reports
```

## Reports Location
All reports are generated in the `build/` directory:
```
build/
├── coverage.out     # Raw coverage data
├── coverage.html    # Visual coverage report
└── benchmark_report.txt   # Benchmark results
```

To view the coverage report:
1. Open `build/coverage.html` in a web browser
2. Green: covered code
3. Red: uncovered code
4. Coverage percentage shown for each function

## Project Structure
```
.
├── build/          # Build outputs and reports
├── main.go         # Pipeline implementation
├── main_test.go    # Tests
├── go.mod          # Module definition
└── Makefile        # Build commands
```

## Make Commands
- `make build`: Builds binary to build/
- `make run`: Builds and runs
- `make unit-test`: Runs unit tests
- `make bench`: Runs benchmarks
- `make coverage`: Generates coverage report
- `make test`: Runs all tests
- `make clean`: Removes build artifacts and reports