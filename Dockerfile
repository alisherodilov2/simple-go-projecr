# Stage 1: Build
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

# Stage 2: Run
FROM alpine:latest
RUN apk add --no-cache bash postgresql-client
WORKDIR /app
COPY --from=builder /app/main .
COPY wait-for-db.sh .
RUN chmod +x wait-for-db.sh
EXPOSE 8080
CMD ["./wait-for-db.sh", "db", "./main"]
