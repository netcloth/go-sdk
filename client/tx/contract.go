package tx

import (
	"fmt"
	"strconv"

	"github.com/netcloth/go-sdk/client/types"
	"github.com/netcloth/go-sdk/config"
	"github.com/netcloth/go-sdk/constants"
	"github.com/netcloth/go-sdk/types/tx"
	"github.com/netcloth/netcloth-chain/modules/auth"
	vmtypes "github.com/netcloth/netcloth-chain/modules/vm/types"
	sdk "github.com/netcloth/netcloth-chain/types"
)

func (c *client) ContractCall(contractBech32Addr string, payload []byte, amount sdk.Coin, commit bool) (r types.BroadcastTxResult, err error) {
	var result types.BroadcastTxResult

	from := c.KeyManager.GetAddr()

	var contractAddr sdk.AccAddress
	if contractBech32Addr != "" {
		contractAddr, err = sdk.AccAddressFromBech32(contractBech32Addr)
		if err != nil {
			return result, err
		}
	}

	accountBody, err := c.liteClient.QueryAccount(from.String())
	if err != nil {
		return result, err
	}

	if amount.Denom != constants.TxDefaultDenom {
		return result, err
	}

	accountAmount := getCoin(accountBody.Result.Value.Coins, constants.TxDefaultDenom)

	totalfee := sdk.NewInt(config.TxDefaultFeeAmount)
	if accountAmount.Amount.LT(totalfee.Add(amount.Amount)) {
		return result, fmt.Errorf("account balance is not enough")
	}

	fee := sdk.Coins{
		{
			Denom:  constants.TxDefaultDenom,
			Amount: sdk.NewInt(config.TxDefaultFeeAmount),
		},
	}

	msg := vmtypes.NewMsgContract(from, contractAddr, payload, amount)

	an, err := strconv.Atoi(accountBody.Result.Value.AccountNumber)
	if err != nil {
		return result, err
	}

	s, err := strconv.Atoi(accountBody.Result.Value.Sequence)
	if err != nil {
		return result, err
	}

	stdSignMsg := tx.StdSignMsg{
		ChainID:       c.chainId,
		AccountNumber: uint64(an),
		Sequence:      uint64(s),
		Fee:           auth.NewStdFee(config.TxDefaultGas, fee),
		Msgs:          []sdk.Msg{msg},
		Memo:          "",
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
