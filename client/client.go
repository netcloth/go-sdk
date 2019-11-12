package client

import (
	"errors"
	"fmt"

	"github.com/netcloth/go-sdk/client/basic"
	"github.com/netcloth/go-sdk/client/lcd"
	"github.com/netcloth/go-sdk/client/rpc"
	"github.com/netcloth/go-sdk/client/tx"
	"github.com/netcloth/go-sdk/config"
	"github.com/netcloth/go-sdk/keys"
)

type NCHClient interface {
	basic.HttpClient
	lcd.LiteClient
	rpc.RPCClient
	tx.TxClient
}

type nchClient struct {
	basic.HttpClient
	lcd.LiteClient
	rpc.RPCClient
	tx.TxClient
}

func NewNCHClient() (nchClient, error) {
	var fake nchClient

	km, err := keys.NewKeyManager(config.KeyStoreFileAbsPath, config.KeyStorePasswd)
	if err != nil {
		panic(err)
	}

	basicClient := basic.NewClient(config.LiteClientRpcEndpoint)
	liteClient := lcd.NewClient(basicClient)
	rpcClient := rpc.NewClient(config.RPCEndpoint)

	status, err := rpcClient.GetStatus()
	if err != nil {
		return fake, err
	}

	if config.ChainID != status.NodeInfo.Network {
		return fake, errors.New(fmt.Sprintf("chainID dismatch:expected chainID[%s], actual chainID[%s]", config.ChainID, status.NodeInfo.Network))
	}

	txClient, err := tx.NewClient(status.NodeInfo.Network, km, liteClient, rpcClient)
	if err != nil {
		return fake, err
	}

	client := nchClient{
		HttpClient: basicClient,
		LiteClient: liteClient,
		RPCClient:  rpcClient,
		TxClient:   txClient,
	}

	return client, nil
}

func NewNCHTXClient() (tx.TxClient, error) {
	km, err := keys.NewKeyManager(config.KeyStoreFileAbsPath, config.KeyStorePasswd)
	if err != nil {
		panic(err)
	}

	basicClient := basic.NewClient(config.LiteClientRpcEndpoint)
	liteClient := lcd.NewClient(basicClient)
	rpcClient := rpc.NewClient(config.RPCEndpoint)

	status, err := rpcClient.GetStatus()
	if err != nil {
		return nil, err
	}

	if config.ChainID != status.NodeInfo.Network {
		return nil, errors.New(fmt.Sprintf("chainID dismatch:expected chainID[%s], actual chainID[%s]", config.ChainID, status.NodeInfo.Network))
	}

	client, err := tx.NewClient(status.NodeInfo.Network, km, liteClient, rpcClient)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func NewNCHQueryClient() (lcd.LiteClient, error) {
	basicClient := basic.NewClient(config.LiteClientRpcEndpoint)
	liteClient := lcd.NewClient(basicClient)
	rpcClient := rpc.NewClient(config.RPCEndpoint) //TODO check release?

	status, err := rpcClient.GetStatus()
	if err != nil {
		return nil, err
	}

	if config.ChainID != status.NodeInfo.Network {
		return nil, errors.New(fmt.Sprintf("chainID dismatch:expected chainID[%s], actual chainID[%s]", config.ChainID, status.NodeInfo.Network))
	}

	return liteClient, nil
}
