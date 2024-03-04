# Use the official Golang image as the base image
FROM golang:1.22 AS builder

# Set the working directory in the container
WORKDIR /app

# Copy the Go module and Go sum files for dependencies
COPY . .

# Download dependencies
RUN go mod download

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd

# Use a minimal image for the final image
FROM alpine:latest

# Set the working directory in the final image
WORKDIR /app

ENV CONFIG_PATH config/local.yaml

# Copy the binary from the builder stage to the final image
COPY --from=builder /app/main .
COPY --from=builder /app/static static
COPY --from=builder /app/config config

# Run the Go application
CMD ["./main"]