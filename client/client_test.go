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
		if res, err := c.QueryAccount("nch1dtpryue8ptzjjm32fwr0a7u5qg6wz02hhnpa30"); err != nil {
			t.Fatal(err)
		} else {
			t.Log(util.ToJsonIgnoreErr(res))
		}

		// query ipal
		if res, err := c.QueryCIPALByAddress("nch1jx2jcycf86vll2yfqttrj85ukws34xjhn8ef4q"); err != nil {
			t.Fatal(err)
		} else {
			t.Log(util.ToJsonIgnoreErr(res))
		}

		//query aipal
		if res, err := c.QueryAIPALByAddress("nch1fxs3zym0tk0gva9mcwdghcwh6d96426d8m4lts"); err != nil {
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
