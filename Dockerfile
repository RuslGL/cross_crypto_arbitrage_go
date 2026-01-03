FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o arb-bot ./cmd/arb-bot

# ---- runtime ----
FROM alpine:3.20

WORKDIR /app
COPY --from=builder /app/arb-bot /app/arb-bot

CMD ["/app/arb-bot"]
