# Fase 1: Build
FROM golang:1.22.0 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 go build -o main cmd/api/main.go

# Fase 2: Imagen final
FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/main ./
COPY --from=builder /app/internal/data/client_portfolio.json /app/internal/data/client_portfolio.json
COPY --from=builder /app/internal/data/items_portfolio.json /app/internal/data/items_portfolio.json


COPY .env .env

ENV PORT=3002 \
    MONGO_DATABASE=portfolio-data \
    MONGO_HOST=mongo \
    MONGO_PORT=27017

EXPOSE 3002

CMD ["./main"]
