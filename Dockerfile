# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the entire project directory into the container
COPY . .

# Download and cache Go modules
RUN go mod download

# Build the Go application for Linux
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/main/main.go

# Expose port 8001 to the outside world
EXPOSE 8080

# Set environment variables
ENV PORT=8001

# Command to run the executable
CMD ["./main"]