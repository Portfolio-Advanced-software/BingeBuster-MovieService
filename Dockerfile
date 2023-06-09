# Use an official Golang runtime as a parent image
FROM golang:1.19.8-alpine

# Set the working directory to /go/src/app
WORKDIR /go/src/app

# Copy the current directory contents into the container at /go/src/app
COPY . /go/src/app

# Download dependencies
RUN go get -d -v ./...

# Build the Go app
RUN go build -o app

# Create a non-root user and switch to that user
RUN adduser -D movie-service
USER movie-service

# Expose port 50055 for the application
EXPOSE 50055

# Define the command to run the executable
CMD ["./app"]
