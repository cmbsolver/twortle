# Stage 1: Build the application
FROM docker.io/library/golang:1.25-alpine AS builder

# Install build dependencies for CGO (required if using SQLite)
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

# Copy dependency files
COPY go.mod ./
# If you are using go.work, you might need to handle it,
# but usually for a single image, go.mod is sufficient.
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
# CGO_ENABLED=1 is usually required for the go-sqlite3 driver
RUN CGO_ENABLED=1 GOOS=linux go build -o twortle-app main.go

# Stage 2: Final lightweight image
FROM docker.io/library/alpine:latest

RUN apk add --no-cache ca-certificates libc6-compat

WORKDIR /root/

# Copy the binary from the builder
COPY --from=builder /app/twortle-app .

# Copy essential runtime directories and files
COPY --from=builder /app/assets ./assets
COPY --from=builder /app/views ./views

# Expose the port your app runs on (adjust if different)
EXPOSE 3000

# Run the application
CMD ["./twortle-app"]