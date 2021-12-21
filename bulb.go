package lightbrite

import (
	"encoding/json"
	"log"
	"net"
)

type Bulb struct {
	IP    net.IP
	Port  int
	State map[string]string
	Mac   string
}

type PilotRequest struct {
	Method string `json:"method"`
	// Id     string      `json:"id"`
	Params PilotParams `json:"params"`
}

type PilotParams map[string]interface{}

func NewPilotRequest(method string, params PilotParams) PilotRequest {
	return PilotRequest{
		Method: method,
		Params: params,
	}
}

func PilotRequestFromJSON(inp []byte) PilotRequest {
	var out PilotRequest
	err := json.Unmarshal(inp, &out)
	if err != nil {
		return PilotRequest{}
	}

	return out
}

func (pr PilotRequest) toJSON() []byte {
	out, err := json.Marshal(pr)
	if err != nil {
		return []byte("")
	}

	return out
}

type PilotResponse struct {
	Result map[string]interface{} `json:"result"`
	Env    string                 `json:"env"`
	Method string                 `json:"method"`
}

func PilotResponseFromJSON(inp []byte) PilotResponse {
	var out PilotResponse
	log.Print(string(inp))
	err := json.Unmarshal(inp, &out)
	if err != nil {
		return PilotResponse{}
	}

	return out
}

func (bl Bulb) SetDataFromResponse(resp PilotResponse) {
	for kx := range resp.Result {
		switch kx {
		default:
			log.Printf("Data: %v\n", resp)
		}
	}
}
