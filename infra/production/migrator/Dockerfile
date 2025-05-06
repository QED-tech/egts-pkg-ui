# Используем многостадийную сборку для уменьшения размера
FROM --platform=$BUILDPLATFORM golang:latest as builder

# Устанавливаем goose с явным указанием целевой платформы
RUN GOOS=linux GOARCH=amd64 go install github.com/pressly/goose/v3/cmd/goose@latest

# Финальный образ
FROM alpine:latest AS final

# Копируем бинарник из стадии builder
COPY --from=builder /go/bin/linux_amd64/goose /usr/local/bin/goose

# Создаем non-root пользователя
ARG UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    appuser

# Переключаемся на non-root пользователя
USER appuser

WORKDIR /app