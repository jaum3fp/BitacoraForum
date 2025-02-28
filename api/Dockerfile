FROM golang:1.23-alpine

WORKDIR /go/src/app

RUN go install github.com/air-verse/air@latest

CMD ["air", "-c", ".air.toml"]