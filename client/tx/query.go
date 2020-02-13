package tx

import (
	"encoding/hex"
	"fmt"

	"github.com/tendermint/tendermint/crypto"

	vmtypes "github.com/netcloth/netcloth-chain/modules/vm/types"
)

func (c *client) QueryContractEvents(contractBech32Addr string, startBlockNum int64, endBlockNum int64) (result []string, err error) {
	// iterator blocks
	for i := startBlockNum; i < endBlockNum; i++ {
		//fmt.Println("block ", i)
		blockResult, err := c.rpcClient.Block(i)
		if err != nil {
			continue
		}
		// iterator txs in block
		txs := blockResult.Block.Data.Txs
		for _, tx := range txs {
			txhash := crypto.Sha256(tx)
			fmt.Println(fmt.Sprintf("block %d, txhash: %x", i, txhash))

			if res, err := c.rpcClient.GetTx(hex.EncodeToString(txhash)); err == nil {
				msg := res.Tx.Msgs[0]
				if msg.Type() == vmtypes.TypeMsgContract {
					msgContract, _ := msg.(vmtypes.MsgContract)
					targetContractAddr := msgContract.To.String()
					if contractBech32Addr != targetContractAddr {
						continue
					}

					// query events by txhash
					eventLog, _ := c.liteClient.QueryContractLog(txhash)
					result = append(result, eventLog.Result.Logs[0].Data)
				}
			}
		}
	}

	return result, nil
}
