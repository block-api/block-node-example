FROM golang:1.18-alpine
RUN apk add build-base

ENV GO111MODULE on 
WORKDIR $GOPATH/src/github.com/block-node-example/sqlite-service

COPY . .

RUN go get -d -v ./...

RUN mkdir -p /var/block-node/sqlite-service/data
RUN make service

WORKDIR  $GOPATH/src/github.com/block-node-example/sqlite-service/build

CMD ["./bin/sqlite-service"]