FROM golang:1.13.4-alpine

WORKDIR /usr/src/backend

RUN apk add git
RUN go get -u golang.org/x/lint/golint

COPY . ./
RUN go mod vendor
RUN go build

CMD ["go", "run", "."]
