package rpc

import (
	"testing"

	"github.com/NetCloth/go-sdk/util"
)

var (
	c RPCClient
)

func TestMain(m *testing.M) {
	c = NewClient("tcp://127.0.0.1:26657")
	m.Run()
}

func TestClient_GetStatus(t *testing.T) {
	if res, err := c.GetStatus(); err != nil {
		t.Fatal(err)
	} else {
		t.Log(util.ToJsonIgnoreErr(res))
	}
}
