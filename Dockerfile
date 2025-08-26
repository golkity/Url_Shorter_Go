FROM golang:1.23.0-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@v1.16.3
RUN swag init -g cmd/main.go

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/url-shortener ./cmd/main.go


FROM alpine:3.20 AS shortner

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

USER appuser

WORKDIR /app

COPY --from=builder /app/url-shortener /app/url-shortener
COPY --from=builder /app/docs /app/docs

EXPOSE 8080

ENTRYPOINT ["/app/url-shortener"]