FROM golang:1.23.2 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o /app/tlunch-backend ./cmd/main/main.go
RUN go build -o /app/migrator ./cmd/migrator/main.go

FROM ubuntu:22.04

RUN apt-get update && apt-get install -y postgresql-client
WORKDIR /app
COPY --from=builder /app/tlunch-backend ./
COPY --from=builder /app/migrator ./
COPY --from=builder /app/internal/migrations ./internal/migrations
COPY configs ./configs


EXPOSE 8080
CMD ["sh", "-c", "./migrator && ./tlunch-backend"]
