
FROM golang:latest

# Set the working directory in the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Initialize Go modules
# RUN go mod init

# Install Go dependencies
RUN go mod download


# Build the Go app
RUN go build -o main .

# Expose port 8082 to the outside world
EXPOSE 8082

# Command to run the executable
CMD ["./main"]
