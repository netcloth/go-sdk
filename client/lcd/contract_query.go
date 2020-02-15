package lcd

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type (
	ContractLog struct {
		Height string `json:"height"`
		Result VMLogs `json:"result"`
	}

	VMLogs struct {
		Logs []VMLog `json:"logs"`
	}

	VMLog struct {
		Address          string   `json:"address"`
		Topics           []string `json:"topics"`
		Data             string   `json:"data"`
		BlockNumber      string   `json:"blockNumber"`
		TransactionHash  string   `json:"transactionHash"`
		TransactionIndex string   `json:"transactionIndex"`
		BlockHash        string   `json:"blockHash"`
		LogIndex         string   `json:"logIndex"`
		Removed          bool     `json:"removed"`
	}
)

func (c *client) QueryContractLog(txId []byte) (ContractLog, error) {
	var res ContractLog

	if _, body, err := c.httpClient.Get(fmt.Sprintf(UriQueryContractLogs, hex.EncodeToString(txId)), nil); err != nil {
		return res, err
	} else {
		if err := json.Unmarshal(body, &res); err != nil {
			return res, err
		} else {
			return res, nil
		}
	}
}
