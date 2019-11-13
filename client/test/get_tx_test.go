package test

import (
	"testing"

	"github.com/netcloth/go-sdk/util"
)

func TestClient_GetTx(t *testing.T) {
	hash := "06CA852A1D6401BE7BE6EE3D402E2FF2B4432A4FDDCFD599535FB9BDFD5CED4E"
	if res, err := c.GetTx(hash); err != nil {
		t.Fatal(err)
	} else {
		t.Log(util.ToJsonIgnoreErr(res))
	}
}
