FROM golang:alpine as builder
RUN apk update && apk upgrade

WORKDIR /go/src/app

ENV GO111MODULE=on

RUN go install github.com/cespare/reflex@latest

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./run ./cmd/.

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder go/src/app/run ./cmd/.

EXPOSE 8000
CMD ["./run"]