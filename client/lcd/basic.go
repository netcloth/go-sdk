package lcd

import (
	"github.com/netcloth/go-sdk/client/basic"
)

type LiteClient interface {
	QueryAccount(address string) (AccountBody, error)

	QueryIPALList() (IPALListBody, error)
	QueryIPALChatServerEndpoints() ([]string, error)

	QueryIPALByAddress(address string) (IPALBody, error)
	QueryIPALByUNCompressedPubKey(uncompressedPubKey string) (IPALBody, error)
	QueryIPALChatServerEndpointByUNCompressedPubKey(uncompressedPubKey string) (string, error)

	QueryCIPALByAddress(address string) (CIPALBody, error)
	QueryCIPALByUNCompressedPubKey(uncompressedPubKey string) (CIPALBody, error)
	QueryCIPALChatServerAddrByUNCompressedPubKey(uncompressedPubKey string) (string, error)
}

type client struct {
	httpClient basic.HttpClient
}

func NewClient(c basic.HttpClient) LiteClient {
	return &client{
		httpClient: c,
	}
}
