package client

import (
	"testing"

	"github.com/NetCloth/go-sdk/keys"
	"github.com/NetCloth/go-sdk/types"
	"github.com/NetCloth/go-sdk/util"
)

const (
	AccAddr = "nch1skhg3tjm09dzcy2zn673006z4c3p47rg929se9"
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
		if res, err := c.QueryAccount(AccAddr); err != nil {
			t.Fatal(err)
		} else {
			t.Log(util.ToJsonIgnoreErr(res))
		}

		// query ipal
		if res, err := c.QueryCIPALByAddress(AccAddr); err != nil {
			t.Fatal(err)
		} else {
			t.Log(util.ToJsonIgnoreErr(res))
		}

		//query aipal
		if res, err := c.QueryAIPALByAddress(AccAddr); err != nil {
			t.Fatal(err)
		} else {
			t.Log(util.ToJsonIgnoreErr(res))
		}

		//query aipallist
		if res, err := c.QueryAIPALList(); err != nil {
			t.Fatal(err)
		} else {
			t.Log(util.ToJsonIgnoreErr(res))
		}
	}
}
