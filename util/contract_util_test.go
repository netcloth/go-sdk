package util

import (
	"fmt"
	"github.com/netcloth/netcloth-chain/hexutil"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	abiFile  = "/Users/sun/Desktop/abi"
	funcName = "ipals"
	res      = "00000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000f42400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000116e6574636c6f74682d6f6666696369616c00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003277b226f70657261746f725f61646472657373223a226e636831306a7a70743332677772616476396d636e72366675756a30746e783772713070736d6d746a75222c226d6f6e696b6572223a226e6574636c6f74682d6f6666696369616c222c2277656273697465223a226e6574636c6f74682e6f7267222c2264657461696c73223a226e6574636c6f74682d6f6666696369616c222c22656e64706f696e7473223a5b7b2274797065223a2231222c22656e64706f696e74223a22687474703a2f2f34372e3130342e3138392e35227d2c7b2274797065223a2233222c22656e64706f696e74223a22687474703a2f2f34372e39302e352e313338227d2c7b2274797065223a2234222c22656e64706f696e74223a227b5c226d696e69417070446f6d61696e735c223a5b7b5c226d6f6e696b65725c223a5c224e6574436c6f746820426c6f675c222c5c22646f6d61696e5c223a5c2268747470733a2f2f626c6f672e6e6574636c6f74682e6f72675c227d2c7b5c226d6f6e696b65725c223a5c22e993bee997bbe7a4be5c222c5c22646f6d61696e5c223a5c2268747470733a2f2f7777772e636861696e6e6577732e636f6d2f5c227d2c7b5c226d6f6e696b65725c223a5c22e99d9ee5b08fe58fb75c222c5c22646f6d61696e5c223a5c2268747470733a2f2f6665697869616f68616f2e636f6d5c227d2c7b5c226d6f6e696b65725c223a5c22e98791e8b4a2e5bfabe8aeaf5c222c5c22646f6d61696e5c223a5c2268747470733a2f2f6d2e6a696e73652e636f6d2f6c697665735c227d2c7b5c226d6f6e696b65725c223a5c224e6574436c6f746820426c6f675c222c5c22646f6d61696e5c223a5c2268747470733a2f2f6d656469756d2e636f6d2f404e6574436c6f74682f5c227d2c7b5c226d6f6e696b65725c223a5c22436f696e6465736b5c222c5c22646f6d61696e5c223a5c2268747470733a2f2f7777772e636f696e6465736b2e636f6d5c227d2c7b5c226d6f6e696b65725c223a5c22436f696e6d61726b65746361705c222c5c22646f6d61696e5c223a5c2268747470733a2f2f7777772e636f696e6d61726b65746361702e636f6d205c227d5d7d227d5d7d00000000000000000000000000000000000000000000000000"
)

func Test_UnpackValuesByABIFile(t *testing.T) {
	d, err := hexutil.Decode(res)
	require.True(t, err == nil)

	values, err := UnpackValuesByABIFile(abiFile, funcName, d)
	require.True(t, err == nil)

	for _, v := range values {
		t.Log(fmt.Sprintf("%v", v))
	}
}
