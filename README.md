# netcloth-chain Go SDK


- **client**: provide httpClient, LiteClient, RpcClient and TxClient for query or send transaction on NCH
- **keys**: implement KeyManage to manage private key and accounts
- **types**: common types
- **util**: define constant and common functions

# Install

## Requirement

Go version above 1.12

## Use go mod(recommend)

Add "github.com/NetCloth/go-sdk" dependency into your go.mod file.

```go
require (
	github.com/NetCloth/go-sdk latest
)
```

# Usage

## Key Manager

Before start using API, you should construct a Key Manager to help sign the transaction msg or verify signature. Key Manager is an Identity Manger to define who you are in the NCH

Wo provide follow construct functions to generate Key Mange(other keyManager will coming soon):

```go
NewKeyStoreKeyManager(file string, auth string) (KeyManager, error)
```

Examples:

for keyStore:

```go
func TestNewKeyStoreKeyManager(t *testing.T) {
	file := ks_1234567890.json
	if km, err := NewKeyStoreKeyManager(file, "1234567890"); err != nil {
		t.Fatal(err)
	} else {
		msg := []byte("hello world")
		signature, err := km.GetPrivKey().Sign(msg)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(km.GetAddr().String())


		assert.Equal(t, km.GetPrivKey().PubKey().VerifyBytes(msg, signature), true)
	}
}
```

## Init Client

```go
import (
	"github.com/NetCloth/go-sdk/client"
	"github.com/NetCloth/go-sdk/types"
)
var (
	baseUrl, nodeUrl string
	networkType = types.Alphanet
)
km, _ := keys.NewKeyStoreKeyManager("../keys/ks_1234567890.json", "1234567890")
c, _ := client.NewNCHClient(baseUrl, nodeUrl, networkType, km)
```

Note:
- `baseUrl`: should be lcd endpoint if you want to use liteClient
- `nodeUrl`: should be nch node endpoint, format is `tcp://host:port`
- `networkType`: `alphanet` or `mainnet`(mainnet will come later)

after you init nchClient, it include follow clients which you can use:

- `liteClient`: lcd client for NCH
- `rpcClient`: query NCH info by rpc
- `txClient`: send transaction on NCH

