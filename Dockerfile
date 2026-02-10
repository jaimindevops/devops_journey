# Stage 1: Build (Using AWS Public Mirror for Golang)
FROM public.ecr.aws/docker/library/golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
# Recommended: Ensure dependencies are downloaded properly
RUN go mod download 
RUN go build -o main .

# Stage 2: Run (Using AWS Public Mirror for Alpine)
FROM public.ecr.aws/docker/library/alpine:latest
WORKDIR /app
# FIX: Force update all packages to fix CVEs
RUN apk update && apk upgrade --no-cache
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
