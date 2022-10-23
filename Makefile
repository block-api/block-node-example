GOBIN = ./build/bin
GO ?= latest
GORUN = env GO111MODULE=on go run
GOBUILD = env GO111MODULE=on go build

.PHONY: cli service build clean

build:
	make service cli
dev:
	@echo "\n> --- run in development mode --"
	DEBUG=true DATA_DIR=./build go run ./cmd/main.go
service:
	mkdir -p $(GOBIN)
	cd ./cmd/service/ && go fmt ./... && $(GOBUILD) -o ./../../$(GOBIN)/example-service
	# cp config.example.yml ./build/config.yml
	chmod +x $(GOBIN)/example-service

	@echo "\n> ---"
	@echo "> Build successful. Executable in: \"$(GOBIN)/example-service\" "
	@echo "> ---\n"
clean:
	rm -rf $(GOBIN)