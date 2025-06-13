FROM golang:1.21-alpine AS builder
WORKDIR /app
RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o kubi8al-webhook ./cmd/main.go

# Final stage
FROM alpine:3.18

WORKDIR /app

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/kubi8al-webhook .

ENV PORT=80
ENV EMMITER_API_ADDRESS=""
ENV WEBHOOK_SECRET=""
EXPOSE ${PORT}

ENTRYPOINT ["/app/kubi8al-webhook"]

LABEL org.opencontainers.image.title="kubi8al-webhook" \
      org.opencontainers.image.authors="Harsh Anand <harsh@devflex.co.in>" \
      org.opencontainers.image.description="Webhook receiver for kubi8al that forwards events to an emitter API" \
      org.opencontainers.image.version="1.0.0" \
      org.opencontainers.image.source="https://github.com/thedevflex/kubi8al-webhook"