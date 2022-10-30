GOBIN = ./build/bin
GO ?= latest
GORUN = env GO111MODULE=on go run
GOBUILD = env GO111MODULE=on go build

.PHONY: dev service build clean

build:
	make service
dev:
	@echo "\n> --- run in development mode --"
	DEBUG=true DATA_DIR=./build go run ./cmd/service/main.go
service:
	mkdir -p $(GOBIN)
	cd ./cmd/service/ && go fmt ./... && $(GOBUILD) -o ./../../$(GOBIN)/ping-pong-service
	cp config.example.yml ./build/config.yml
	cp .env.tpl ./build/.env
	chmod +x $(GOBIN)/ping-pong-service

	@echo "\n> ---"
	@echo "> Build successful. Executable in: \"$(GOBIN)/ping-pong-service\" "
	@echo "> ---\n"
clean:
	rm -rf $(GOBIN)