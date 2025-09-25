FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /auth-service ./cmd/main.go


FROM gcr.io/distroless/static-debian12

COPY --from=builder /auth-service /auth-service

EXPOSE 8080

ENTRYPOINT ["/auth-service"]