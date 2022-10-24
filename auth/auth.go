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
	block := AuthBlock{
		block.Block{
			Name:    "auth-block",
			Actions: make(map[string]block.BlockAction),
		},
	}
	block.AddAction("authorization", block.ActionAuthorization)
	block.AddAction("ping", block.ActionPing)

	return block
}

func (ab *AuthBlock) GetName() string {
	return ab.Block.Name
}

func (ab *AuthBlock) Actions() map[string]block.BlockAction {
	return ab.Block.Actions
}

func (ab *AuthBlock) AddAction(name string, action block.BlockAction) {
	ab.Block.Actions[name] = action
}

func (ab *AuthBlock) ActionAuthorization(payload interface{}) error {
	log.Default("-- ActionAuthorization --")
	fmt.Println(payload)
	return nil
}

func (ab *AuthBlock) ActionPing(payload interface{}) error {
	log.Default("-- ActionPing --")
	fmt.Println(payload)
	return nil
}
