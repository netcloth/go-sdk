package client

import (
	"testing"

	"github.com/NetCloth/go-sdk/keys"
	"github.com/NetCloth/go-sdk/types"
	"github.com/NetCloth/go-sdk/util"
)

var (
	baseUrl     = "http://127.0.0.1:1317"
	nodeUrl     = "tcp://127.0.0.1:26657"
	networkType = types.Alphanet
	km          keys.KeyManager
)

func TestMain(m *testing.M) {
	if k, err := keys.NewKeyStoreKeyManager("../keys/ks_1234567890.json", "1234567890"); err != nil {
		panic(err)
	} else {
		km = k
	}
	m.Run()
}

func TestNewNCHClient(t *testing.T) {
	c, err := NewNCHClient(baseUrl, nodeUrl, networkType, km)
	if err != nil {
		t.Fatal(err)
	} else {
		// query account
		if res, err := c.QueryAccount("nch1ue8f3uzdwv9mhyvppucepz6zwpquqpzz9untnm"); err != nil {
			t.Fatal(err)
		} else {
			t.Log(util.ToJsonIgnoreErr(res))
		}

		// query ipal
		if res, err := c.QueryIPAL("nch18q4pv9qvmqx7dcd2jq3dl3d0755urk835w2ak4"); err != nil {
			t.Fatal(err)
		} else {
			t.Log(util.ToJsonIgnoreErr(res))
		}
	}
}
