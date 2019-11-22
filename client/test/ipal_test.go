package test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/netcloth/netcloth-chain/modules/ipal"
	sdk "github.com/netcloth/netcloth-chain/types"

	"github.com/netcloth/go-sdk/client"
	"github.com/netcloth/go-sdk/util"
)

const (
	AccAddr = "nch1ugus2df3sydca3quula5yjqfntuq5aaxweezpt"
)

func Test_IPALClaim(t *testing.T) {
	client, err := client.NewNCHClient("/Users/sky/go/src/github.com/netcloth/go-sdk/config/sdk.yaml")
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

func Test_IPALQuery(t *testing.T) {
	liteClient, err := client.NewNCHQueryClient("/Users/sky/go/src/github.com/netcloth/go-sdk/config/sdk.yaml")
	require.True(t, err == nil)

	if res, err := liteClient.QueryIPALByAddress("nch1f2h4shfaugqgmryg9wxjyu8ehhddc5yuh0t0fw"); err != nil {
		t.Fatal(err)
	} else {
		t.Log(util.ToJsonIgnoreErr(res))
	}
}

func Test_QueryIPALChatServersEndpointByAddresses(t *testing.T) {
	client, err := client.NewNCHClient("/Users/sky/go/src/github.com/netcloth/go-sdk/config/sdk.yaml")
	require.True(t, err == nil)

	addrs := []string{0: "nch196mwu4e5l86t73rhw690xkfdagx6lkmkrxpsta", 1: "nch1f2h4shfaugqgmryg9wxjyu8ehhddc5yuh0t0fw"}
	if res, err := client.QueryIPALChatServersEndpointByAddresses(addrs); err != nil {
		t.Fatal(err)
	} else {
		t.Log(util.ToJsonIgnoreErr(res))
	}
}

func Test_IPALListQuery(t *testing.T) {
	liteClient, err := client.NewNCHQueryClient("/Users/sky/go/src/github.com/netcloth/go-sdk/config/sdk.yaml")
	require.True(t, err == nil)

	if res, err := liteClient.QueryIPALList(); err != nil {
		t.Fatal(err)
	} else {
		t.Log(util.ToJsonIgnoreErr(res))
	}
}
