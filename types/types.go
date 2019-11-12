package types

import (
	"github.com/netcloth/netcloth-chain/modules/auth"
	"github.com/netcloth/netcloth-chain/modules/bank"
	"github.com/netcloth/netcloth-chain/modules/cipal"
	"github.com/netcloth/netcloth-chain/modules/ipal"
)

type (
	MsgSend             = bank.MsgSend
	MsgServiceNodeClaim = ipal.MsgServiceNodeClaim
	MsgIPALClaim        = cipal.MsgIPALClaim
	StdFee              = auth.StdFee
)
