package lcd

import (
	"encoding/json"
	"errors"
	"fmt"

	ipaltypes "github.com/netcloth/netcloth-chain/modules/ipal/types"
	sdk "github.com/netcloth/netcloth-chain/types"

	"github.com/netcloth/go-sdk/client/types"
	"github.com/netcloth/go-sdk/keys"
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

	IPALsBody struct {
		Height string       `json:"height"`
		Result []IPALResult `json:"result"`
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
		if endpoint.Type == EndpointTypeChat {
			return endpoint.Endpoint, nil
			break
		}
	}
	return "", errors.New("no chat endpoint")
}

func QueryServiceNodesParamsFromBech32Addresses(addresses []string) (params ipaltypes.QueryServiceNodesParams) {
	for _, address := range addresses {
		addr, err := sdk.AccAddressFromBech32(address)
		if err != nil {
			continue
		}
		params.AccAddrs = append(params.AccAddrs, addr)
	}

	return
}

func (c *client) QueryIPALChatServersEndpointByAddresses(addresses []string) (map[string]string, error) {
	var r IPALsBody

	params := QueryServiceNodesParamsFromBech32Addresses(addresses)
	if _, body, err := c.httpClient.Post(UriQueryIPALs, nil, params); err != nil {
		return nil, err
	} else {
		if err := json.Unmarshal(body, &r); err != nil {
			return nil, err
		} else {
			result := make(map[string]string)
			for _, sn := range r.Result {
				for _, ep := range sn.Endpoints {
					if ep.Type == EndpointTypeChat {
						result[sn.OperatorAddress] = ep.Endpoint
					}
				}
			}
			return result, nil
		}
	}
}

func (c *client) QueryIPALsEndpointByAddressesByType(addresses []string, endpointType string) (map[string]string, error) {
	var r IPALsBody

	params := QueryServiceNodesParamsFromBech32Addresses(addresses)
	if _, body, err := c.httpClient.Post(UriQueryIPALs, nil, params); err != nil {
		return nil, err
	} else {
		if err := json.Unmarshal(body, &r); err != nil {
			return nil, err
		} else {
			result := make(map[string]string)
			for _, sn := range r.Result {
				for _, ep := range sn.Endpoints {
					if ep.Type == endpointType {
						result[sn.OperatorAddress] = ep.Endpoint
					}
				}
			}
			return result, nil
		}
	}
}
