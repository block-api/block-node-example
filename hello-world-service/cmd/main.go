package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/block-api/block-node-example/hello-world-service/helloworld"
	"github.com/block-api/block-node/block"
	"github.com/block-api/block-node/log"
)

func main() {
	options := block.BlockNodeOptions{
		Name:    "hello-world-service",
		Version: 1,
	}

	blockNode := block.NewBlockNode(&options)
	helloWorldBlock := helloworld.NewHelloWorldBlock(&blockNode)

	blockNode.AddBlock(&helloWorldBlock)
	blockNode.Start()

	go func(helloWorldBlock *helloworld.HelloWorldBlock) {
		http.HandleFunc("/hello", helloWorldBlock.ApiHello)
		http.HandleFunc("/ping", helloWorldBlock.ApiPing)

		http.ListenAndServe(":8090", nil)
	}(&helloWorldBlock)

	var osSignal chan os.Signal = make(chan os.Signal)
	signal.Notify(osSignal, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-osSignal:
			log.Warning("shutting down, please wait")
			err := blockNode.Stop()

			if err != nil {
				log.Panic(err.Error())
			}

			os.Exit(0)
		}
	}
}
