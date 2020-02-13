package tx

import (
	"fmt"

	"github.com/netcloth/go-sdk/client/types"
)

func (c *client) QueryContractEvents(contractBech32Addr string, startBlockNum int64, endBlockNum int64) (result []types.MsgDeleteResult, err error) {
	//var result types.MsgDeleteResult

	for i := startBlockNum; i < endBlockNum; i++ {
		fmt.Println("block ", i)
		blockResult, err := c.rpcClient.Block(i)
		if err != nil {
			continue
		}
		txs := blockResult.Block.Data.Txs
		fmt.Println(fmt.Sprintf("tx size: %d", len(txs)))
		for tx := range txs {
			fmt.Println(fmt.Sprintf("---- %v -----", tx))
		}
	}

	return result, nil
}
