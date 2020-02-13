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
	fromBech32Addr = "nch13f5tmt88z5lkx8p45hv7a327nc0tpjzlwsq35e"
	toBech32Addr   = "nch1zypvh2q606ztw4elfgla0p6x4eruz3md6euv2t"
	timestamp      = 1581065043
	r              = "1100000000000000000000000000000000000000000000000000000000000001"
	s              = "1100000000000000000000000000000000000000000000000000000000000001"
	v              = "1100000000000000000000000000000000000000000000000000000000000001"

	contractBech32Addr = "nch17awtgfpq30xgzs5eld93pev5u588yvmqpruv3s"
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
	client, err := client.NewNCHClient(yaml_path)
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

const txHash = "89A5480747829A680437B663C7681DE1C2E7869D3031B37D3136CF4330201B2E"

func Test_ContractQuery(t *testing.T) {
	client, err := client.NewNCHClient(yaml_path)
	require.True(t, err == nil)

	txId, err := hexutil.Decode(txHash)
	r, err := client.QueryContractLog(txId)
	require.True(t, err == nil)

	t.Log(r.Result.Logs[0].Data)
}

func Test_QueryContractEvents(t *testing.T) {
	startBlockNum := int64(10400)
	endBlockNum := int64(10509)

	client, err := client.NewNCHClient(yaml_path)
	require.True(t, err == nil)

	res, err := client.QueryContractEvents(contractBech32Addr, startBlockNum, endBlockNum)
	require.True(t, err == nil)

	fmt.Println("result:")
	for _, item := range res {
		t.Log(item)
	}
}
