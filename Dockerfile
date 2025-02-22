# Stage 1: Build
FROM golang:1.24-alpine AS builder

# Install necessary dependencies
RUN apk add --no-cache git

# Set the working directory
WORKDIR /app

# Copy module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Change working directory to the correct location (where main.go is)
WORKDIR /app/cmd

# Build the application
RUN go build -o /app/real-estate-management

# Stage 2: Minimal Runtime
FROM alpine:latest

# Install CA Certificates
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /app

# Copy the compiled binary
COPY --from=builder /app/real-estate-management .

# Copy the migration files to the correct location
COPY --from=builder /app/deployment/migration /app/deployment/migration

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["./real-estate-management"]
