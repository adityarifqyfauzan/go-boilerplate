# Build stage
FROM golang:1.23.4-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy source code
COPY . .

# Download dependencies
RUN go mod download

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final stage
FROM alpine:3.19

# Install runtime dependencies
RUN apk add --no-cache ca-certificates tzdata

# Set working directory
WORKDIR /app

# Copy the binary and necessary files from builder
COPY --from=builder /app/main .
COPY --from=builder /app/locales ./locales
COPY --from=builder /app/internal ./internal
COPY --from=builder /app/config ./config
COPY --from=builder /app/pkg ./pkg
COPY --from=builder /app/go.mod .
COPY --from=builder /app/go.sum .

# Expose port
EXPOSE 5001

# Run the application
CMD ["./main"] 