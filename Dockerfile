FROM golang:latest

RUN mkdir -p /go/src/github.com/moemoe89/go-memcached-elen

WORKDIR /go/src/github.com/moemoe89/go-memcached-elen

COPY . /go/src/github.com/moemoe89/go-memcached-elen

RUN go mod download
RUN go install

ENTRYPOINT /go/bin/go-memcached-elen
