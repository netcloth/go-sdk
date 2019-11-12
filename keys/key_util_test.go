package keys

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"

	"github.com/tendermint/tendermint/crypto/secp256k1"

	"github.com/stretchr/testify/require"
)

const (
	uncompressedPubKey = "048c36b163c26f492abc874648b7258450394fe78133bcc4d920895d0ce8c3ac4e61c6895d129a36c94e4a9a724b8061e32560eb8e4aa9435373fc93afa22425de"
	compressedPubKey   = "028c36b163c26f492abc874648b7258450394fe78133bcc4d920895d0ce8c3ac4e"
	address            = "6983fed7a7159b82b64a5cd5ce3b3adf6b6aa3f9"
	addressBech32      = "nch1dxpla4a8zkdc9dj2tn2uuwe6ma4k4glewcte7w"
)

func Test_UNCompressedPubKey2CompressedPubKey(t *testing.T) {
	pubKeyBytes, err := UNCompressedPubKey2CompressedPubKey(uncompressedPubKey)
	require.True(t, err == nil)
	require.Equal(t, compressedPubKey, fmt.Sprintf("%x", pubKeyBytes))
}

func Test_UNCompressedPubKey2Address(t *testing.T) {
	addr, err := UNCompressedPubKey2Address(uncompressedPubKey)
	require.True(t, err == nil)
	require.Equal(t, address, strings.ToLower(addr.String()))
}

func Test_UNCompressedPubKey2AddressBech32(t *testing.T) {
	addr, err := UNCompressedPubKey2AddressBech32(uncompressedPubKey)
	require.True(t, err == nil)
	require.Equal(t, addressBech32, addr)
}

func Test_PubKeyHexString2AddressBech32(t *testing.T) {
	addr, err := PubKeyHexString2AddressBech32(compressedPubKey)
	require.True(t, err == nil)
	require.Equal(t, addressBech32, addr)
}

func Test_PubKey2AddressBech32(t *testing.T) {
	pubKey, err := hex.DecodeString(compressedPubKey)
	require.True(t, err == nil)

	var pk secp256k1.PubKeySecp256k1
	copy(pk[:], pubKey)

	addrBech32, err := PubKey2AddressBech32(pk)
	require.True(t, err == nil)
	require.Equal(t, addressBech32, addrBech32)
}
