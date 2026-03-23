# Use official Go image
FROM golang:1.22-alpine

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first (dependency caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port
EXPOSE 8080

# Run the application
CMD ["./main"]