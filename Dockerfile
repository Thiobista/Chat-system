# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server .

# Final stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8082

CMD ["./server"]