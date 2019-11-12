package client

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/netcloth/netcloth-chain/modules/cipal"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

type MySignature struct {
	PubKey    string `json:"pub_key"`
	Signature string `json:"signature"`
}

type CIPALUserRequest struct {
	Params cipal.ADParam `json:"params"`
	Sig    MySignature   `json:"signature"`
}

//sdk API for user
/*
	bz: client req json
		e.g.:"{\"params\":{\"expiration\":\"2019-11-13T06:28:32.279212Z\",\"service_info\":{\"address\":\"nch1ugus2df3sydca3quula5yjqfntuq5aaxweezpt\",\"type\":1},\"user_address\":\"nch1ugus2df3sydca3quula5yjqfntuq5aaxweezpt\"},\"signature\":{\"pub_key\":\"02b2bf9a87dd7cf1ad998721ffef00713a4d5fb2bae0316eea04268ae877a0bcac\",\"signature\":\"c2b822b4ddfbce95feb44112bce0022bbce67126b9fd50447e9ce3e0d03000ac287698c4994a892834921d5c2e3fe901d13f078eb54b289286831afb8046afe0\"}}"
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
