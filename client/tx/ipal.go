package tx

import (
	"fmt"
	"os"
	"strconv"

	"github.com/netcloth/go-sdk/constants"

	"github.com/netcloth/go-sdk/config"

	"github.com/netcloth/netcloth-chain/modules/ipal"

	"github.com/netcloth/netcloth-chain/modules/auth"
	"github.com/netcloth/netcloth-chain/modules/cipal"
	sdk "github.com/netcloth/netcloth-chain/types"

	"github.com/netcloth/go-sdk/client/types"
	"github.com/netcloth/go-sdk/types/tx"
)

func (c *client) CIPALClaim(req cipal.IPALUserRequest, memo string, commit bool) (types.BroadcastTxResult, error) {
	var result types.BroadcastTxResult
	from := c.KeyManager.GetAddr()

	msg := buildBankIPALClaimMsg(from, req)

	accountBody, err := c.liteClient.QueryAccount(from.String())
	if err != nil {
		return result, err
	}

	amount := getCoin(accountBody.Result.Value.Coins, config.TxDefaultDenom)

	totalfee := sdk.NewInt(config.TxDefaultFeeAmount)
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

	fmt.Fprintf(os.Stderr, "stdSignMsg = %v\n", stdSignMsg)
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

func buildBankIPALClaimMsg(from sdk.AccAddress, userRequest cipal.IPALUserRequest) cipal.MsgIPALClaim {
	msg := cipal.MsgIPALClaim{
		From:        from,
		UserRequest: userRequest,
	}
	return msg
}

func (c *client) IPALClaim(moniker, website, details string, endpoints ipal.Endpoints, bond sdk.Coin, commit bool) (r types.BroadcastTxResult, err error) {
	var result types.BroadcastTxResult

	from := c.KeyManager.GetAddr() //from is operator_address

	accountBody, err := c.liteClient.QueryAccount(from.String())
	if err != nil {
		return result, err
	}

	if bond.Denom != config.TxDefaultDenom {
		return result, err
	}

	amount := getCoin(accountBody.Result.Value.Coins, config.TxDefaultDenom)

	totalfee := sdk.NewInt(config.TxDefaultFeeAmount)
	if amount.Amount.LT(totalfee.Add(bond.Amount)) {
		return result, fmt.Errorf("account balance is not enough")
	}

	fee := sdk.Coins{
		{
			Denom:  config.TxDefaultDenom,
			Amount: sdk.NewInt(config.TxDefaultFeeAmount),
		},
	}

	msg := ipal.NewMsgServiceNodeClaim(from, moniker, website, details, endpoints, bond)

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
