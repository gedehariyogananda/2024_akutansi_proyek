# Stage 1: Build
FROM golang:1.22-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o main .

# Stage 2: Final image (using a lightweight image like Alpine)
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

COPY .env /app/.env

EXPOSE ${SERVER_PORT}

CMD ["/app/main"]

