package auth

import (
	"fmt"
	"net/http"

	"github.com/block-api/block-node/block"
	"github.com/block-api/block-node/common/types"
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

func (ab *AuthBlock) HttpGetHello(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	payload := transporter.PayloadMessage{
		Data: "it is working",
	}

	// this action will be executed locally - even if there are other "block-node-example" nodes available in the network
	// in this case there would be no request sent over network, it is done that way to reduce latency
	// "v1.block-node-example.auth.authorization"
	targetAction := types.TargetAction{
		Name:    "block-node-example",
		Version: 1,
		Block:   "auth",
		Action:  "authorization",
	}

	pocket := transporter.NewPocket(transporter.ChanMessage, ab.BlockNode().VersionName(), ab.BlockNode().NodeID(), nil, &targetAction, payload)
	ab.BlockNode().Send(pocket)

	fmt.Fprintf(w, "hello\n")
}
