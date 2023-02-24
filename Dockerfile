# syntax=docker/dockerfile

FROM golang:1.20.1-alpine3.17
WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./


EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
