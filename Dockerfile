FROM golang:latest

WORKDIR /app

# Copy the go.mod and go.sum files first to utilize Docker's cache
COPY go.mod ./
COPY go.sum ./

# Download dependencies specified in go.mod and go.sum
RUN go mod download

# Clean up the go.mod and go.sum files
RUN go mod tidy

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Command to run the built application
CMD ["./main"]
