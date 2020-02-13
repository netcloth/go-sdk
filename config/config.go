package config

import (
	"errors"
	"fmt"
	"io/ioutil"

	oconfig "github.com/olebedev/config"
)

const (
	DefautlKeyStoreFileAbsPath   = "/Users/zhuliting/go/nch-sdk/config/keystore"
	DefaultKeyStorePasswd        = "12345678"
	DefaultLiteClientRpcEndpoint = "http://127.0.0.1:1317"
	DefaultRPCEndpoint           = "http://127.0.0.1:26657"
	DefaultChainID               = "nch-chain"
	DefaultTxDefaultGas          = uint64(200000)
	DefaultTxDefaultFeeAmount    = int64(500000)
)

var (
	KeyStoreFileAbsPath   = DefautlKeyStoreFileAbsPath
	KeyStorePasswd        = DefaultKeyStorePasswd
	LiteClientRpcEndpoint = DefaultLiteClientRpcEndpoint
	RPCEndpoint           = DefaultRPCEndpoint
	ChainID               = DefaultChainID
	TxDefaultGas          = DefaultTxDefaultGas
	TxDefaultFeeAmount    = DefaultTxDefaultFeeAmount
)

func Init(sdkConfigFileAbsPath string) {
	data, err := ioutil.ReadFile(sdkConfigFileAbsPath)
	if err != nil {
		panic(err)
	}

	SDK, err := oconfig.ParseYaml(string(data))
	if err != nil {
		panic(err)
	}

	KeyStoreFileAbsPath, err = SDK.String("keystore.KeyStoreFileAbsPath")
	if err != nil {
		panic(err)
	}

	KeyStorePasswd, err = SDK.String("keystore.KeyStorePasswd")
	if err != nil {
		panic(err)
	}

	LiteClientRpcEndpoint, err = SDK.String("endpoint.LiteClientRpcEndpoint")
	if err != nil {
		panic(err)
	}

	LiteClientRpcEndpoint, err = SDK.String("endpoint.LiteClientRpcEndpoint")
	if err != nil {
		panic(err)
	}

	ChainID, err = SDK.String("ChainID")
	if err != nil {
		panic(err)
	}

	tmpTxDefaultGas, err := SDK.Int("feeParams.TxDefaultGas")
	if err != nil {
		panic(err)
	}

	if tmpTxDefaultGas <= 0 {
		panic(errors.New(fmt.Sprintf("feeParams.TxDefaultGas must > 0")))
	}
	TxDefaultGas = uint64(tmpTxDefaultGas)

	tmpTxDefaultFeeAmount, err := SDK.Int("feeParams.TxDefaultFeeAmount")
	if err != nil {
		panic(err)
	}

	if tmpTxDefaultFeeAmount <= 0 {
		panic(errors.New(fmt.Sprintf("feeParams.tmpTxDefaultFeeAmount must > 0")))
	}
	TxDefaultFeeAmount = int64(tmpTxDefaultFeeAmount)
}
