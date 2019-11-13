package test

import (
	"testing"

	"github.com/netcloth/go-sdk/client"

	"github.com/netcloth/go-sdk/util"
	"github.com/stretchr/testify/require"
)

func Test_IPALQuery(t *testing.T) {
	liteClient, err := client.NewNCHQueryClient("/Users/sky/go/src/github.com/netcloth/go-sdk/config/sdk.yaml")
	require.True(t, err == nil)

	if res, err := liteClient.QueryIPALByAddress("nch1f2h4shfaugqgmryg9wxjyu8ehhddc5yuh0t0fw"); err != nil {
		t.Fatal(err)
	} else {
		t.Log(util.ToJsonIgnoreErr(res))
	}
}
