# Stage 1: Build the Go binary
FROM golang:1.22.2-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./main.go

FROM alpine:3.16

WORKDIR /app

COPY --from=builder /app/main .

COPY .env .

RUN chmod +x ./main

CMD ["./main"]
