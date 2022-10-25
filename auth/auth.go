package auth

import (
	"fmt"

	"github.com/block-api/block-node/block"
	"github.com/block-api/block-node/log"
)

type AuthBlock struct {
	block.Block
}

func NewAuthBlock() AuthBlock {
	authBlock := AuthBlock{
		block.NewBlock("auth"),
	}

	authBlock.AddAction("authorization", authBlock.ActionAuthorization)

	return authBlock
}

func (ab *AuthBlock) ActionAuthorization(payload any) error {
	log.Default("-- ActionAuthorization --")
	fmt.Println(payload)
	return nil
}
