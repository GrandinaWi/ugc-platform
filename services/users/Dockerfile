# ---------- build stage ----------
FROM golang:1.25-alpine AS builder

WORKDIR /app

# зависимости
COPY go.mod go.sum ./
RUN go mod download

# код
COPY . .

# сборка
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o users ./cmd/app


# ---------- runtime stage ----------
FROM alpine:3.20

WORKDIR /app

# сертификаты для https / postgres
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/users /app/users

EXPOSE 8080

CMD ["/app/users"]
