# Use the official Go 1.21 image
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Copy the source code to the container
COPY . .

# Build the Go app
RUN go build -o myapp .

# Expose the port your app runs on
EXPOSE 8080

# Command to run the binary
CMD ["./myapp"]
