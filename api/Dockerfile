FROM golang:1.23-alpine

WORKDIR /go/src/app

RUN go install github.com/air-verse/air@v1.61.7

CMD ["air", "-c", ".air.toml"]
