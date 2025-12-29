

# ---- BUILD UI ----
FROM oven/bun:1.3.4-alpine AS ui
WORKDIR /app

COPY ./cmd/ui .
RUN bun install
RUN bun run build

# ---- BUILD BACKEND ----
FROM golang:1.25.5-alpine AS builder
WORKDIR /app

COPY --exclude=./cmd/ui . .
COPY ./cmd/ui/embed.go ./cmd/ui/embed.go

# COPY ui/build ./ui/build
COPY --from=ui /app/build ./cmd/ui/build

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
ENV APP_ENV=production
RUN go mod tidy
RUN go build -ldflags="-s -w" -o ./dist/server ./cmd/main.go

# ---- FINAL (scratch) ----
FROM scratch
WORKDIR /app

COPY --from=builder /app/dist/server .

EXPOSE 3000
ENV APP_ENV=production
CMD ["./server"]

