package types

import (
	"github.com/netcloth/netcloth-chain/modules/auth"
	"github.com/netcloth/netcloth-chain/modules/bank"
	"github.com/netcloth/netcloth-chain/modules/cipal"
	"github.com/netcloth/netcloth-chain/modules/ipal"
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
