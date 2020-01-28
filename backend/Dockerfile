FROM golang:1.13.4-alpine

WORKDIR /usr/src/backend

COPY . ./

RUN go mod vendor

RUN go build

CMD ["go", "run", "."]
