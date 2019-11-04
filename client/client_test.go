package client

import (
	"testing"

	"github.com/NetCloth/go-sdk/keys"
	"github.com/NetCloth/go-sdk/types"
	"github.com/NetCloth/go-sdk/util"
)

var (
	baseUrl     = "http://xxxxx"
	nodeUrl     = "tcp://127.0.0.1:25557"
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
		if res, err := c.QueryAccount("faa1282eufkw9qgm55symgqqg38nremslvggpylkht"); err != nil {
			t.Fatal(err)
		} else {
			t.Log(util.ToJsonIgnoreErr(res))
		}
	}
}
