# Start from a small, secure Linux version with Go installed
FROM golang:1.21-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy your source code into the container
COPY main.go .

# Initialize a Go module (required for modern Go)
RUN go mod init my-go-app

# Build the application
RUN go build -o server main.go

# Tell Docker this app listens on port 8080
EXPOSE 8080

# The command to run when the container starts
CMD ["./server"]
