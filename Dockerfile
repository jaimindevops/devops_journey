# Start from a small, secure Linux version
FROM golang:1.21-alpine

# Set the working directory
WORKDIR /app

# Copy the source code first
COPY main.go .

# 1. Initialize the module INSIDE the container
RUN go mod init my-go-app

# 2. Download the Redis library INSIDE the container
RUN go get github.com/redis/go-redis/v9

# 3. Build the application
RUN go build -o server main.go

# Expose the port
EXPOSE 8080

# Run the server
CMD ["./server"]
