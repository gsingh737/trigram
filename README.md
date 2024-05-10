
# New Relic Code Challenge Solution

## How to Run the Program

**Prerequisite:** Go 1.22.2

### Command Line
```bash
# Running with files
go run ./cmd/main.go -- ./texts/moby_dick.txt ./texts/brothers_karamazov.txt

# Running with stdin
cat ./texts/*.txt | go run ./cmd/main.go
```

### Running Tests
```bash
# Run all tests
go test ./...
```

## Project Structure

- **cmd**: Contains the main executable entry point.
- **pkg**: Modular packages containing specific functionalities.
