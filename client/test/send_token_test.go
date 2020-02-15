package test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/netcloth/go-sdk/client"
	"github.com/netcloth/go-sdk/client/types"
	"github.com/netcloth/go-sdk/util"
)

func TestClient_SendToken(t *testing.T) {
	c, err := client.NewNCHTXClient(yaml_path)
	require.True(t, err == nil)

	coins := []types.Coin{
		{
			Denom:  "pnch",
			Amount: "100",
		},
	}
	if res, err := c.SendToken("nch1ugus2df3sydca3quula5yjqfntuq5aaxweezpt", coins, "", false); err != nil {
		t.Fatal(err)
	} else {
		t.Log(util.ToJsonIgnoreErr(res))
	}
}
