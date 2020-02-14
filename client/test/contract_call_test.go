package test

import (
	"encoding/base64"
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
	// contract args
	fromBech32Addr = "nch13f5tmt88z5lkx8p45hv7a327nc0tpjzlwsq35e"
	toBech32Addr   = "nch1zypvh2q606ztw4elfgla0p6x4eruz3md6euv2t"
	pubKey         = "1100000000000000000000000000000000000000000000000000000000000001"
	timestamp      = 1581065043
	r              = "1100000000000000000000000000000000000000000000000000000000000001"
	s              = "1100000000000000000000000000000000000000000000000000000000000001"

	// contract address
	contractBech32Addr = "nch1vdm60zm5jr5yn3aj0dkfkghf23x4t2yw9y2j4f"

	/*
		    function recall(address from, address to, uint timestamp, bytes32 pubkey, byte t, bytes32 r, bytes32 s, byte v) public {
			from 和to 是地址的二进制，如果是bech32地址需要先转为二进制的地址即[20]byte类型，再按照二进制的字符串形式打印成40个字符的字符串
			timestamp是时间戳，需要填充为32字节
			pubkey为公钥，需要填充为32字节
			pubKeyType为公钥类型，1表示个人，2表示群,需要填充为32字节
			最后3个%s是签名的r，s的二进制字符串
	*/

	addressPadZeros = "000000000000000000000000"

	// fc6a54a8 is recall function signature
	payloadTemplate = "fc6a54a8%s%s%064x%s%s%s"
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
	//fmt.Println(fromAddrStr)

	toAddrBin, err := sdk.AccAddressFromBech32(toBech32Addr)
	require.True(t, err == nil)
	toAddrStr := hexutil.Encode(toAddrBin)
	//fmt.Println(fmt.Sprintf("%x", toAddrBin.Bytes()))
	//fmt.Println(toAddrStr)

	payloadStr := fmt.Sprintf(payloadTemplate, addressPadZeros+fromAddrStr, addressPadZeros+toAddrStr, timestamp, pubKey, r, s)

	t.Log(payloadStr)
	payload, err := hexutil.Decode(payloadStr)
	require.True(t, err == nil)

	res, err := client.ContractCall(contractBech32Addr, payload, amount, true)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(util.ToJsonIgnoreErr(res))
	}
}

const txHash = "EA5D94E02EF77AF61D3CEA6056348D1E400C0728E059EB57F3D8A3E896390427"

func Test_ContractQuery(t *testing.T) {
	client, err := client.NewNCHClient(yaml_path)
	require.True(t, err == nil)

	txId, err := hexutil.Decode(txHash)
	r, err := client.QueryContractLog(txId)
	require.True(t, err == nil)

	t.Log(r.Result.Logs[0].Data)

	/*
		item := r.Result.Logs[0].Data
		s, _ := base64.StdEncoding.DecodeString(item)
		fmt.Println(fmt.Sprintf("%d, %x", len(s), s))

		// 第一个byte32为from地址
		a := fmt.Sprintf("%x", s[12:32])
		// 第二个byte32为to地址
		b := fmt.Sprintf("%x", s[44:64])
		// 为int64类型的timestame
		c := fmt.Sprintf("%x", s[64:96])

		d := fmt.Sprintf("%x", s[96:128])

		//e ;= fmt.Sprintf("%x", s[128:])

		// 输出
		accA, _ := sdk.AccAddressFromHex(a)
		fmt.Println(fmt.Sprintf("%s --> %s", a, accA.String()))

		accB, _ := sdk.AccAddressFromHex(b)
		fmt.Println(fmt.Sprintf("%s --> %s", b, accB.String()))

		timestamp, _ := strconv.ParseUint(c, 16, 64)
		fmt.Println(fmt.Sprintf("%s --> %d", c, timestamp))

		fmt.Println(fmt.Sprintf("%s --> %s", d, d))
	*/
}

func Test_QueryContractEvents(t *testing.T) {
	// 遍历 [start, end] 之间的区块
	startBlockNum := int64(1287)
	endBlockNum := int64(1289)

	client, err := client.NewNCHClient(yaml_path)
	require.True(t, err == nil)

	// 查询合约相关的事件
	res, err := client.QueryContractEvents(contractBech32Addr, startBlockNum, endBlockNum)
	require.True(t, err == nil)

	// 根据abi，解析出事件的data
	fmt.Println("result:")
	for _, item := range res {
		s, _ := base64.StdEncoding.DecodeString(item)
		fmt.Println(fmt.Sprintf("%d, %x", len(s), s))

		// 第一个byte32为from地址
		a := fmt.Sprintf("%x", s[12:32])
		// 第二个byte32为to地址
		b := fmt.Sprintf("%x", s[44:64])
		// 为int64类型的timestame
		c := fmt.Sprintf("%x", s[64:96])
		// pubkey
		d := fmt.Sprintf("%x", s[96:128])

		// 输出
		accA, _ := sdk.AccAddressFromHex(a)
		fmt.Println(fmt.Sprintf("%s --> %s", a, accA.String()))

		accB, _ := sdk.AccAddressFromHex(b)
		fmt.Println(fmt.Sprintf("%s --> %s", b, accB.String()))

		timestamp, _ := strconv.ParseUint(c, 16, 64)
		fmt.Println(fmt.Sprintf("%s --> %d", c, timestamp))

		fmt.Println(fmt.Sprintf("%s --> %s", d, d))

		t.Log(item)
	}
}
