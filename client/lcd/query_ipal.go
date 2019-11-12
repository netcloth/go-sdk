package lcd

import (
	"encoding/json"
	"fmt"

	"github.com/netcloth/go-sdk/client/types"
)

type (
	Endpoint struct {
		Type     string `json:"type"`
		Endpoint string `json:"endpoint"`
	}

	IPALResult struct {
		OperatorAddress string     `json:"operator_address"`
		Moniker         string     `json:"moniker"`
		Website         string     `json:"website"`
		Details         string     `json:"details"`
		Endpoints       []Endpoint `json:"endpoints"`
		Bond            types.Coin `json:"bond"`
	}

	IPALBody struct {
		Height string     `json:"height"`
		Result IPALResult `json:"result"`
	}
)

func (c *client) QueryIPALByAddress(address string) (IPALBody, error) {
	var r IPALBody

	if _, body, err := c.httpClient.Get(fmt.Sprintf(UriQueryIPAL, address), nil); err != nil {
		return r, err
	} else {
		if err := json.Unmarshal(body, &r); err != nil {
			return r, err
		} else {
			return r, nil
		}
	}
}
