package client

import (
	"encoding/hex"
	"testing"
	"time"

	"github.com/tendermint/tendermint/crypto/secp256k1"

	"github.com/netcloth/go-sdk/client/basic"
	"github.com/netcloth/go-sdk/client/lcd"
	"github.com/netcloth/go-sdk/client/rpc"
	"github.com/netcloth/go-sdk/client/tx"
	"github.com/netcloth/go-sdk/keys"
	"github.com/netcloth/go-sdk/types"
	"github.com/netcloth/go-sdk/util"
	"github.com/netcloth/netcloth-chain/modules/auth"
	"github.com/netcloth/netcloth-chain/modules/cipal"
	"github.com/netcloth/netcloth-chain/modules/ipal"
	sdk "github.com/netcloth/netcloth-chain/types"
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

		//demo for GetBech32AddrByPubkeyStr and GetBech32AddrByPubkey
		pubkeyStr := "020268AA87DA53D0667FF931E741635E1409CB2E105D409B3C6253E13FF57BDEDC"
		t.Log(keys.GetBech32AddrByPubKeyStr(pubkeyStr))

		var pubkey secp256k1.PubKeySecp256k1
		pubkeyHex, err := hex.DecodeString(pubkeyStr)
		if err != nil {
			return
		}
		copy(pubkey[:], pubkeyHex)
		t.Log(keys.GetBech32AddrByPubKey(pubkey))
	}
}

func Test_IPALClaim(t *testing.T) {
	km, err := keys.NewKeystoreByImportKeystore("../keys/ks_12345678.txt", "12345678")
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
	km, err := keys.NewKeystoreByImportKeystore("../keys/ks_12345678.txt", "12345678")
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

func Test_CIPALClaim1(t *testing.T) {
	km, err := keys.NewKeystoreByImportKeystore("../keys/ks_12345678.txt", "12345678")
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

	msg := "{\"expriration\":\"2019-11-08T20:18:44.524268Z\",\"service_info\":{\"address\":\"nch19vnsnnseazkyuxgkt0098gqgvfx0wxmv96479m\",\"type\":\"1\"},\"user_address\":\"nch1edyenjf04mrsq3ueghmmga35hjeetgudjp4z36\"}"
	msgbz := []byte(msg)

	sigstr := "8271c4c00774e3de49f468367e610caf07c1319a1f0b8724f427ab1e918b703d429834765e8c9e8b7333919ced0ff68c28fea433e557816e06815447106d68be"
	sigbz, err := hex.DecodeString(sigstr)
	if err != nil {
		t.Log(err)
		return
	}

	pubkeystr := "028c36b163c26f492abc874648b7258450394fe78133bcc4d920895d0ce8c3ac4e"
	pubkeybz, err := hex.DecodeString(pubkeystr)
	if err != nil {
		t.Log(err)
		return
	}

	var pubkey secp256k1.PubKeySecp256k1
	copy(pubkey[:], pubkeybz)
	t.Log(pubkey)
	t.Log(pubkey.VerifyBytes(msgbz, sigbz))

	req := cipal.NewIPALUserRequest(AccAddr, AccAddr, 6, expiration, stdSig)
	if res, err := c.CIPALClaim(req, "memo", false); err != nil {
		t.Fatal(err)
	} else {
		t.Log(util.ToJsonIgnoreErr(res))
	}
}
