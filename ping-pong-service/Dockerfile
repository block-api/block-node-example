FROM golang:1.18-alpine
RUN apk add build-base

ENV GO111MODULE on 
WORKDIR $GOPATH/src/github.com/block-node-example/ping-pong-service

COPY . .

RUN go get -d -v ./...

RUN mkdir -p /var/block-node/ping-pong-service/data
RUN make service

WORKDIR  $GOPATH/src/github.com/block-node-example/ping-pong-service/build

CMD ["./bin/ping-pong-service"]