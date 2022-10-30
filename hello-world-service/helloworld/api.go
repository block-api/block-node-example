package helloworld

import (
	"encoding/json"
	"net/http"

	"github.com/block-api/block-node/common/types"
	"github.com/block-api/block-node/errors"
	"github.com/block-api/block-node/log"
	"github.com/block-api/block-node/transporter"
)

func (ab *HelloWorldBlock) ApiHello(w http.ResponseWriter, req *http.Request) {
	var response map[string]string = make(map[string]string)

	w.Header().Add("Content-Type", "application/json")

	urlQuery := req.URL.Query()
	name := urlQuery.Get("name")

	if name == "" {
		response["error"] = "name is missing"

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Warning(err.Error())
			w.WriteHeader(500)

			return
		}

		w.WriteHeader(400)
		w.Write(jsonResponse)
		return
	}

	// this action will be executed locally - even if there are other "block-node-example" nodes available in the network
	// in this case there would be no request sent over network, it is done that way to reduce latency
	// "v1.block-node-example.auth.hello"
	targetAction := types.TargetAction{
		Name:    "hello-world-service",
		Version: 1,
		Block:   "helloworld",
		Action:  "hello",
	}

	payload := transporter.PayloadMessage{
		Data: name,
	}

	resPayload, err := ab.BlockNode().Send(&payload, &targetAction)
	if err != nil {
		log.Warning(err.Error())

		if err == errors.ErrInvalidTargetAction {
			response["error"] = errors.ErrInvalidTargetAction.Error()
			w.WriteHeader(400)
		}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Warning(err.Error())
			w.WriteHeader(500)

			return
		}

		w.Write(jsonResponse)

		return
	}

	json, _ := resPayload.JSON()

	w.WriteHeader(200)
	w.Write(json)
}

func (ab *HelloWorldBlock) ApiPing(w http.ResponseWriter, req *http.Request) {
	var response map[string]string = make(map[string]string)

	w.Header().Add("Content-Type", "application/json")

	targetAction := types.TargetAction{
		Name:    "node-test",
		Version: 1,
		Block:   "pingpong",
		Action:  "ping",
	}

	resPayload, err := ab.BlockNode().Send(nil, &targetAction)
	if err != nil {
		log.Warning(err.Error())

		if err == errors.ErrInvalidTargetAction {
			response["error"] = errors.ErrInvalidTargetAction.Error()
			w.WriteHeader(400)
		}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Warning(err.Error())
			w.WriteHeader(500)

			return
		}

		w.Write(jsonResponse)

		return
	}

	json, _ := resPayload.JSON()

	w.WriteHeader(200)
	w.Write(json)
}
