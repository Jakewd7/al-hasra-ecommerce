FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main ./cmd/main.go

FROM alpine:3.19

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

RUN mkdir -p /app/uploads

COPY --from=builder /app/main .

EXPOSE 8000

RUN addgroup -S appgroup && adduser -S appuser -G appgroup \
    && chown -R appuser:appgroup /app
USER appuser

CMD ["./main"]