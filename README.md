[![CircleCI](https://circleci.com/gh/moemoe89/go-memcached-elen.svg?style=svg)](https://circleci.com/gh/moemoe89/go-memcached-elen)
[![codecov](https://codecov.io/gh/moemoe89/go-memcached-elen/branch/master/graph/badge.svg)](https://codecov.io/gh/moemoe89/go-memcached-elen)
[![Go Report Card](https://goreportcard.com/badge/github.com/moemoe89/go-memcached-elen)](https://goreportcard.com/report/github.com/moemoe89/go-memcached-elen)

# GO-MEMCACHED-ELEN #

Practicing Memcached Using Golang (Mux Framework) as Programming Language, Memcached as Database

## Directory structure
Your project directory structure should look like this
```
  + your_gopath/
  |
  +--+ src/github.com/moemoe89
  |  |
  |  +--+ go-memcached-elen/
  |     |
  |     +--+ main.go
  |        + api/
  |        + routers/
  |        + ... any other source code
  |
  +--+ bin/
  |  |
  |  +-- ... executable file
  |
  +--+ pkg/
     |
     +-- ... all dependency_library required

```

## Setup and Build

* Setup Golang <https://golang.org/>
* Setup Memcached <https://www.memcached.org/>
* Under `$GOPATH`, do the following command :
```
$ mkdir -p src/github.com/moemoe89
$ cd src/github.com/moemoe89
$ git clone <url>
$ mv <cloned directory> go-memcached-elen
```

## Running Application
Make config file for local :
```
$ cp config-sample.json config.json
```
Build
```
$ go build
```
Run
```
$ go run main.go
```

## How to Run with Docker
Make config file for docker :
```
$ cp config-sample.json config.json
```
Build
```
$ docker-compose build
```
Run
```
$ docker-compose up
```
Stop
```
$ docker-compose down
```