package lcd

const (
	UriQueryAccount      = "/auth/accounts/%s"
	UriQueryCIPAL        = "/cipal/query/%s"
	UriQueryCIPALs       = "/cipal/batch_query"
	UriQueryIPAL         = "/ipal/node/%s"
	UriQueryIPALs        = "/ipal/nodes"
	UriQueryIPALList     = "/ipal/list"
	UriQueryContractLogs = "/vm/logs/%s"
)

const (
	EndpointTypeServerChat      = "1"
	EndpointTypeClientChat      = "1"
	EndpointTypeClientGroupChat = "2"
)

var ClientChatEndpointType2ServerChatEndpointType = map[string]string{
	EndpointTypeClientChat:      EndpointTypeServerChat,
	EndpointTypeClientGroupChat: EndpointTypeServerChat,
}
