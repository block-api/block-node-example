GOBIN = ./build/bin
GO ?= latest
GORUN = env GO111MODULE=on go run
GOBUILD = env GO111MODULE=on go build

.PHONY: service build clean

build:
	make service
dev:
	@echo "\n> --- run in development mode --"
	DEBUG=true DATA_DIR=./build go run ./cmd/main.go
service:
	mkdir -p $(GOBIN)
	cd ./cmd/ && go fmt ./... && $(GOBUILD) -o ./../$(GOBIN)/hello-world-service
	chmod +x $(GOBIN)/hello-world-service

	cp config.example.yml ./build/config.yml
	cp .env.tpl ./build/.env
	
	@echo "\n> ---"
	@echo "> Build successful. Executable in: \"$(GOBIN)/hello-world-service\" "
	@echo "> ---\n"
clean:
	rm -rf $(GOBIN)