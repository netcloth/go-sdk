package tx

import (
	"github.com/NetCloth/netcloth-chain/codec"
	"github.com/NetCloth/netcloth-chain/modules/aipal"
	"github.com/NetCloth/netcloth-chain/modules/auth"
	"github.com/NetCloth/netcloth-chain/modules/bank"
	"github.com/NetCloth/netcloth-chain/modules/ipal"
	sdk "github.com/NetCloth/netcloth-chain/types"

	"github.com/tendermint/go-amino"
)

var Cdc *amino.Codec

func init() {
	var cdc = codec.New()
	bank.RegisterCodec(cdc)
	auth.RegisterCodec(cdc)
	ipal.RegisterCodec(cdc)
	aipal.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)

	Cdc = cdc
}
