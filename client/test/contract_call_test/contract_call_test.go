package contract_call_test

import (
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
	contractBech32Addr = "nch1kl6qunfqus9xlt4zpt89q2r3ty2y2l8f4348wt"
)

var (
	amount = sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(0))
)

func Test_ContractCall(t *testing.T) {
	const (
		functionSig     = "81a8a747" // the first 4 bytes of sig of function: recall
		payloadTemplate = "%s%s%s000000000000000000000000%s%064x%064x%s%s%064x"

		fromPubkeyHexString = "8c36b163c26f492abc874648b7258450394fe78133bcc4d920895d0ce8c3ac4e"
		toPubKeyHexString   = "8c36b163c26f492abc874648b7258450394fe78133bcc4d920895d0ce8c3ac4e"
		fromAddr            = "692a70d2e424a56d2c6c27aa97d1a86395877b3a"
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
	payloadStr := fmt.Sprintf(payloadTemplate, functionSig, fromPubkeyHexString, toPubKeyHexString, fromAddr, revokeType, timestamp, rHexString, sHexString, v)
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

const txHash = "4ACAD763B7CFF863F832FD3DB15A67C9C7B960E5698151CD5A0C51FD7DEF2CD5"

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

func Test_QueryContractEvents(t *testing.T) {
	const (
		startBlockNum = 4400
		endBlockNum   = 4446
	)

	client, err := client.NewNCHClient(yamlPath)
	require.True(t, err == nil)

	res, err := client.QueryContractEvents(contractBech32Addr, startBlockNum, endBlockNum)
	require.True(t, err == nil)
	t.Log(res)

	for _, item := range res {
		t.Log(item)

		fromPubkeyStr := item[:64]
		toPubkeyStr := item[64:128]
		revokeTypeStr := item[128:192]
		timestampStr := item[192:]

		revokeType, _ := strconv.ParseUint(revokeTypeStr, 16, 64)
		timestamp, _ := strconv.ParseUint(timestampStr, 16, 64)

		t.Log(fromPubkeyStr)
		t.Log(toPubkeyStr)
		t.Log(revokeTypeStr)
		t.Log(timestampStr)

		t.Log(revokeType)
		t.Log(timestamp)
	}
}
