FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o food-ordering-service .

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/food-ordering-service .

EXPOSE 3000

CMD ["./food-ordering-service"]