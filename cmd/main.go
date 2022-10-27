package main

import (
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

	authBlock := auth.NewAuthBlock()
	userBlock := user.NewUserBlock()

	blockNode := block.NewBlockNode(&options)

	blockNode.AddBlock(&authBlock, &userBlock)
	blockNode.Start()

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
