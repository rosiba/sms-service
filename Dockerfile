FROM golang:1.24-alpine
LABEL authors="rosiba"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o sms-service ./cmd/sms-service

CMD ["./sms-service"]