# Stage 1: Build the Go application
FROM golang:1.23 AS builder

WORKDIR /app

COPY config.env ./

# Copy the Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o transaction_processor_binary main.go

# Stage 2: Run the application in a minimal Alpine image
FROM alpine:latest

RUN apk add --no-cache file
COPY --from=builder /app/config.env /app/config.env
# Copy the built binary
COPY --from=builder /app/transaction_processor_binary /app/transaction_processor_binary

# Set execution permissions
RUN chmod +x /app/transaction_processor_binary

WORKDIR /app

EXPOSE 8080

CMD ["/app/transaction_processor_binary"]
