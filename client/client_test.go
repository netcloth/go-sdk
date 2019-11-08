package client

import (
	"encoding/hex"
	"testing"
	"time"

	"github.com/tendermint/tendermint/crypto/secp256k1"

	"github.com/NetCloth/netcloth-chain/modules/cipal"

	"github.com/NetCloth/netcloth-chain/modules/auth"
	"github.com/NetCloth/netcloth-chain/modules/ipal"

	sdk "github.com/NetCloth/netcloth-chain/types"

	"github.com/NetCloth/go-sdk/client/basic"
	"github.com/NetCloth/go-sdk/client/lcd"
	"github.com/NetCloth/go-sdk/client/rpc"
	"github.com/NetCloth/go-sdk/client/tx"
	"github.com/NetCloth/go-sdk/keys"
	"github.com/NetCloth/go-sdk/types"
	"github.com/NetCloth/go-sdk/util"
)

const (
	AccAddr = "nch1ugus2df3sydca3quula5yjqfntuq5aaxweezpt"
)

var (
	baseUrl     = "http://127.0.0.1:1317"
	nodeUrl     = "tcp://127.0.0.1:26657"
	networkType = types.Alphanet
	km          keys.KeyManager
)

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

		pubkeyStr := "020268AA87DA53D0667FF931E741635E1409CB2E105D409B3C6253E13FF57BDEDC"
		t.Log(keys.GetBech32AddrByPubkeyStr(pubkeyStr))

		var pubkey secp256k1.PubKeySecp256k1
		pubkeyHex, err := hex.DecodeString(pubkeyStr)
		if err != nil {
			return
		}
		copy(pubkey[:], pubkeyHex)
		t.Log(keys.GetBech32AddrByPubkey(pubkey))
	}
}

func Test_IPALClaim(t *testing.T) {
	km, err := keys.NewKeystoreByImportKeystore("./ks_12345678.txt", "12345678")
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
	km, err := keys.NewKeystoreByImportKeystore("./ks_12345678.txt", "12345678")
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
