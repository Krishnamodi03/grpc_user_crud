# Step 1: Build the Go application
FROM golang:1.22-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the project
COPY . .

# Build the Go app
RUN go build -o grpc_user_crud main.go

# Step 2: Create a lightweight image to run the Go application
FROM alpine:latest

# Set working directory inside the container
WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder /app/grpc_user_crud .

COPY .env ./

# Expose gRPC port
EXPOSE 50052

# Command to run the binary
CMD ["./grpc_user_crud"]
