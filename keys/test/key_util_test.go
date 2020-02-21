package test

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"

	sdk "github.com/netcloth/netcloth-chain/types"

	"github.com/netcloth/netcloth-chain/hexutil"

	btcsecp256k1 "github.com/btcsuite/btcd/btcec"

	"github.com/netcloth/go-sdk/keys"

	"github.com/tendermint/tendermint/crypto/secp256k1"

	"github.com/stretchr/testify/require"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"

	"github.com/tendermint/tendermint/crypto"
)

const (
	uncompressedPubKey = "048c36b163c26f492abc874648b7258450394fe78133bcc4d920895d0ce8c3ac4e61c6895d129a36c94e4a9a724b8061e32560eb8e4aa9435373fc93afa22425de"
	compressedPubKey   = "028c36b163c26f492abc874648b7258450394fe78133bcc4d920895d0ce8c3ac4e"
	address            = "6983fed7a7159b82b64a5cd5ce3b3adf6b6aa3f9"
	addressBech32      = "nch1dxpla4a8zkdc9dj2tn2uuwe6ma4k4glewcte7w"
)

func Test_UNCompressedPubKey2CompressedPubKey(t *testing.T) {
	pubKeyBytes, err := keys.UNCompressedPubKey2CompressedPubKey(uncompressedPubKey)
	require.True(t, err == nil)
	require.Equal(t, compressedPubKey, fmt.Sprintf("%x", pubKeyBytes))
}

func Test_UNCompressedPubKey2Address(t *testing.T) {
	addr, err := keys.UNCompressedPubKey2Address(uncompressedPubKey)
	require.True(t, err == nil)
	require.Equal(t, address, strings.ToLower(addr.String()))
}

func Test_UNCompressedPubKey2AddressBech32(t *testing.T) {
	addr, err := keys.UNCompressedPubKey2AddressBech32(uncompressedPubKey)
	require.True(t, err == nil)
	require.Equal(t, addressBech32, addr)
}

func Test_PubKeyHexString2AddressBech32(t *testing.T) {
	addr, err := keys.PubKeyHexString2AddressBech32(compressedPubKey)
	require.True(t, err == nil)
	require.Equal(t, addressBech32, addr)
}

func Test_PubKey2AddressBech32(t *testing.T) {
	pubKey, err := hex.DecodeString(compressedPubKey)
	require.True(t, err == nil)

	var pk secp256k1.PubKeySecp256k1
	copy(pk[:], pubKey)

	addrBech32, err := keys.PubKey2AddressBech32(pk)
	require.True(t, err == nil)
	require.Equal(t, addressBech32, addrBech32)
}

func Test_test(t *testing.T) {
	h1 := ethcrypto.Keccak256([]byte("abc"))
	h2 := crypto.Sha256([]byte("abc"))

	t.Log(h1, h2)
}

func Test_1(t *testing.T) {
	hash := crypto.Sha256([]byte("abdadfasdfadfcd"))
	hash1, err := hexutil.Decode("0xce0677bb30baa8cf067c88db9811f4333d131bf8bcf12fe7065d211dce971008")
	t.Log(fmt.Sprintf("%x", hash1))
	t.Log(fmt.Sprintf("%x", hash))

	pri, err := btcsecp256k1.NewPrivateKey(btcsecp256k1.S256())
	t.Log(fmt.Sprintf("%x", pri))
	t.Log(fmt.Sprintf("%x", pri.D))
	t.Log(err)
	sig, err := pri.Sign(hash)
	t.Log(fmt.Sprintf("%x", sig.Serialize()))
	t.Log(err)
	sig1, err := btcsecp256k1.SignCompact(btcsecp256k1.S256(), pri, hash, true)
	t.Log(fmt.Sprintf("%x", sig1))
	sig2, err := btcsecp256k1.SignCompact(btcsecp256k1.S256(), pri, hash, false)
	t.Log(fmt.Sprintf("%x", sig2))
	t.Log(fmt.Sprintf("%d:%d:%d\n", len(sig.Serialize()), len(sig1), len(sig2)))

	x, err := hexutil.Decode("2a6e636831787735396864307a74677a35366c6d307534716567336335767330726b6b72727775656e7072")
	t.Log(fmt.Sprintf("%s\n", string(x)))
	addr, err := sdk.AccAddressFromBech32("nch1xw59hd0ztgz56lm0u4qeg3c5vs0rkkrrwuenpr")

	t.Log(hexutil.Encode(addr.Bytes()))

}

func Test_2(t *testing.T) {
	hash, _ := hexutil.Decode("0xce0677bb30baa8cf067c88db9811f4333d131bf8bcf12fe7065d211dce971008")

	curve := btcsecp256k1.S256()

	pk, _ := btcsecp256k1.NewPrivateKey(curve)
	t.Log(fmt.Sprintf("pk %x", pk))

	pubkey := pk.PubKey()
	t.Log(fmt.Sprintf("pubkey: %x", pubkey.SerializeCompressed()))

	t.Log(fmt.Sprintf("uncompressed pubkey: %x", pubkey.SerializeUncompressed()))
	addr, _ := keys.UNCompressedPubKey2Address(fmt.Sprintf("%x", pubkey.SerializeUncompressed()))
	t.Log(fmt.Sprintf("addr = %v", addr))

	sig, _ := btcsecp256k1.SignCompact(curve, pk, hash, false)
	t.Log(fmt.Sprintf("sig: %x", sig))

	rpubkey, _, _ := btcsecp256k1.RecoverCompact(btcsecp256k1.S256(), sig, hash)
	t.Log(fmt.Sprintf("recovered pubkey: %x", rpubkey))
}
