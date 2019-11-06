package tx

import (
	"fmt"
	"strconv"

	"github.com/NetCloth/netcloth-chain/modules/auth"
	"github.com/NetCloth/netcloth-chain/modules/bank"
	"github.com/NetCloth/netcloth-chain/modules/ipal"

	sdk "github.com/NetCloth/netcloth-chain/types"

	"github.com/NetCloth/go-sdk/client/types"
	"github.com/NetCloth/go-sdk/types/tx"
	"github.com/NetCloth/go-sdk/util/constant"
)

func (c *client) IPALClaim(userRequest ipal.IPALUserRequest, memo string, commit bool) (types.BroadcastTxResult, error) {
	var (
		result types.BroadcastTxResult
	)
	from := c.keyManager.GetAddr()

	// check userRequest

	msg := buildBankIPALClaimMsg(from, userRequest)

	accountBody, err := c.liteClient.QueryAccount(from.String())
	if err != nil {
		return result, err
	}

	//  check balance is enough
	amount := getCoin(accountBody.Result.Value.Coins, constant.TxDefaultFeeDenom)

	totalfee := sdk.NewInt(constant.TxDefaultFeeAmount)
	if amount.Amount.LT(totalfee) {
		return result, fmt.Errorf("account balance is not enough")
	}

	fee := sdk.Coins{
		{
			Denom:  constant.TxDefaultFeeDenom,
			Amount: sdk.NewInt(constant.TxDefaultFeeAmount),
		},
	}
	an, err := strconv.Atoi(accountBody.Result.Value.AccountNumber)
	s, err := strconv.Atoi(accountBody.Result.Value.Sequence)
	stdSignMsg := tx.StdSignMsg{
		ChainID:       c.chainId,
		AccountNumber: uint64(an),
		Sequence:      uint64(s),
		Fee:           auth.NewStdFee(constant.TxDefaultGas, fee),
		Msgs:          []sdk.Msg{msg},
		Memo:          memo,
	}

	for _, m := range stdSignMsg.Msgs {
		if err := m.ValidateBasic(); err != nil {
			return result, err
		}
	}

	txBytes, err := c.keyManager.Sign(stdSignMsg)
	if err != nil {
		return result, err
	}

	var txBroadcastType string
	if commit {
		txBroadcastType = constant.TxBroadcastTypeCommit
	} else {
		txBroadcastType = constant.TxBroadcastTypeSync
	}

	return c.rpcClient.BroadcastTx(txBroadcastType, txBytes)
}

func getCoin(icoins []types.Coin, denom string) sdk.Coin {
	for _, vcoin := range icoins {
		if vcoin.Denom == denom {
			amount, ok := sdk.NewIntFromString(vcoin.Amount)
			if ok {
				return sdk.Coin{
					Denom:  vcoin.Denom,
					Amount: amount,
				}
			}

		}
	}
	return sdk.Coin{}
}

// buildBankSendMsg builds the sending coins msg
func buildBankIPALClaimMsg(from sdk.AccAddress, userRequest ipal.IPALUserRequest) bank.MsgSend {
	msg := ipal.MsgIPALClaim{
		From:        from,
		UserRequest: userRequest,
	}
	return msg
}
