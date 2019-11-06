package client

import (
	"testing"

	"github.com/NetCloth/go-sdk/client/tx"

	"github.com/NetCloth/go-sdk/client/basic"
	"github.com/NetCloth/go-sdk/client/lcd"
	"github.com/NetCloth/go-sdk/client/rpc"
	"github.com/NetCloth/go-sdk/client/types"
	"github.com/NetCloth/go-sdk/keys"
	commontypes "github.com/NetCloth/go-sdk/types"
	"github.com/NetCloth/go-sdk/util"
)

func TestClient_SendToken(t *testing.T) {
	km, err := keys.NewKeystoreByImportKeystore("./ks_12345678.txt", "12345678")
	if err != nil {
		panic(err)
	}

	basicClient := basic.NewClient("http://127.0.0.1:1317")
	lite := lcd.NewClient(basicClient)
	rpcClient := rpc.NewClient("tcp://127.0.0.1:26657")

	c, err := tx.NewClient("nch-prinet-sky", commontypes.Alphanet, km, lite, rpcClient)

	coins := []types.Coin{
		{
			Denom:  "unch",
			Amount: "100",
		},
	}
	if res, err := c.SendToken(AccAddr, coins, "", false); err != nil {
		t.Fatal(err)
	} else {
		t.Log(util.ToJsonIgnoreErr(res))
	}
}
