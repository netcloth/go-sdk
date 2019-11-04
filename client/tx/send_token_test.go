package tx

import (
	"fmt"
	"math"
	"testing"

	"github.com/NetCloth/go-sdk/client/basic"
	"github.com/NetCloth/go-sdk/client/lcd"
	"github.com/NetCloth/go-sdk/client/rpc"
	"github.com/NetCloth/go-sdk/client/types"
	"github.com/NetCloth/go-sdk/keys"
	commontypes "github.com/NetCloth/go-sdk/types"
	"github.com/NetCloth/go-sdk/util"
)

var (
	c TxClient
)

func TestMain(m *testing.M) {
	km, err := keys.NewKeyStoreKeyManager("./ks_1234567890.json", "1234567890")
	if err != nil {
		panic(err)
	}
	basicClient := basic.NewClient("http://127.0.0.1:1317")
	lite := lcd.NewClient(basicClient)
	rpcClient := rpc.NewClient("tcp://127.0.0.1:25567")

	c, err = NewClient("nch-", commontypes.Alphanet, km, lite, rpcClient)
	if err != nil {
		panic(err)
	}
	m.Run()
}

func TestClient_SendToken(t *testing.T) {
	receiver := "nch1p3fuppcxud5rjsaywuyuguh6achmj5p0r6z6ve"
	amount := fmt.Sprintf("%.0f", 0.12*math.Pow10(18))
	coins := []types.Coin{
		{
			Denom:  "unch",
			Amount: amount,
		},
	}
	memo := "send from NetCloth/go-sdk"
	if res, err := c.SendToken(receiver, coins, memo, false); err != nil {
		fmt.Println(res)
		t.Fatal(err)
	} else {
		t.Log(util.ToJsonIgnoreErr(res))
	}
}
