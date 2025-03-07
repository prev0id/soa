FROM golang:1.24-alpine
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o user_service ./cmd/user_service

EXPOSE 8081

CMD ["./user_service"]
