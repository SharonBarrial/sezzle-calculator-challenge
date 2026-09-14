# --- Build stage ---
FROM golang:1.24-alpine AS builder
WORKDIR /app

COPY go.mod ./
COPY . .

# CGO_ENABLED=0 produces a static binary we can run on a minimal base image.
RUN CGO_ENABLED=0 GOOS=linux go build -o /backend .

# --- Runtime stage ---
FROM alpine:3.20
RUN addgroup -S app && adduser -S app -G app
WORKDIR /app
COPY --from=builder /backend .
USER app

EXPOSE 8080
ENV PORT=8080

CMD ["./backend"]