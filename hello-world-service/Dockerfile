FROM golang:1.18-alpine
RUN apk add build-base

ENV GO111MODULE on 
WORKDIR $GOPATH/src/github.com/block-node-example/hello-world-service

COPY . .

RUN go get -d -v ./...

RUN mkdir -p /var/block-node/hello-world-service/data
RUN make service

WORKDIR  $GOPATH/src/github.com/block-node-example/hello-world-service/build

EXPOSE 8090

CMD ["./bin/hello-world-service"]