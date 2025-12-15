# -------------------------------
# BUILD (PRODUCTION)
# -------------------------------
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
# Собираем бинарник для продакшена
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main ./cmd/server

# -------------------------------
# PRODUCTION (DISTROLESS)
# -------------------------------
FROM gcr.io/distroless/static-debian11 AS production

WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8003

# Запуск приложения
CMD ["/app/main"]
