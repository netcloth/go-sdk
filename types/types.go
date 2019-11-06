package types

import (
	"github.com/NetCloth/netcloth-chain/modules/auth"
	"github.com/NetCloth/netcloth-chain/modules/bank"
	"github.com/NetCloth/netcloth-chain/modules/cipal"
	"github.com/NetCloth/netcloth-chain/modules/ipal"
)

type NetworkType int

const (
	_ NetworkType = iota
	Alphanet
	Mainnet
)

type (
	MsgSend             = bank.MsgSend
	MsgServiceNodeClaim = ipal.MsgServiceNodeClaim
	MsgIPALClaim        = cipal.MsgIPALClaim
	StdFee              = auth.StdFee
)
