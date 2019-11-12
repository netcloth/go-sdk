package keys

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	sdk "github.com/netcloth/netcloth-chain/types"
)

func TestNewKeyStoreKeyManager(t *testing.T) {
	file := "./ks_1234567890.json"
	if km, err := NewKeyStoreKeyManager(file, "1234567890"); err != nil {
		t.Fatal(err)
	} else {
		msg := []byte("hello world")
		signature, err := km.GetPrivKey().Sign(msg)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(km.GetAddr().String())

		assert.Equal(t, km.GetPrivKey().PubKey().VerifyBytes(msg, signature), true)
	}
}

func Test_ImportKeystore(t *testing.T) {
	file := "./ks_12345678.txt"
	if km, err := NewKeystoreByImportKeystore(file, "12345678"); err != nil {
		t.Fatal(err)
	} else {
		msg := []byte("hello world")
		signature, err := km.GetPrivKey().Sign(msg)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(fmt.Sprintf("prikey hex:%X", km.GetPrivKey().PubKey().Bytes()))
		pubkey, err := sdk.GetAccPubKeyBech32("nchpub1addwnpepq2etlx58m470rtvesusllmcqwyay6hajhtsrzmh2qsng46rh5z72ckx7334")
		//addr: nch1f2h4shfaugqgmryg9wxjyu8ehhddc5yuh0t0fw
		if err != nil {
			t.Log(err)
			return
		}

		accAddr := sdk.AccAddress(pubkey.Address().Bytes())
		t.Log(fmt.Sprintf("addr = %s\n", accAddr.String()))

		pk, err := sdk.Bech32ifyAccPub(pubkey)
		t.Log(fmt.Sprintf("pubkey = %s\n", pk))

		t.Log(fmt.Sprintf("pubkey hex = %X", pubkey.Bytes()))

		type kk [33]byte
		assert.Equal(t, km.GetPrivKey().PubKey().VerifyBytes(msg, signature), true)
		v, _ := km.GetUCPubKey()
		t.Log(fmt.Sprintf("v = %x\n", v))

		addr, err := GetUCAddressBech32("04b2bf9a87dd7cf1ad998721ffef00713a4d5fb2bae0316eea04268ae877a0bcacd41b5b363911a30c0254ca12148d48e3cd4562e3e4b5d8cd3e6d2107a69754e6")
		t.Log(fmt.Sprintf("addr = %s\n", addr))
		t.Log(fmt.Sprintf("%s\n", km.GetAddr().String()))

	}
}
