# Use an official Golang runtime as a base image
FROM golang:1.19.3

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY go.mod go.sum ./
# Download and install any required dependencies
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/app ./cmd/userservice

EXPOSE 6602

# Command to run the executable
CMD ["./app"]
