package constant

const (
	TxBroadcastTypeSync   = "sync"
	TxBroadcastTypeAsync  = "async"
	TxBroadcastTypeCommit = "commit"

	TxDefaultGas       = 200000
	TxDefaultFeeAmount = 0
	TxDefaultDenom     = "unch"

	ChainID             = "nch-prinet-sky"
	NetworkTypeMainnet  = "mainnet"
	NetworkTypeAlphanet = "nch-alphanet"

	KeyStoreFileAbsPath = "/Users/sky/go/src/github.com/netcloth/go-sdk/keystorefile/keystore"

	LiteClientRpcEndpoint = "http://127.0.0.1:1317"
	RPCEndpoint           = "http://127.0.0.1:26657"
)
