# Use a multi-stage build to optimize the image size and security

# Build stage
FROM golang:alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker's cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN go build -o main .

# Runtime stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose the port your application runs on
EXPOSE 8080

# Command to run the application
CMD ["./main"]