package test

import (
	"testing"
	"time"

	"github.com/netcloth/go-sdk/util/constant"

	"github.com/netcloth/netcloth-chain/modules/auth"
	"github.com/netcloth/netcloth-chain/modules/cipal"
	"github.com/netcloth/netcloth-chain/modules/ipal"
	sdk "github.com/netcloth/netcloth-chain/types"

	"github.com/netcloth/go-sdk/client/basic"
	"github.com/netcloth/go-sdk/client/lcd"
	"github.com/netcloth/go-sdk/client/rpc"
	"github.com/netcloth/go-sdk/client/tx"
	"github.com/netcloth/go-sdk/keys"
	"github.com/netcloth/go-sdk/util"
)

const (
	AccAddr = "nch1ugus2df3sydca3quula5yjqfntuq5aaxweezpt"
)

func Test_IPALClaim(t *testing.T) {
	km, err := keys.NewKeyManager(constant.KeyStoreFileAbsPath, "12345678")
	if err != nil {
		panic(err)
	}

	basicClient := basic.NewClient("http://127.0.0.1:1317")
	lite := lcd.NewClient(basicClient)
	rpcClient := rpc.NewClient("tcp://127.0.0.1:26657")
	c, err := tx.NewClient("nch-prinet-sky", 1, km, lite, rpcClient)

	bond := sdk.Coin{
		Denom:  "unch",
		Amount: sdk.NewInt(1000000),
	}

	var eps ipal.Endpoints
	ep := ipal.NewEndpoint(10, "192.168.100.100:20000")
	eps = append(eps, ep)
	if res, err := c.IPALClaim("sky", "sky weibsite", "sky details", eps, bond, false); err != nil {
		t.Fatal(err)
	} else {
		t.Log(util.ToJsonIgnoreErr(res))
	}
}

func Test_CIPALClaim(t *testing.T) {
	km, err := keys.NewKeyManager(constant.KeyStoreFileAbsPath, "12345678")
	if err != nil {
		panic(err)
	}

	basicClient := basic.NewClient("http://127.0.0.1:1317")
	lite := lcd.NewClient(basicClient)
	rpcClient := rpc.NewClient("tcp://127.0.0.1:26657")
	c, err := tx.NewClient("nch-prinet-sky", 1, km, lite, rpcClient)
	if err != nil {
		panic(err)
	}

	expiration := time.Now().UTC().AddDate(0, 0, 1)
	adMsg := cipal.NewADParam(AccAddr, AccAddr, 6, expiration)
	sigBytes, err := km.SignBytes(adMsg.GetSignBytes())
	if err != nil {
		panic(err)
	}

	stdSig := auth.StdSignature{
		PubKey:    km.GetPrivKey().PubKey(),
		Signature: sigBytes,
	}

	req := cipal.NewIPALUserRequest(AccAddr, AccAddr, 6, expiration, stdSig)
	if res, err := c.CIPALClaim(req, "memo", false); err != nil {
		t.Fatal(err)
	} else {
		t.Log(util.ToJsonIgnoreErr(res))
	}
}
