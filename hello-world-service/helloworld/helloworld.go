package helloworld

import (
	"github.com/block-api/block-node/block"
	"github.com/block-api/block-node/transporter"
)

type HelloWorldBlock struct {
	block.Block
}

func NewHelloWorldBlock(bn *block.BlockNode) HelloWorldBlock {
	helloWorldBlock := HelloWorldBlock{
		block.NewBlock(bn, "helloworld"),
	}

	helloWorldBlock.AddAction("hello", helloWorldBlock.ActionHello)

	return helloWorldBlock
}

func (hw *HelloWorldBlock) ActionHello(payload transporter.PayloadMessage) (*transporter.PayloadMessage, error) {
	responseMessage := transporter.PayloadMessage{
		Data: "Hello " + payload.Data.(string),
	}

	return &responseMessage, nil
}
