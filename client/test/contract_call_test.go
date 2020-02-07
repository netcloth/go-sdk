package test

import (
	"fmt"
	"testing"

	"github.com/netcloth/go-sdk/client"
	"github.com/netcloth/go-sdk/util"
	"github.com/netcloth/netcloth-chain/hexutil"
	sdk "github.com/netcloth/netcloth-chain/types"
	"github.com/stretchr/testify/require"
)

const (
	fromBech32Addr = "nch1yclzyuya2usepg80md5d6mrqypr8mhk3gl5a7x"
	toBech32Addr   = "nch1yclzyuya2usepg80md5d6mrqypr8mhk3gl5a7x"
	timestamp      = 1581065043
	r              = "1100000000000000000000000000000000000000000000000000000000000001"
	s              = "1100000000000000000000000000000000000000000000000000000000000001"
	v              = "1100000000000000000000000000000000000000000000000000000000000001"

	contractBech32Addr = "nch1yclzyuya2usepg80md5d6mrqypr8mhk3gl5a7x"
	payloadTemplate1   = "0xfc6a54a80000000000000000000000000dd023d5c543054c8612a2291b647c32d5714f510000000000000000000000000dd023d5c543054c8612a2291b647c32d5714f510000000000000000000000000000000000000000000000000000000000000001100000000000000000000000000000000000000000000000000000000000000120000000000000000000000000000000000000000000000000000000000000020300000000000000000000000000000000000000000000000000000000000000"

	/*
		第一个%s和第二个%s是地址的二进制，如果是bech32地址需要先转为二进制的地址即[20]byte类型，再按照二进制的字符串形式打印成40个字符的字符串
		%064x是时间戳
		最后3个%s是签名的r，s，v的二进制字符串
	*/
	payloadTemplate = "0xfc6a54a80000000000000000000000000%s0000000000000000000000000%s%064x%s%s%s"
)

var (
	amount = sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(0))
)

func Test_ContractCall(t *testing.T) {
	client, err := client.NewNCHClient("/Users/sky/go/src/github.com/netcloth/go-sdk/config/sdk.yaml")
	require.True(t, err == nil)

	fromAddrBin, err := sdk.AccAddressFromBech32(fromBech32Addr)
	require.True(t, err == nil)
	fromAddrStr := hexutil.Encode(fromAddrBin.Bytes())

	toAddrBin, err := sdk.AccAddressFromBech32(toBech32Addr)
	require.True(t, err == nil)
	toAddrStr := hexutil.Encode(toAddrBin)

	payloadStr := fmt.Sprintf(payloadTemplate, fromAddrStr, toAddrStr, timestamp, r, s, v)
	t.Log(payloadStr)
	payload, err := hexutil.Decode(payloadStr)
	require.True(t, err == nil)

	res, err := client.ContractCall(contractBech32Addr, payload, amount, true)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(util.ToJsonIgnoreErr(res))
	}

	//txId := res.CommitResult.Hash
}

func Test_ContractQuery(t *testing.T) {
	client, err := client.NewNCHClient("/Users/sky/go/src/github.com/netcloth/go-sdk/config/sdk.yaml")
	require.True(t, err == nil)

	txId, err := hexutil.Decode("d6063cf0ac432a27645f9defb75747fd3cd7e6157f5c4e23b8de0898c2571f50")
	r, err := client.QueryContractLog(txId)
	require.True(t, err == nil)

	t.Log(r.Result.Logs[0].Data)
}
