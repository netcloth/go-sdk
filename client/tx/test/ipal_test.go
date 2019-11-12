package test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/netcloth/go-sdk/client"

	"github.com/netcloth/netcloth-chain/modules/auth"
	"github.com/netcloth/netcloth-chain/modules/cipal"
	"github.com/netcloth/netcloth-chain/modules/ipal"
	sdk "github.com/netcloth/netcloth-chain/types"

	"github.com/netcloth/go-sdk/util"
)

const (
	AccAddr = "nch1ugus2df3sydca3quula5yjqfntuq5aaxweezpt"
)

func Test_IPALClaim(t *testing.T) {
	client, err := client.NewNCHClient()
	require.True(t, err == nil)

	bond := sdk.Coin{
		Denom:  "unch",
		Amount: sdk.NewInt(1000000),
	}

	var eps ipal.Endpoints
	ep := ipal.NewEndpoint(1, "192.168.100.100:20000")
	eps = append(eps, ep)
	if res, err := client.IPALClaim("sky", "sky weibsite", "sky details", eps, bond, false); err != nil {
		t.Fatal(err)
	} else {
		t.Log(util.ToJsonIgnoreErr(res))
	}
}

func Test_CIPALClaim(t *testing.T) {
	client, err := client.NewNCHClient()
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
