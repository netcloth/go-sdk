package test

import (
	"testing"

	"github.com/netcloth/go-sdk/client"
	"github.com/netcloth/go-sdk/util"
	"github.com/stretchr/testify/require"
)

func Test_IPALListQuery(t *testing.T) {
	liteClient, err := client.NewNCHQueryClient()
	require.True(t, err == nil)

	if res, err := liteClient.QueryIPALList(); err != nil {
		t.Fatal(err)
	} else {
		t.Log(util.ToJsonIgnoreErr(res))
	}
}
