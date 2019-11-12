package lcd

import (
	"encoding/json"
	"fmt"

	"github.com/netcloth/go-sdk/client/types"
)

type (
	AccountValue struct {
		Address       string       `json:"address"`
		Coins         []types.Coin `json:"coins"`
		AccountNumber string       `json:"account_number"`
		Sequence      string       `json:"sequence"`
	}

	AccountResult struct {
		Type  string       `json:"type"`
		Value AccountValue `json:"value"`
	}

	AccountBody struct {
		Height string        `json:"height"`
		Result AccountResult `json:"result"`
	}
)

func (c *client) QueryAccount(address string) (accountInfo AccountBody, err error) {
	path := fmt.Sprintf(UriQueryAccount, address)

	if _, body, err := c.httpClient.Get(path, nil); err != nil {
		return accountInfo, err
	} else {
		if err := json.Unmarshal(body, &accountInfo); err != nil {
			return accountInfo, err
		} else {
			return accountInfo, nil
		}
	}
}
