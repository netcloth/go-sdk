package client

import (
	"fmt"
	"testing"
)

func Test_SignatureVerify(t *testing.T) {
	//cliMsg为客户端发给服务端的json字符串，保证公钥和地址的开头不要有0x，service_info.type是数字1不是字符串"1"，cliMsgWrong是错误的例子
	cliMsg := "{\"params\":{\"expiration\":\"2019-11-13T06:28:32.279212Z\",\"service_info\":{\"address\":\"nch1ugus2df3sydca3quula5yjqfntuq5aaxweezpt\",\"type\":1},\"user_address\":\"nch1ugus2df3sydca3quula5yjqfntuq5aaxweezpt\"},\"signature\":{\"pub_key\":\"02b2bf9a87dd7cf1ad998721ffef00713a4d5fb2bae0316eea04268ae877a0bcac\",\"signature\":\"c2b822b4ddfbce95feb44112bce0022bbce67126b9fd50447e9ce3e0d03000ac287698c4994a892834921d5c2e3fe901d13f078eb54b289286831afb8046afe0\"}}"
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
