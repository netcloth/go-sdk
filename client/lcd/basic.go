package lcd

import (
	"github.com/netcloth/go-sdk/client/basic"
)

type LiteClient interface {
	QueryAccount(address string) (AccountBody, error)

	QueryIPALList() (IPALListBody, error)
	QueryIPALChatServerEndpoints() ([]string, error)
	QueryIPALEndpointsByType(endpointType string) ([]string, error)

	QueryIPALByAddress(address string) (IPALBody, error)
	QueryIPALByUNCompressedPubKey(uncompressedPubKey string) (IPALBody, error)
	QueryIPALChatServerEndpointByUNCompressedPubKey(uncompressedPubKey string) (string, error)
	QueryIPALChatServersEndpointByAddresses(addresses []string) (map[string]string, error)
	QueryIPALsEndpointByAddressesByType(addresses []string, endpointType string) (map[string]string, error)

	QueryCIPALByAddress(address string) (CIPALBody, error)
	QueryCIPALByUNCompressedPubKey(uncompressedPubKey string) (CIPALBody, error)
	QueryCIPALChatServerAddrByUNCompressedPubKey(uncompressedPubKey string) (string, error)
	QueryCIPALChatServersAddrByUNCompressedPubKeys(uncompressedPubKey []string) (map[string]string, error)
	QueryCIPALsAddrByUNCompressedPubKeysByType(uncompressedPubKey []string, endpointType string) (map[string]string, error)

	QueryContractLog(txId []byte) (ContractLog, error)
}

type client struct {
	httpClient basic.HttpClient
}

func NewClient(c basic.HttpClient) LiteClient {
	return &client{
		httpClient: c,
	}
}
