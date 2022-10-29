package auth

import (
	"fmt"

	"github.com/block-api/block-node/block"
	"github.com/block-api/block-node/log"
	"github.com/block-api/block-node/transporter"
)

type AuthBlock struct {
	block.Block
}

func NewAuthBlock(bn *block.BlockNode) AuthBlock {
	authBlock := AuthBlock{
		block.NewBlock(bn, "auth"),
	}

	authBlock.AddAction("authorization", authBlock.ActionAuthorization)

	return authBlock
}

func (ab *AuthBlock) ActionAuthorization(payload []byte) (any, error) {
	var response any

	body, err := block.DecodePayload[transporter.PayloadMessage](payload)
	if err != nil {
		log.Warning(err.Error())
		return response, err
	}

	fmt.Println(body)

	return response, nil
}
