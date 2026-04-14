# === Stage 1 (Builder)
FROM golang:1.25.4-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum file to the working directory sothat we can download the dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code to the working directory
COPY . .

# Build the Go Application
RUN go build -o main ./cmd/api

# === Stage 2 (Runtime)
FROM alpine:latest

WORKDIR /app

# Copy from the builder stage the built binary
COPY --from=builder /app/main .

# If .env file is needed, copy it to the runtime stage
COPY .env .

# Expose the port that the application will run on
EXPOSE 8080

# Command to run the application
CMD ["./main"]