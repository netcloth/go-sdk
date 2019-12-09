package lcd

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/netcloth/go-sdk/keys"
	"github.com/netcloth/netcloth-chain/modules/cipal/types"
)

type (
	CIPALServiceInfo struct {
		Type    string `json:"type"`
		Address string `json:"address"`
	}

	CIPALResult struct {
		UserAddress  string             `json:"user_address"`
		ServiceInfos []CIPALServiceInfo `json:"service_infos"`
	}

	CIPALBody struct {
		Height string      `json:"height"`
		Result CIPALResult `json:"result"`
	}

	CIPALsBody struct {
		Height string        `json:"height"`
		Result []CIPALResult `json:"result"`
	}
)

func (c *client) QueryCIPALByAddress(address string) (CIPALBody, error) {
	var r CIPALBody

	if _, body, err := c.httpClient.Get(fmt.Sprintf(UriQueryCIPAL, address), nil); err != nil {
		return r, err
	} else {
		if err := json.Unmarshal(body, &r); err != nil {
			return r, err
		} else {
			return r, nil
		}
	}
}

func (c *client) QueryCIPALByUNCompressedPubKey(uncompressedPubKey string) (CIPALBody, error) {
	var r CIPALBody

	addrBech32, err := keys.UNCompressedPubKey2AddressBech32(uncompressedPubKey)
	if err != nil {
		return r, err
	}

	return c.QueryCIPALByAddress(addrBech32)
}

func (c *client) QueryCIPALChatServerAddrByUNCompressedPubKey(uncompressedPubKey string) (string, error) {
	cipalInfo, err := c.QueryCIPALByUNCompressedPubKey(uncompressedPubKey)
	if err != nil {
		return "", err
	}

	for _, serviceInfo := range cipalInfo.Result.ServiceInfos {
		if ClientChatEndpointType2ServerChatEndpointType[serviceInfo.Type] == EndpointTypeServerChat {
			return serviceInfo.Address, nil
			break
		}
	}

	return "", errors.New("no chat endpoint")
}

func Bech32AddrsFromUNCompressedPubKeys(uncompressedPubKeys []string) (bech32Addrs []string) {
	for _, pubKey := range uncompressedPubKeys {
		bech32Addr, err := keys.UNCompressedPubKey2AddressBech32(pubKey)
		if err != nil {
			bech32Addrs = append(bech32Addrs, bech32Addr)
		}
	}

	return
}

func QueryCIPALsParamsFromUNCompressedPubKeys(uncompressedPubKeys []string) (params types.QueryCIPALsParams) {
	params.AccAddrs = Bech32AddrsFromUNCompressedPubKeys(uncompressedPubKeys)
	return
}

func (c *client) QueryCIPALChatServersAddrByUNCompressedPubKeys(uncompressedPubKeys []string) (map[string]string, error) {
	var r CIPALsBody

	params := QueryCIPALsParamsFromUNCompressedPubKeys(uncompressedPubKeys)
	if _, body, err := c.httpClient.Post(UriQueryCIPALs, nil, params); err != nil {
		return nil, err
	} else {
		if err := json.Unmarshal(body, &r); err != nil {
			return nil, err
		} else {
			result := make(map[string]string)
			for _, cipal := range r.Result {
				for _, si := range cipal.ServiceInfos {
					if ClientChatEndpointType2ServerChatEndpointType[si.Type] == EndpointTypeServerChat {
						result[cipal.UserAddress] = si.Address
						break
					}
				}
			}
			return result, nil
		}
	}
}

func (c *client) QueryCIPALsAddrByUNCompressedPubKeysByType(uncompressedPubKeys []string, endpointType string) (map[string]string, error) {
	var r CIPALsBody

	params := QueryCIPALsParamsFromUNCompressedPubKeys(uncompressedPubKeys)
	if _, body, err := c.httpClient.Post(UriQueryCIPALs, nil, params); err != nil {
		return nil, err
	} else {
		if err := json.Unmarshal(body, &r); err != nil {
			return nil, err
		} else {
			result := make(map[string]string)
			for _, cipal := range r.Result {
				for _, si := range cipal.ServiceInfos {
					if ClientChatEndpointType2ServerChatEndpointType[si.Type] == ClientChatEndpointType2ServerChatEndpointType[endpointType] {
						result[cipal.UserAddress] = si.Address
						break
					}
				}
			}
			return result, nil
		}
	}
}
