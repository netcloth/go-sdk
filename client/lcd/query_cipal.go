package lcd

import (
	"encoding/json"
	"fmt"
)

type (
	CIPALResult struct {
		UserAddress string `json:"user_address"`
		ServerIp    string `json:"server_ip"`
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
