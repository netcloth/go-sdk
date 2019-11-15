# netcloth-chain Go SDK


- **client**: provide httpClient, LiteClient, RpcClient and TxClient for query or send transaction on NCH
- **keys**: implement KeyManage to manage private key and accounts
- **types**: common types
- **util**: define constant and common functions

# Install

## Requirement

Go version above 1.12

## Use go mod(recommend)

Add "github.com/netcloth/go-sdk" dependency into your go.mod file.

```go
require (
	github.com/netcloth/go-sdk latest
)
```

# Usage

## Key Manager

Before start using API, you should have some accounts which have unch tokens. then exporting keysotre file by nchcli tool through command below:
```cassandraql
nchcli keys export <account_name>

Enter passphrase to decrypt your key:
enter your passphrase of <account_name> account
Enter passphrase to encrypt the exported key:
enter passphrase to encrypt the keystore file which can be used to import keystore to sdk

e.g.:
nchcli keys export lucy

```

When you have a keystore file and corresponding passphrase, you should construct a Key Manager to help sign the transaction msg or verify signature. Key Manager is an Identity Manger to define who you are in the NCH

We provide follow construct functions to generate Key Manager(other keyManager will coming soon):

```go
NewKeyManager(keystoreFile string, passphrase string) (KeyManager, error)
```

Examples:

for keyStore:

```go
func TestNewKeyManager(t *testing.T) {
	if km, err := keys.NewKeyManager(config.KeyStoreFileAbsPath, config.KeyStorePasswd); err != nil {
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
	"github.com/netcloth/go-sdk/client"
)

client, err := client.NewNCHClient("/Users/sky/go/src/github.com/netcloth/go-sdk/config/sdk.yaml")
```

Note:
- `baseUrl`: should be lcd endpoint if you want to use liteClient
- `nodeUrl`: should be nch node endpoint, format is `tcp://host:port`
- `networkType`: `alphanet` or `mainnet`(mainnet will come later)

after you init nchClient, it include follow clients which you can use:

- `liteClient`: lcd client for NCH
- `rpcClient`: query NCH info by rpc
- `txClient`: send transaction on NCH

