FROM golang:latest

RUN mkdir /app
ADD . /app
ADD .env /app/.env
WORKDIR /app
RUN go build -o main .

EXPOSE ${SERVER_PORT}

CMD ["sh", "-c", "/app/main"]