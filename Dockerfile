# Build stage
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o bank .

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bank /app/
EXPOSE 8080
CMD ["./bank"]
