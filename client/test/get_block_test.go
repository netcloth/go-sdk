package test

import (
	"testing"

	"github.com/netcloth/go-sdk/util"
)

func TestClient_GetBlock(t *testing.T) {
	const height int64 = 25
	if res, err := c.Block(height); err != nil {
		t.Fatal(err)
	} else {
		t.Log(util.ToJsonIgnoreErr(res))
	}
}
