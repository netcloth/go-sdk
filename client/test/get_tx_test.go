package test

import (
	"testing"

	"github.com/netcloth/go-sdk/types"
	"github.com/netcloth/go-sdk/types/tx"
	"github.com/netcloth/go-sdk/util"
)

func TestClient_GetTx(t *testing.T) {
	hash := "D38D2EE53C254BD5723F0004FD0D206C9B53BE512E322D0332141088398DDC51"
	if res, err := c.GetTx(hash); err != nil {
		t.Fatal(err)
	} else {
		t.Log(util.ToJsonIgnoreErr(res))
	}
}

func TestClient_GetTxMsgs(t *testing.T) {
	hash := "D38D2EE53C254BD5723F0004FD0D206C9B53BE512E322D0332141088398DDC51"
	if res, err := c.GetTx(hash); err != nil {
		t.Fatal(err)
	} else {
		t.Log(util.ToJsonIgnoreErr(res))
		var msgSend types.MsgSend
		for _, msg := range res.Tx.Msgs {
			err := tx.Cdc.UnmarshalJSON(msg.GetSignBytes(), &msgSend)
			if err != nil {
				continue
			}

			t.Log(msgSend.FromAddress.String())
			t.Log(msgSend.ToAddress.String())
			t.Log(msgSend.Amount.String())
		}
	}
}
