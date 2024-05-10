
# New Relic Code Challenge Solution

## How to Run the Program

**Prerequisite:** Go 1.22.2

### Command Line
**Build and Run the Program**
```bash
# Build the Go binary
make build

# Run the program with files
make run-files ARGS="-- ./texts/moby_dick.txt ./texts/brothers_karamazov.txt"

# Run the program with stdin
make run-stdin

# Run all tests
make test

# Clean up the Go binary
make clean

# Build the Docker image
make docker

# Run with files using Docker
make docker-run

# Run with stdin using Docker
make docker-stdin

# Build and Run Tests in Docker
make docker-test


## Project Structure

- **cmd**: Contains the main executable entry point.
- **pkg**: Modular packages containing specific functionalities.
