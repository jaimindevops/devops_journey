# Stage 1: Build
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main .

# Stage 2: Run (The Secure Layer)
FROM alpine:latest
WORKDIR /app
# FIX: Force update all packages to fix CVEs
RUN apk update && apk upgrade --no-cache
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
