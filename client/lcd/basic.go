package lcd

import (
	"github.com/netcloth/go-sdk/client/basic"
)

type LiteClient interface {
	QueryAccount(address string) (AccountBody, error)
	QueryCIPALByAddress(address string) (CIPALBody, error)
	QueryAIPALByAddress(address string) (AIPALBody, error)
	QueryAIPALList() (AIPALListBody, error)
}

type client struct {
	httpClient basic.HttpClient
}

func NewClient(c basic.HttpClient) LiteClient {
	return &client{
		httpClient: c,
	}
}
