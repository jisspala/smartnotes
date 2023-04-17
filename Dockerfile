# Use the official golang:1.20 image as the base image
FROM golang:1.20

# Set the working directory to /app
WORKDIR /smartnotes

# Copy the Go modules files to the working directory
COPY go.mod go.sum ./

# Download the Go modules
RUN go mod download

# Copy the rest of the application code to the working directory
COPY . .

# Build the application
RUN go build -o smartnotes .

# Expose port 8080 for the application
EXPOSE 8001

# Start the application
CMD ["./smartnotes"]