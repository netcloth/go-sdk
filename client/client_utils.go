package client

import (
	"github.com/netcloth/go-sdk/client/basic"
	"github.com/netcloth/go-sdk/client/lcd"
	"github.com/netcloth/go-sdk/client/rpc"
	"github.com/netcloth/go-sdk/client/tx"
	"github.com/netcloth/go-sdk/keys"
	"github.com/netcloth/go-sdk/util/constant"
)

func NewClient() {
	km, err := keys.NewKeyManager(constant.KeyStoreFileAbsPath, "12345678")
	if err != nil {
		panic(err)
	}

	basicClient := basic.NewClient("http://127.0.0.1:1317")
	lite := lcd.NewClient(basicClient)
	rpcClient := rpc.NewClient("tcp://127.0.0.1:26657")

	c, err := tx.NewClient("nch-prinet-sky", commontypes.Alphanet, km, lite, rpcClient)
}
