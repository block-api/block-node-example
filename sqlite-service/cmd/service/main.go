package main

import (
	"github.com/block-api/block-node-example/sqlite-service/user"
	usermigration "github.com/block-api/block-node-example/sqlite-service/user/migration"
	"github.com/block-api/block-node/block"
	"github.com/block-api/block-node/log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	options := block.BlockNodeOptions{
		Name:    "sqlite-service",
		Version: 1,
	}

	blockNode := block.NewBlockNode(&options)
	userBlock := user.NewUserBlock(&blockNode)

	_ = blockNode.AddBlock(&userBlock)

	err := blockNode.Start()
	if err != nil {
		panic(err)
	}

	addDatabaseMigrations(&blockNode)

	err = blockNode.Database().RunMigrations()
	if err != nil {
		log.Panic(err.Error())
	}

	var osSignal = make(chan os.Signal)
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

func addDatabaseMigrations(bn *block.BlockNode) {
	dbDefault := bn.Database().GetSQLite("default")
	if dbDefault != nil {
		dbDefault.AddMigration(usermigration.UserMigration())
	}
}
