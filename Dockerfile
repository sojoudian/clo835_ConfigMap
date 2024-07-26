# Build stage
FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

# Run stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/app .
RUN chmod +x ./app
CMD ["./app"]
