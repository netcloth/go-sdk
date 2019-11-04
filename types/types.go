package types

import (
	"github.com/NetCloth/netcloth-chain/modules/aipal"
	"github.com/NetCloth/netcloth-chain/modules/auth"
	"github.com/NetCloth/netcloth-chain/modules/bank"
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
	MsgServiceNodeClaim = aipal.MsgServiceNodeClaim
	MsgIPALClaim        = ipal.MsgIPALClaim
	StdFee              = auth.StdFee
)
