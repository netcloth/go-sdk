package sigtest

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	sdk "github.com/NetCloth/netcloth-chain/types"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

type ServiceInfo struct {
	Type    uint64 `json:"type" yaml:"type"`
	Address string `json:"address" yaml:"address"`
}

type ADParamType struct {
	UserAddress string      `json:"user_address" yaml:"user_address"`
	ServiceInfo ServiceInfo `json:"service_info" yaml:"service_info"`
	Expiration  time.Time   `json:"expiration"`
}

type MySignature struct {
	PubKey    string `json:"pub_key"`
	Signature string `json:"signature"`
}

type CIPALUserRequest struct {
	Params ADParamType `json:"params"`
	Sig    MySignature `json:"signature"`
}

//sdk API for user
/*
	bz: client req json
		e.g.:"{\"params\":{\"expriration\":\"2019-11-08T20:18:44.524268Z\",\"service_info\":{\"address\":\"nch19vnsnnseazkyuxgkt0098gqgvfx0wxmv96479m\",\"type\":1},\"user_address\":\"nch1edyenjf04mrsq3ueghmmga35hjeetgudjp4z36\"},\"signature\":{\"pub_key\":\"028c36b163c26f492abc874648b7258450394fe78133bcc4d920895d0ce8c3ac4e\",\"signature\":\"8271c4c00774e3de49f468367e610caf07c1319a1f0b8724f427ab1e918b703d429834765e8c9e8b7333919ced0ff68c28fea433e557816e06815447106d68be\"}}"
*/
func CIPALClaimReqFromBytes(bz []byte) (req CIPALUserRequest, pubkey secp256k1.PubKeySecp256k1, signature []byte, err error) {
	err = json.Unmarshal(bz, &req)
	if err != nil {
		fmt.Printf("parse bz failed: %s\n", err)
		return
	}

	pubKeyHex, err := hex.DecodeString(req.Sig.PubKey)
	if err != nil {
		fmt.Printf("parse pubkey failed: %s\n", err)
		return
	}
	copy(pubkey[:], pubKeyHex)

	signature, err = hex.DecodeString(req.Sig.Signature)
	if err != nil {
		fmt.Printf("parse signature failed: %s\n", err)
		return
	}

	return
}

func (p ADParamType) GetSignBytes() []byte {
	b, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func Test_SignatureVerify(t *testing.T) {
	//cliMsg为客户端发给服务端的json字符串，保证公钥和地址的开头不要有0x，service_info.type是数字1不是字符串"1"，cliMsgWrong是错误的例子
	//cliMsgWrong := "{\"params\":{\"expriration\":\"2019-11-08T20:18:44.524268Z\",\"service_info\":{\"address\":\"nch19vnsnnseazkyuxgkt0098gqgvfx0wxmv96479m\",\"type\":\"1\"},\"user_address\":\"nch1edyenjf04mrsq3ueghmmga35hjeetgudjp4z36\"},\"signature\":{\"pub_key\":\"0x028c36b163c26f492abc874648b7258450394fe78133bcc4d920895d0ce8c3ac4e\",\"signature\":\"0x8271c4c00774e3de49f468367e610caf07c1319a1f0b8724f427ab1e918b703d429834765e8c9e8b7333919ced0ff68c28fea433e557816e06815447106d68be\"}}"
	//cliMsg因为是手动把cliMsgWrong的错误格式改正了，实际的签名信息是cliMsgWrong，所以签名验证还是会失败
	cliMsg := "{\"params\":{\"expriration\":\"2019-11-08T20:18:44.524268Z\",\"service_info\":{\"address\":\"nch19vnsnnseazkyuxgkt0098gqgvfx0wxmv96479m\",\"type\":1},\"user_address\":\"nch1edyenjf04mrsq3ueghmmga35hjeetgudjp4z36\"},\"signature\":{\"pub_key\":\"028c36b163c26f492abc874648b7258450394fe78133bcc4d920895d0ce8c3ac4e\",\"signature\":\"8271c4c00774e3de49f468367e610caf07c1319a1f0b8724f427ab1e918b703d429834765e8c9e8b7333919ced0ff68c28fea433e557816e06815447106d68be\"}}"
	cliMsgBZ := []byte(cliMsg)
	userReq, pubkey, signature, err := CIPALClaimReqFromBytes(cliMsgBZ)
	if err != nil {
		fmt.Printf("parse client req failed: %s\n", err.Error())
		return
	}
	fmt.Printf("userReq = %v\n", userReq)
	fmt.Printf("pubkey = %v\n", pubkey)
	fmt.Printf("signature = %x\n", signature)
	fmt.Printf("final params = %s\n", string(userReq.Params.GetSignBytes()))

	p := pubkey.VerifyBytes(userReq.Params.GetSignBytes(), signature)
	if p {
		fmt.Printf("verfiy signature PASS")
	} else {
		fmt.Printf("verfiy signature FAIL")
	}
}
