package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/block-api/block-node-example/auth"
	"github.com/block-api/block-node-example/user"
	"github.com/block-api/block-node/block"
	"github.com/block-api/block-node/common/types"
	"github.com/block-api/block-node/log"
	"github.com/block-api/block-node/transporter"
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

	/**
	 * temporairly placed pocket sending test
	 */
	target := ""
	payload := transporter.PayloadDiscovery{
		NodeID:  blockNode.NodeID().String(),
		Name:    blockNode.GetName(),
		Version: blockNode.Version(),
		Event:   transporter.EventConnected,
		Blocks:  blockNode.Blocks(),
	}

	pocket := transporter.NewPocket[transporter.PayloadDiscovery](transporter.ChanDiscovery, blockNode.NodeID(), types.NodeID(target), payload)
	blockNode.Network().Send(pocket)
	/** ** ** **/

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
