package pingpong

import (
	"github.com/block-api/block-node/block"
	"github.com/block-api/block-node/transporter"
)

type PingPongBlock struct {
	block.Block
}

func NewPingPongBlock(bn *block.BlockNode) PingPongBlock {
	serviceBlock := PingPongBlock{
		block.NewBlock(bn, "pingpong"),
	}

	serviceBlock.AddAction("ping", serviceBlock.ActionPing)

	return serviceBlock
}

func (ab *PingPongBlock) ActionPing(payload transporter.PayloadMessage) (*transporter.PayloadMessage, error) {
	var response *transporter.PayloadMessage = new(transporter.PayloadMessage)
	response.Data = "pong"

	return response, nil
}
