package user

import (
	"fmt"

	"github.com/block-api/block-node/block"
	"github.com/block-api/block-node/log"
)

type UserBlock struct {
	block.Block
}

func NewUserBlock() UserBlock {
	userBlock := UserBlock{
		block.NewBlock("user"),
	}

	userBlock.AddAction("get-user", userBlock.ActionGetUser)

	return userBlock
}

func (ub *UserBlock) ActionGetUser(payload any) error {
	log.Default("-- ActionGetUser --")
	fmt.Println(payload)
	return nil
}
