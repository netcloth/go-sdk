package test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/netcloth/netcloth-chain/modules/auth"
	"github.com/netcloth/netcloth-chain/modules/cipal"

	"github.com/netcloth/go-sdk/client"
	"github.com/netcloth/go-sdk/client/lcd"
	"github.com/netcloth/go-sdk/util"
)

func Test_CIPALClaim(t *testing.T) {
	client, err := client.NewNCHClient("/Users/sky/go/src/github.com/netcloth/go-sdk/config/sdk.yaml")
	require.True(t, err == nil)

	expiration := time.Now().UTC().AddDate(0, 0, 1)
	adMsg := cipal.NewADParam(AccAddr, AccAddr, 6, expiration)
	sigBytes, err := client.TxClient.SignBytes(adMsg.GetSignBytes())
	if err != nil {
		panic(err)
	}

	stdSig := auth.StdSignature{
		PubKey:    client.TxClient.GetPrivKey().PubKey(),
		Signature: sigBytes,
	}

	req := cipal.NewIPALUserRequest(AccAddr, AccAddr, 6, expiration, stdSig)
	if res, err := client.CIPALClaim(req, "memo", false); err != nil {
		t.Fatal(err)
	} else {
		t.Log(util.ToJsonIgnoreErr(res))
	}
}

func Test_QueryCIPALChatServersAddrByUNCompressedPubKeys(t *testing.T) {
	client, err := client.NewNCHClient("/Users/sky/go/src/github.com/netcloth/go-sdk/config/sdk.yaml")
	require.True(t, err == nil)

	require.True(t, false) //TODO fix this unit test: for addrs below is bech32 addr but not uncompressedPubkey
	addrs := []string{0: "nch196mwu4e5l86t73rhw690xkfdagx6lkmkrxpsta", 1: "nch1f2h4shfaugqgmryg9wxjyu8ehhddc5yuh0t0fw"}
	if res, err := client.QueryCIPALChatServersAddrByUNCompressedPubKeys(addrs); err != nil {
		t.Fatal(err)
	} else {
		t.Log(util.ToJsonIgnoreErr(res))
	}
}

func Test_QueryCIPALsAddrByUNCompressedPubKeysByType(t *testing.T) {
	client, err := client.NewNCHClient("/Users/sky/go/src/github.com/netcloth/go-sdk/config/sdk.yaml")
	require.True(t, err == nil)

	require.True(t, false) //TODO fix this unit test: for addrs below is bech32 addr but not uncompressedPubkey
	addrs := []string{0: "nch196mwu4e5l86t73rhw690xkfdagx6lkmkrxpsta", 1: "nch1f2h4shfaugqgmryg9wxjyu8ehhddc5yuh0t0fw"}
	if res, err := client.QueryCIPALsAddrByUNCompressedPubKeysByType(addrs, lcd.EndpointTypeClientChat); err != nil {
		t.Fatal(err)
	} else {
		t.Log(util.ToJsonIgnoreErr(res))
	}
}
