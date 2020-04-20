package test

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/netcloth/go-sdk/client"
	"github.com/netcloth/go-sdk/util"
	"github.com/netcloth/netcloth-chain/hexutil"
	sdk "github.com/netcloth/netcloth-chain/types"
	"github.com/stretchr/testify/require"
)

const (
	// contract args
	functionName   = "revoke"
	fromBech32Addr = "nch13f5tmt88z5lkx8p45hv7a327nc0tpjzlwsq35e"
	toBech32Addr   = "nch1zypvh2q606ztw4elfgla0p6x4eruz3md6euv2t"
	pubKey         = 1025
	timestamp      = 1581065043
	r              = 2049
	s              = 4097

	// contract address
	contractBech32Addr = "nch1h50h6zyf2asumja34hh6gf8grmtdna08qzcmqr"
	contractAbi        = `[{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"from","type":"address"},{"indexed":false,"internalType":"address","name":"to","type":"address"},{"indexed":false,"internalType":"uint256","name":"timestamp","type":"uint256"},{"indexed":false,"internalType":"int64","name":"pk","type":"int64"}],"name":"Recall","type":"event"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"timestamp","type":"uint256"},{"internalType":"int64","name":"r","type":"int64"},{"internalType":"int64","name":"s","type":"int64"}],"name":"ecrecoverDecode","outputs":[{"internalType":"address","name":"addr","type":"address"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"address","name":"from","type":"address"}],"name":"queryParams","outputs":[{"internalType":"int64","name":"pubkey","type":"int64"},{"internalType":"uint256","name":"timestamp","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"timestamp","type":"uint256"},{"internalType":"int64","name":"pk","type":"int64"},{"internalType":"int64","name":"r","type":"int64"},{"internalType":"int64","name":"s","type":"int64"}],"name":"recall","outputs":[],"stateMutability":"nonpayable","type":"function"}]`
	/*
		合约代码参考：https://github.com/iavl/sol-demo/blob/master/demo2.sol
			function recall(address from, address to, uint timestamp, int64 pk, int64 r, int64 s) public {
			from 和to 是地址的二进制，如果是bech32地址需要先转为二进制的地址即[20]byte类型，再按照二进制的字符串形式打印成40个字符的字符串
			timestamp是时间戳，需要填充为32字节
			pubkey为公钥，需要填充为32字节
			r，s为签名
	*/

	// 地址20个字节，构造bytes32需要再填充12个字节的0
	addressPadZeros = "000000000000000000000000"

	// recall 函数参数
	payloadTemplate = "%s%s%064x%064x%064x%064x"
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
	//fmt.Println(fmt.Sprintf("%x", fromAddrBin.Bytes()))

	toAddrBin, err := sdk.AccAddressFromBech32(toBech32Addr)
	require.True(t, err == nil)
	toAddrStr := hexutil.Encode(toAddrBin.Bytes())

	// 构造合约的payload
	payloadStr := fmt.Sprintf(payloadTemplate, addressPadZeros+fromAddrStr, addressPadZeros+toAddrStr, timestamp, pubKey, r, s)
	//fmt.Println(fmt.Sprintf("payload:         %s ", payloadStr))
	argsBinary, err := hex.DecodeString(payloadStr)

	var payload []byte
	abiObj, _ := abi.JSON(strings.NewReader(contractAbi))
	m, _ := abiObj.Methods[functionName]

	readyArgs, err := m.Inputs.UnpackValues(argsBinary)
	require.NoError(t, err)

	payload, err = abiObj.Pack(functionName, readyArgs...)
	require.NoError(t, err)

	res, err := client.ContractCall(contractBech32Addr, payload, amount, true)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(util.ToJsonIgnoreErr(res))
	}
}

const txHash = "724B5E5141AD94E10C5945D55394A858963422CBC4D772E796E90E0D66943FE1"

func Test_ContractQuery(t *testing.T) {
	const (
		abiFilePath   = "/Users/sun/go/src/github.com/netcloth/go-sdk/config/contract.abi"
		eventFuncName = "Recall"
	)

	client, err := client.NewNCHClient(yaml_path)
	require.True(t, err == nil)

	txId, err := hexutil.Decode(txHash)
	r, err := client.QueryContractLog(txId)
	require.True(t, err == nil)

	//t.Log(r.Result)
	//if len(r.Result.Logs) == 0 {
	//	return
	//}

	t.Log(fmt.Sprintf("%s", r.Result.Logs[0].Data))
	s, _ := base64.StdEncoding.DecodeString(r.Result.Logs[0].Data)
	vs, err := util.UnpackValuesByABIFile(abiFilePath, eventFuncName, s)
	t.Log(err)

	for _, v := range vs {
		t.Log(fmt.Sprintf("%v\n", v))
	}
}

func Test_QueryContractEvents(t *testing.T) {
	// 遍历 [start, end] 之间的区块
	startBlockNum := int64(60)
	endBlockNum := int64(80)

	client, err := client.NewNCHClient(yaml_path)
	require.True(t, err == nil)

	// 查询合约相关的事件
	events, err := client.QueryContractEvents(contractBech32Addr, startBlockNum, endBlockNum)
	require.True(t, err == nil)

	const abiFilePath = "/Users/sun/nch/contract/r/r.abi"
	// 根据abi，解析出事件的data
	for _, e := range events {
		s, _ := hexutil.Decode(e)

		r, err := util.UnpackEventValuesByABIFile(abiFilePath, "Recall", s)
		if err != nil {
			t.Log(err)
		}

		require.Nil(t, err)

		t.Log(fmt.Sprintf("%v", r))
	}
}
