package lcd

import (
	"encoding/json"
	"fmt"
)

type (
	VMLogs struct {
		Logs []VMLog `json:"logs"`
	}

	VMLog struct {
		address          string   `json:"address"`
		topics           []string `json:"topics"`
		data             string   `json:"data"`
		blockNumber      uint64   `json:"blockNumber"`
		transactionHash  string   `json:"transactionHash"`
		transactionIndex uint64   `json:"transactionIndex"`
		blockHash        string   `json:"blockHash"`
		logIndex         uint64   `json:"logIndex"`
		removed          bool     `json:"removed"`
	}
)

func (l *VMLog) String() string {
	return l.data
}

func (c *client) QueryContractLog(txId string) (VMLogs, error) {
	var r VMLogs

	if _, body, err := c.httpClient.Get(fmt.Sprintf(UriQueryContractLogs, txId), nil); err != nil {
		return r, err
	} else {
		if err := json.Unmarshal(body, &r); err != nil {
			return r, err
		} else {
			return r, nil
		}
	}
}
