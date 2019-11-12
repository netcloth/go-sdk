package lcd

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/netcloth/go-sdk/keys"
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
		if serviceInfo.Type == "1" { //TODO remove magic number "1"
			return serviceInfo.Address, nil
			break
		}
	}

	return "", errors.New("no chat endpoint")
}
