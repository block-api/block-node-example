package auth

import (
	"fmt"

	"github.com/block-api/block-node/block"
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
	authBlock.AddAction("hello", authBlock.ActionHello)

	return authBlock
}

func (ab *AuthBlock) ActionAuthorization(payload transporter.PayloadMessage) (*transporter.PayloadMessage, error) {
	fmt.Println(payload)

	return nil, nil
}

func (ab *AuthBlock) ActionHello(payload transporter.PayloadMessage) (*transporter.PayloadMessage, error) {
	responseMessage := transporter.PayloadMessage{
		Data: "Hello " + payload.Data.(string),
	}

	return &responseMessage, nil
}
