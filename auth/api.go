package auth

import (
	"net/http"

	"github.com/block-api/block-node/common/types"
	"github.com/block-api/block-node/transporter"
)

func (ab *AuthBlock) ApiHello(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	// this action will be executed locally - even if there are other "block-node-example" nodes available in the network
	// in this case there would be no request sent over network, it is done that way to reduce latency
	// "v1.block-node-example.auth.authorization"
	targetAction := types.TargetAction{
		Name:    "block-node-example",
		Version: 1,
		Block:   "auth",
		Action:  "authorization",
	}

	payload := transporter.PayloadMessage{
		Data: "it is working",
	}

	ab.BlockNode().Send(&payload, &targetAction)

	w.WriteHeader(204)
}
