package pingpong

import (
	"github.com/block-api/block-node/block"
	"github.com/block-api/block-node/transporter"
)

type PingPongBlock struct {
	block.Block
}

func NewPingPongBlock(bn *block.BlockNode) PingPongBlock {
	block := PingPongBlock{
		block.NewBlock(bn, "pingpong"),
	}

	block.AddAction("ping", block.ActionPing)

	return block
}

func (ab *PingPongBlock) ActionPing(payload transporter.PayloadMessage) (*transporter.PayloadMessage, error) {
	var response *transporter.PayloadMessage = new(transporter.PayloadMessage)
	response.Data = "pong"

	return response, nil
}
