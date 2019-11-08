package client

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/NetCloth/netcloth-chain/modules/auth"
	sdk "github.com/NetCloth/netcloth-chain/types"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

type ReqServiceInfo struct {
	Type    string `json:"type" yaml:"type"`
	Address string `json:"address" yaml:"address"`
}

type ServiceInfo struct {
	Type    uint64 `json:"type" yaml:"type"`
	Address string `json:"address" yaml:"address"`
}

type ReqADParamType struct {
	UserAddress string         `json:"user_address" yaml:"user_address"`
	ServiceInfo ReqServiceInfo `json:"service_info" yaml:"service_info"`
	Expiration  time.Time      `json:"expriration"`
}

type ADParamType struct {
	UserAddress string      `json:"user_address" yaml:"user_address"`
	ServiceInfo ServiceInfo `json:"service_info" yaml:"service_info"`
	Expiration  time.Time   `json:"expriration"`
}

type ReqCIPALUserRequest struct {
	Params ReqADParamType `json:"params"`
	Sig    MySignature    `json:"signature"`
}

type CIPALUserRequest struct {
	Params ADParamType `json:"params"`
	Sig    MySignature `json:"signature"`
}

type MySignature struct {
	PubKey    string `json:"pub_key"`
	Signature string `json:"signature"`
}

func (p ReqADParamType) GetSignBytes() []byte {
	b, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func Test_SigVerify1(t *testing.T) {
	var userReq ReqCIPALUserRequest
	cliMsg := "{\"params\":{\"expriration\":\"2019-11-08T20:18:44.524268Z\",\"service_info\":{\"address\":\"nch19vnsnnseazkyuxgkt0098gqgvfx0wxmv96479m\",\"type\":\"1\"},\"user_address\":\"nch1edyenjf04mrsq3ueghmmga35hjeetgudjp4z36\"},\"signature\":{\"pub_key\":\"0x028c36b163c26f492abc874648b7258450394fe78133bcc4d920895d0ce8c3ac4e\",\"signature\":\"0x8271c4c00774e3de49f468367e610caf07c1319a1f0b8724f427ab1e918b703d429834765e8c9e8b7333919ced0ff68c28fea433e557816e06815447106d68be\"}}"
	cliMsgBZ := []byte(cliMsg)

	err := json.Unmarshal(cliMsgBZ, &userReq)
	if err != nil {
		fmt.Printf("read json failed: %s\n", err.Error())
		return
	}
	fmt.Printf("userReq = %v\n", userReq)

	var secp256k1PubKey secp256k1.PubKeySecp256k1
	pubKey, _ := hex.DecodeString(userReq.Sig.PubKey[2:])
	copy(secp256k1PubKey[:], pubKey)
	signature, _ := hex.DecodeString(userReq.Sig.Signature[2:])
	fmt.Printf("pubkey = %x\n", pubKey)
	fmt.Printf("sig = %x\n", signature)

	stdSig := auth.StdSignature{
		PubKey:    secp256k1PubKey,
		Signature: signature,
	}

	var params ReqADParamType
	params.ServiceInfo.Type = userReq.Params.ServiceInfo.Type
	params.ServiceInfo.Address = userReq.Params.ServiceInfo.Address
	params.UserAddress = userReq.Params.UserAddress
	params.Expiration = userReq.Params.Expiration

	fmt.Printf("final params = %s\n", string(params.GetSignBytes()))
	p := stdSig.VerifyBytes(params.GetSignBytes(), stdSig.Signature)
	fmt.Printf("p = %v\n", p)
}
