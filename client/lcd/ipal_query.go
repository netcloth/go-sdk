package lcd

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/netcloth/go-sdk/keys"

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

func (c *client) QueryIPALByUNCompressedPubKey(uncompressedPubKey string) (IPALBody, error) {
	var r IPALBody

	addrBech32, err := keys.UNCompressedPubKey2AddressBech32(uncompressedPubKey)
	if err != nil {
		return r, err
	}

	return c.QueryIPALByAddress(addrBech32)
}

func (c *client) QueryIPALChatServerEndpointByUNCompressedPubKey(uncompressedPubKey string) (string, error) {
	ipalInfo, err := c.QueryIPALByUNCompressedPubKey(uncompressedPubKey)
	if err != nil {
		return "", err
	}

	for _, endpoint := range ipalInfo.Result.Endpoints {
		if endpoint.Type == "1" { //TODO remove magic number "1"
			return endpoint.Endpoint, nil
			break
		}
	}
	return "", errors.New("no chat endpoint")
}
