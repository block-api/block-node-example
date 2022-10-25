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
		block.NewBlock("auth-block"),
	}

	authBlock.AddAction("authorization", authBlock.ActionAuthorization)
	authBlock.AddAction("ping", authBlock.ActionPing)

	return authBlock
}

func (ab *AuthBlock) ActionAuthorization(payload any) error {
	log.Default("-- ActionAuthorization --")
	fmt.Println(payload)
	return nil
}

func (ab *AuthBlock) ActionPing(payload any) error {
	log.Default("-- ActionPing --")
	fmt.Println(payload)
	return nil
}
