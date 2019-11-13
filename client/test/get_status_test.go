package test

import (
	"testing"

	"github.com/netcloth/go-sdk/client/rpc"
	"github.com/netcloth/go-sdk/util"
)

var (
	c rpc.RPCClient
)

func TestMain(m *testing.M) {
	c = rpc.NewClient("tcp://127.0.0.1:26657")
	m.Run()
}

func TestClient_GetStatus(t *testing.T) {
	if res, err := c.GetStatus(); err != nil {
		t.Fatal(err)
	} else {
		t.Log(util.ToJsonIgnoreErr(res))
	}
}
