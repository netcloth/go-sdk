package tx

import (
	"fmt"

	"github.com/NetCloth/go-sdk/client/lcd"
	"github.com/NetCloth/go-sdk/client/rpc"
	"github.com/NetCloth/go-sdk/client/types"
	"github.com/NetCloth/go-sdk/keys"
	"github.com/NetCloth/go-sdk/util/constant"

	commontypes "github.com/NetCloth/go-sdk/types"
)

type TxClient interface {
	SendToken(receiver string, coins []types.Coin, memo string, commit bool) (types.BroadcastTxResult, error)
}

type client struct {
	chainId    string
	keyManager keys.KeyManager
	liteClient lcd.LiteClient
	rpcClient  rpc.RPCClient
}

func NewClient(chainId string, networkType commontypes.NetworkType, keyManager keys.KeyManager,
	liteClient lcd.LiteClient, rpcClient rpc.RPCClient) (TxClient, error) {
	var (
		network string
	)
	switch networkType {
	case commontypes.Mainnet:
		network = constant.NetworkTypeMainnet
	case commontypes.Alphanet:
		network = constant.NetworkTypeAlphanet
	default:
		return &client{}, fmt.Errorf("invalid networktype, %d", networkType)
	}

	fmt.Println(network)
	// sdktypes.SetNetworkType(network)

	return &client{
		chainId:    chainId,
		keyManager: keyManager,
		liteClient: liteClient,
		rpcClient:  rpcClient,
	}, nil
}
