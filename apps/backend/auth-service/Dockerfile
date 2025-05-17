# Build stage
FROM golang:1.21-alpine AS build

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.work ./
COPY apps/backend/auth-service/go.mod ./apps/backend/auth-service/
COPY packages/go/common/go.mod ./packages/go/common/

# Set working directory to the service
WORKDIR /app/apps/backend/auth-service

# Download dependencies
RUN go mod download

# Copy the rest of the application
COPY . /app

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /auth-service

# Final stage
FROM alpine:latest

# Install CA certificates
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /root/

# Copy the binary from the build stage
COPY --from=build /auth-service .

# Expose port
EXPOSE 8081

# Command to run the application
CMD ["./auth-service"]
