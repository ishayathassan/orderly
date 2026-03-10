# ---------- Build Stage ----------
FROM golang:1.25-alpine AS builder

WORKDIR /app

# install git (needed for some go modules)
RUN apk add --no-cache git

# copy go modules first (better caching)
COPY go.mod go.sum ./
RUN go mod download

# copy source code
COPY . .

# build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o oms-service ./cmd/oms

# ---------- Runtime Stage ----------
FROM alpine:latest

WORKDIR /app

# copy binary from builder
COPY --from=builder /app/oms-service .
COPY --from=builder /app/internal/database/migrations ./internal/database/migrations

EXPOSE 8080

CMD ["./oms-service"]