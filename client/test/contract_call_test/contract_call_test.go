package contract_call_test

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strconv"
	"testing"

	"github.com/netcloth/go-sdk/client"
	"github.com/netcloth/go-sdk/util"
	"github.com/netcloth/netcloth-chain/hexutil"
	sdk "github.com/netcloth/netcloth-chain/types"
	"github.com/stretchr/testify/require"
)

const (
	yamlPath           = "/Users/sky/go/src/github.com/netcloth/go-sdk/config/sdk.yaml"
	contractBech32Addr = "nch13kmemljzcnm6jyku8xnejmc62vdunfpps7jjj9"
)

var (
	amount = sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(0))
)

func Test_ContractCall(t *testing.T) {
	const (
		functionSig         = "4ee604a3" // the first 4 bytes of sig of function: revoke
		payloadTemplate     = "%s%s%s%064x%064x%s%s%064x"
		fromPubkeyHexString = "8c36b163c26f492abc874648b7258450394fe78133bcc4d920895d0ce8c3ac4e"
		toPubKeyHexString   = "8c36b163c26f492abc874648b7258450394fe78133bcc4d920895d0ce8c3ac4e"
		revokeType          = 0
		timestamp           = 1581065043
		rHexString          = "1f9b85de5bfdea9b1310e6712b7585d19202ef1190eed0d5ecc6449108893fde"
		sHexString          = "7d36d77636032619603036c0deb4c6d985ebc44496f4cfc81a0acc4ed63c21b3"
		v                   = 0
	)

	client, err := client.NewNCHClient(yamlPath)
	t.Log(err)
	require.True(t, err == nil)

	// 构造合约的payload
	payloadStr := fmt.Sprintf(payloadTemplate, functionSig, fromPubkeyHexString, toPubKeyHexString, revokeType, timestamp, rHexString, sHexString, v)
	fmt.Println(fmt.Sprintf("payload: %s", payloadStr))
	payload, err := hex.DecodeString(payloadStr)
	require.NoError(t, err)

	res, err := client.ContractCall(contractBech32Addr, payload, amount, true)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(util.ToJsonIgnoreErr(res))
	}
}

const txHash = "A1AA0EFFB1FD1C1E93830B71789FB4A67FB2E076307BCEA635522143BB12A7D7"

func Test_ContractQuery(t *testing.T) {
	client, err := client.NewNCHClient(yamlPath)
	require.True(t, err == nil)

	txId, err := hexutil.Decode(txHash)
	r, err := client.QueryContractLog(txId)
	require.True(t, err == nil)

	t.Log(r.Result.Logs[0].Data)

	item := r.Result.Logs[0].Data

	fromPubkeyStr := item[:64]
	toPubkeyStr := item[64:128]
	revokeTypeStr := item[128:192]
	timestampStr := item[192:]

	t.Log(fromPubkeyStr)
	t.Log(toPubkeyStr)
	t.Log(revokeTypeStr)
	t.Log(timestampStr)
}

type MsgDeleteResult struct {
	from      string `json:"from" yaml:"from"`
	to        string `json:"to" yaml:"to"`
	pubkey    uint64 `json:"pubkey" yaml:"pubkey"`
	timestamp uint64 `json:"timestamp" yaml:"timestamp"`
}

func (res MsgDeleteResult) String() string {
	return fmt.Sprintf(
		`
from: %s
to: %s
pubkey: %d
timestamp: %d`, res.from, res.to, res.pubkey, res.timestamp)
}
func Test_QueryContractEvents(t *testing.T) {
	// 遍历 [start, end] 之间的区块
	startBlockNum := int64(6280)
	endBlockNum := int64(6470)

	client, err := client.NewNCHClient(yamlPath)
	require.True(t, err == nil)

	// 查询合约相关的事件
	res, err := client.QueryContractEvents(contractBech32Addr, startBlockNum, endBlockNum)
	require.True(t, err == nil)

	// 根据abi，解析出事件的data
	var results []MsgDeleteResult
	for _, item := range res {
		var result MsgDeleteResult

		s, _ := base64.StdEncoding.DecodeString(item)

		// 第一个byte32为from地址
		a := fmt.Sprintf("%x", s[12:32])
		// 第二个byte32为to地址
		b := fmt.Sprintf("%x", s[44:64])
		// 为int64类型的timestame
		c := fmt.Sprintf("%x", s[64:96])
		// pubkey
		d := fmt.Sprintf("%x", s[96:128])

		// address - from
		accA, _ := sdk.AccAddressFromHex(a)
		// address - to
		accB, _ := sdk.AccAddressFromHex(b)
		// uint - timestamp
		timestamp, _ := strconv.ParseUint(c, 16, 64)
		// int64 - public key
		pk, _ := strconv.ParseUint(d, 16, 64)

		result.from = accA.String()
		result.to = accB.String()
		result.pubkey = pk
		result.timestamp = timestamp
		results = append(results, result)

		t.Log(item)
	}

	fmt.Println(results)
}
