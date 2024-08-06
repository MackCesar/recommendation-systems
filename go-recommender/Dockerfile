# Use the official Golang image to build the application
FROM golang:1.18 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Copy the Go application source code
COPY main.go .

# Copy the data directory
COPY data/ ./data/

# Download dependencies
RUN go mod download

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-recommender .

# Use a lightweight image to run the Go application
FROM alpine:3.14

# Set the working directory inside the container
WORKDIR /app

# Copy the compiled binary and data from the builder stage
COPY --from=builder /app/go-recommender .
COPY --from=builder /app/data/ ./data/

# Expose the application port
EXPOSE 8080

# Command to run the application with the default file path
ENTRYPOINT ["./go-recommender"]
CMD ["--file", "data/Video_Games.jsonl"]