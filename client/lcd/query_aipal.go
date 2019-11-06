package lcd

import (
	"encoding/json"
	"fmt"

	"github.com/NetCloth/go-sdk/client/types"
)

type (
	Endpoint struct {
		Type     string `json:"type"`
		Endpoint string `json:"endpoint"`
	}

	AIPALResult struct {
		OperatorAddress string     `json:"operator_address"`
		Moniker         string     `json:"moniker"`
		Website         string     `json:"website"`
		Details         string     `json:"details"`
		Endpoints       []Endpoint `json:"endpoints"`
		Bond            types.Coin `json:"bond"`
	}

	AIPALBody struct {
		Height string      `json:"height"`
		Resutl AIPALResult `json:"result"`
	}
)

func (c *client) QueryAIPALByAddress(address string) (AIPALBody, error) {
	var r AIPALBody

	if _, body, err := c.httpClient.Get(fmt.Sprintf(UriQueryAIPAL, address), nil); err != nil {
		return r, err
	} else {
		if err := json.Unmarshal(body, &r); err != nil {
			return r, err
		} else {
			return r, nil
		}
	}
}
