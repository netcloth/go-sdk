package tx

import (
	"fmt"
	"strconv"

	"github.com/netcloth/go-sdk/constants"

	"github.com/netcloth/go-sdk/config"

	"github.com/netcloth/netcloth-chain/modules/auth"
	"github.com/netcloth/netcloth-chain/modules/bank"
	sdk "github.com/netcloth/netcloth-chain/types"

	"github.com/netcloth/go-sdk/client/types"
	"github.com/netcloth/go-sdk/types/tx"
)

func (c *client) SendToken(receiver string, coins []types.Coin, memo string, commit bool) (types.BroadcastTxResult, error) {
	var (
		result types.BroadcastTxResult
	)
	from := c.KeyManager.GetAddr()

	to, err := types.AccAddrFromBech32(receiver)
	if err != nil {
		return result, err
	}

	sdkCoins, err := buildCoins(coins)
	if err != nil {
		return result, err
	}
	msg := buildBankSendMsg(from, to, sdkCoins)

	accountBody, err := c.liteClient.QueryAccount(from.String())
	if err != nil {
		return result, err
	}

	//  check balance is enough
	amount := getCoin(accountBody.Result.Value.Coins, config.TxDefaultDenom)

	totalfee := sdk.NewInt(config.TxDefaultFeeAmount)
	for _, val := range sdkCoins {
		if val.Denom == config.TxDefaultDenom {
			totalfee = totalfee.Add(val.Amount)
		}
	}

	if amount.Amount.LT(totalfee) {
		return result, fmt.Errorf("account balance is not enough")
	}

	fee := sdk.Coins{
		{
			Denom:  config.TxDefaultDenom,
			Amount: sdk.NewInt(config.TxDefaultFeeAmount),
		},
	}
	an, err := strconv.Atoi(accountBody.Result.Value.AccountNumber)
	s, err := strconv.Atoi(accountBody.Result.Value.Sequence)
	stdSignMsg := tx.StdSignMsg{
		ChainID:       c.chainId,
		AccountNumber: uint64(an),
		Sequence:      uint64(s),
		Fee:           auth.NewStdFee(config.TxDefaultGas, fee),
		Msgs:          []sdk.Msg{msg},
		Memo:          memo,
	}

	for _, m := range stdSignMsg.Msgs {
		if err := m.ValidateBasic(); err != nil {
			return result, err
		}
	}

	txBytes, err := c.KeyManager.Sign(stdSignMsg)
	if err != nil {
		return result, err
	}

	var txBroadcastType string
	if commit {
		txBroadcastType = constants.TxBroadcastTypeCommit
	} else {
		txBroadcastType = constants.TxBroadcastTypeSync
	}

	return c.rpcClient.BroadcastTx(txBroadcastType, txBytes)
}

func buildCoins(icoins []types.Coin) (sdk.Coins, error) {
	var (
		coins []sdk.Coin
	)
	if len(icoins) == 0 {
		return coins, nil
	}
	for _, v := range icoins {
		amount, ok := sdk.NewIntFromString(v.Amount)
		if ok {
			coins = append(coins, sdk.Coin{
				Denom:  v.Denom,
				Amount: amount,
			})
		} else {
			return coins, fmt.Errorf("can't parse str to Int, coin is %+v", icoins)
		}
	}

	return coins, nil
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
func buildBankSendMsg(from sdk.AccAddress, to sdk.AccAddress, coins sdk.Coins) bank.MsgSend {
	msg := bank.MsgSend{
		FromAddress: from,
		ToAddress:   to,
		Amount:      coins,
	}
	return msg
}
