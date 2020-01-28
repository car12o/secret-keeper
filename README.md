# secret-keeper

A secret server to store and share secrets. Each secret can be read only a limited number of times after that it will expire and won’t be available anymore besides the secret may have TTL and after the expiration time the secret won’t be available anymore as well.

The API response can be in XML or JSON, based on the Accept header.

Service exports metrics to a Prometheus.

## Installation - normal

#### Requirements
The first need [Go](https://golang.org/) installed (**version 1.12+ is required**)
The second need [MongoDB](https://www.mongodb.com/) installed (**version 4.1.13+ is required**)
The third need **export** environment variables (**check docker-compose file**)

1. Install dependencies
```sh
$ go mod vendor
```
2. Install application
```sh
$ go install -v
```
3. Run application
```sh
$ secret-keeper
```

## Installation - using Docker

#### Requirements
The first need [Docker](https://docs.docker.com/install) installed (**version 18.09.7+ is required**)
The second need [docker-compose](https://docs.docker.com/compose) installed (**version 3.1+ is required**)

1. Run docker-compose
```sh
$ docker-compose up
```
