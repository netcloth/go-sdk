package test

import (
	"testing"

	"github.com/netcloth/go-sdk/client/basic"
	"github.com/netcloth/go-sdk/client/lcd"
	"github.com/netcloth/go-sdk/util"
)

func Test_IPALQuery(t *testing.T) {
	basicClient := basic.NewClient("http://127.0.0.1:1317")
	lite := lcd.NewClient(basicClient)

	if res, err := lite.QueryIPALByAddress("nch1f2h4shfaugqgmryg9wxjyu8ehhddc5yuh0t0fw"); err != nil {
		t.Fatal(err)
	} else {
		t.Log(util.ToJsonIgnoreErr(res))
	}
}
