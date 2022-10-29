package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/block-api/block-node-example/auth"
	"github.com/block-api/block-node-example/user"
	"github.com/block-api/block-node/block"
	"github.com/block-api/block-node/log"
)

func main() {
	options := block.BlockNodeOptions{
		Name:    "block-node-example",
		Version: 1,
	}

	blockNode := block.NewBlockNode(&options)

	authBlock := auth.NewAuthBlock(&blockNode)
	userBlock := user.NewUserBlock(&blockNode)

	blockNode.AddBlock(&authBlock, &userBlock)
	blockNode.Start()

	go func(authBlock *auth.AuthBlock) {
		http.HandleFunc("/hello", authBlock.ApiHello)
		http.HandleFunc("/ping", authBlock.ApiPing)

		http.ListenAndServe(":8090", nil)
	}(&authBlock)

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
