package client

import (
	"encoding/hex"
	"testing"

	"github.com/tendermint/tendermint/crypto/secp256k1"

	"github.com/netcloth/go-sdk/keys"
	"github.com/netcloth/go-sdk/types"
	"github.com/netcloth/go-sdk/util"
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

		//demo for GetBech32AddrByPubkeyStr and GetBech32AddrByPubkey
		pubkeyStr := "020268AA87DA53D0667FF931E741635E1409CB2E105D409B3C6253E13FF57BDEDC"
		t.Log(keys.PubKeyHexString2AddressBech32(pubkeyStr))

		var pubkey secp256k1.PubKeySecp256k1
		pubkeyHex, err := hex.DecodeString(pubkeyStr)
		if err != nil {
			return
		}
		copy(pubkey[:], pubkeyHex)
		t.Log(keys.PubKey2AddressBech32(pubkey))
	}
}
