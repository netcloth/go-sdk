package tx

import (
	"github.com/netcloth/go-sdk/client/lcd"
	"github.com/netcloth/go-sdk/client/rpc"
	"github.com/netcloth/go-sdk/client/types"
	"github.com/netcloth/go-sdk/keys"
	"github.com/netcloth/netcloth-chain/modules/cipal"
	"github.com/netcloth/netcloth-chain/modules/ipal"
	sdk "github.com/netcloth/netcloth-chain/types"
)

type TxClient interface {
	keys.KeyManager
	SendToken(receiver string, coins []types.Coin, memo string, commit bool) (types.BroadcastTxResult, error)
	IPALClaim(Moniker, website, details string, endpoints ipal.Endpoints, bond sdk.Coin, commit bool) (types.BroadcastTxResult, error)
	CIPALClaim(req cipal.IPALUserRequest, memo string, commit bool) (types.BroadcastTxResult, error)
}

type client struct {
	keys.KeyManager
	chainId    string
	liteClient lcd.LiteClient
	rpcClient  rpc.RPCClient
}

func NewClient(chainId string, keyManager keys.KeyManager, liteClient lcd.LiteClient, rpcClient rpc.RPCClient) (TxClient, error) {
	return &client{
		chainId:    chainId,
		KeyManager: keyManager,
		liteClient: liteClient,
		rpcClient:  rpcClient,
	}, nil
}
