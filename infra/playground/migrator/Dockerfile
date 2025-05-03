FROM --platform=$BUILDPLATFORM golang:latest

# Install goose
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Create non-root user
ARG UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    appuser

# Copy goose binary to user's PATH
RUN cp /go/bin/goose /usr/local/bin/

# Switch to non-root user
USER appuser

WORKDIR /app