# Stage 1: Build the Go application

FROM golang:1.19-alpine AS builder

# Install necessary tools for building
RUN apk add --no-cache git

# Set the working directory
WORKDIR /app

# Copy module dependency files first
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project into the container
COPY . .

# Build the Go application from the /cmd/app directory
RUN go build -o ip2country ./cmd/app

# Stage 2: Create a minimal image for running the app
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/ip2country .

# Copy the CSV file into the runtime container
COPY resources/csv/ip2location.csv resources/csv/

# Copy the .env file to the working directory
COPY .env ./

# Set the entry point for the container
CMD ["./ip2country"]