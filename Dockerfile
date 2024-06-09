# Use the official Golang image as a build stage
FROM golang:1.21 as builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app for Linux
RUN GOOS=linux GOARCH=amd64 go build -o myapp

# Verify that the binary was created
RUN ls -l /app/myapp

# Use a minimal image for the final container
FROM alpine:latest

# Set the current working directory inside the container
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/myapp .

# Verify that the binary was copied correctly
RUN ls -l /root/myapp

# Ensure the binary is executable
RUN chmod +x /root/myapp

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./myapp"]
