FROM golang:1.24-alpine AS builder
LABEL authors="rosiba"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o sms-service ./cmd/sms-service

FROM alpine

RUN apk --no-cache add curl

WORKDIR /app

COPY --from=builder /app/.env .
COPY --from=builder /app/sms-service .

CMD ["/app/sms-service"]