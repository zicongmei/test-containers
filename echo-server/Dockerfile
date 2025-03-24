# Use the official Golang image as a base image
FROM golang:1.22-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to cache dependencies
COPY go.mod ./

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o server .

# Use a minimal Alpine Linux image for the final stage
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the built executable from the builder stage
COPY --from=builder /app/server .

# Expose the port the application listens on
EXPOSE 8080

# Command to run the executable
CMD ["./server"]