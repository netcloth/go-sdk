package tx

import (
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
//c TxClient
)

func TestMain(m *testing.M) {
	km, err := keys.NewKeystoreByImportKeystore("./ks_12345678.txt", "12345678")
	if err != nil {
		panic(err)
	}
	basicClient := basic.NewClient("http://127.0.0.1:1317")
	lite := lcd.NewClient(basicClient)
	rpcClient := rpc.NewClient("tcp://127.0.0.1:26657")

	_, err = NewClient("nch-prinet-sky", commontypes.Alphanet, km, lite, rpcClient)
	if err != nil {
		panic(err)
	}
	m.Run()
}

func TestClient_SendToken(t *testing.T) {

	km, err := keys.NewKeystoreByImportKeystore("./ks_12345678.txt", "12345678")
	if err != nil {
		panic(err)
	}
	basicClient := basic.NewClient("http://127.0.0.1:1317")
	lite := lcd.NewClient(basicClient)
	rpcClient := rpc.NewClient("tcp://127.0.0.1:26657")

	c, err := NewClient("nch-prinet-sky", commontypes.Alphanet, km, lite, rpcClient)

	receiver := "nch1dtpryue8ptzjjm32fwr0a7u5qg6wz02hhnpa30"
	//amount := fmt.Sprintf("%.0f", 0.12*math.Pow10(18))
	coins := []types.Coin{
		{
			Denom:  "unch",
			Amount: "100",
		},
	}
	//memo := "send from NetCloth/go-sdk"
	if res, err := c.SendToken(receiver, coins, "", false); err != nil {
		t.Fatal(err)
	} else {
		t.Log(util.ToJsonIgnoreErr(res))
	}
}
