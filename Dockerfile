# Use the official Golang image to create a build artifact.
FROM golang:1.23-alpine AS builder

# Install PostgreSQL client
RUN apk add --no-cache postgresql-client bash

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main ./cmd/app

# Start a new stage from scratch
FROM alpine:latest  

# Install PostgreSQL client and bash
RUN apk add --no-cache postgresql-client bash

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
COPY scripts/wait-for-postgres.sh ./wait-for-postgres.sh

# Make the script executable
RUN chmod +x ./wait-for-postgres.sh

# Wait for postgres and start the application
CMD /bin/sh -c "./wait-for-postgres.sh db && ./main" 