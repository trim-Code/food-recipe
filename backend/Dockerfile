# Use the official Golang image as the base image
FROM golang:1.23-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download the Go modules
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main ./cmd

# Print the contents of the /app directory
RUN ls -la /app

# Expose the port the application runs on
EXPOSE 7001

# Command to run the executable
CMD ["./main"]
