package lcd

import (
	"github.com/NetCloth/go-sdk/client/basic"
)

type LiteClient interface {
	QueryAccount(address string) (AccountBody, error)
	QueryIPAL(address string) (IPALObj, error)
}

type client struct {
	httpClient basic.HttpClient
}

func NewClient(c basic.HttpClient) LiteClient {
	return &client{
		httpClient: c,
	}
}
