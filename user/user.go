package user

import (
	"fmt"

	"github.com/block-api/block-node/block"
	"github.com/block-api/block-node/log"
)

type UserBlock struct {
	block.Block
}

func NewUserBlock(bn *block.BlockNode) UserBlock {
	userBlock := UserBlock{
		block.NewBlock(bn, "user"),
	}

	userBlock.AddAction("get-user", userBlock.ActionGetUser)

	return userBlock
}

func (ub *UserBlock) ActionGetUser(payload []byte) (any, error) {
	var response any
	log.Default("-- ActionGetUser --")
	fmt.Println(payload)

	return response, nil
}
