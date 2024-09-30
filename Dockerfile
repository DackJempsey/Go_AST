# Build stage
FROM golang:1.19-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code to the working directory
COPY . .

RUN go mod init go_ast/go_ast

# Build the Go application
RUN go build -o server



# Final stage
FROM golang:1.19-alpine

# Set the working directory inside the container
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/server .
COPY --from=builder /app/comp .

# Expose port 8080 to the host
EXPOSE 8080

# Command to run the executable
CMD ["./server"]
# CMD ["./comp"]

