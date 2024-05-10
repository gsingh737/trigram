# Use an official Golang image as a build stage
FROM golang:1.22.3 AS builder

# Set the working directory
WORKDIR /app

# Copy the Go modules and source files
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the Go binary
RUN go build -o /trigrams ./cmd/main.go

# Use a lightweight base image
FROM gcr.io/distroless/base-debian11

# Copy the Go binary from the build stage
COPY --from=builder /trigrams /trigrams

# Set the command to run the binary
ENTRYPOINT ["/trigrams"]
