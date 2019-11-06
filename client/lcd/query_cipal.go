package lcd

import (
	"encoding/json"
	"fmt"
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
		Resutl CIPALResult `json:"result"`
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
