FROM golang:1.24-alpine
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o api_service ./cmd/api_service

EXPOSE 8080

CMD ["./api_service"]
