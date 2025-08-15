FROM golang:1.24-alpine

WORKDIR /app

# Install bash, netcat, git, and Air
RUN apk add --no-cache bash netcat-openbsd git && \
    go install github.com/air-verse/air@latest

# Copy Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy all source code
COPY . .

# Expose your app port
EXPOSE 8080

# Run Air with your configuration
CMD ["air", "-c", ".air.toml"]
