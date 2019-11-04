package client

import (
	"github.com/NetCloth/go-sdk/client/basic"
	"github.com/NetCloth/go-sdk/client/lcd"
	"github.com/NetCloth/go-sdk/client/rpc"
	"github.com/NetCloth/go-sdk/client/tx"
	"github.com/NetCloth/go-sdk/keys"
	"github.com/NetCloth/go-sdk/types"
)

type nchClient struct {
	basic.HttpClient
	lcd.LiteClient
	rpc.RPCClient
	tx.TxClient
}

type NCHClient interface {
	basic.HttpClient
	lcd.LiteClient
	rpc.RPCClient
	tx.TxClient
}

func NewNCHClient(baseUrl, nodeUrl string, networkType types.NetworkType, km keys.KeyManager) (NCHClient, error) {
	var (
		ic nchClient
	)
	basicClient := basic.NewClient(baseUrl)
	liteClient := lcd.NewClient(basicClient)
	rpcClient := rpc.NewClient(nodeUrl)
	status, err := rpcClient.GetStatus()
	if err != nil {
		return ic, err
	}
	chainId := status.NodeInfo.Network
	txClient, err := tx.NewClient(chainId, networkType, km, liteClient, rpcClient)
	if err != nil {
		return ic, err
	}

	ic = nchClient{
		HttpClient: basicClient,
		LiteClient: liteClient,
		RPCClient:  rpcClient,
		TxClient:   txClient,
	}

	return ic, nil
}
