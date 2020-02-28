FROM golang:latest

RUN mkdir -p /go/src/github.com/moemoe89/practicing-memcached-golang

WORKDIR /go/src/github.com/moemoe89/practicing-memcached-golang

COPY . /go/src/github.com/moemoe89/practicing-memcached-golang

RUN go mod download
RUN go install

ENTRYPOINT /go/bin/practicing-memcached-golang
