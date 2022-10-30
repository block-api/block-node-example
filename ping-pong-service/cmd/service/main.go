package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/block-api/auth-service/pingpong"
	"github.com/block-api/block-node/block"
	"github.com/block-api/block-node/log"
)

func main() {
	options := block.BlockNodeOptions{
		Name:    "ping-pong-service",
		Version: 1,
	}

	blockNode := block.NewBlockNode(&options)
	pingpongBlock := pingpong.NewPingPongBlock(&blockNode)

	blockNode.AddBlock(&pingpongBlock)
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
