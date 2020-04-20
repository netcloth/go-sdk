package keys

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/ripemd160"

	btcsecp256k1 "github.com/btcsuite/btcd/btcec"
	"github.com/netcloth/netcloth-chain/types"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

func UNCompressedPubKey2CompressedPubKey(uncompressedPubKey string) (compressedPubKey []byte, err error) {
	uncompressedPubKeyHex, err := hex.DecodeString(uncompressedPubKey)
	if err != nil {
		return nil, err
	}

	pubkey, err := btcsecp256k1.ParsePubKey(uncompressedPubKeyHex, btcsecp256k1.S256())
	if err != nil {
		return nil, err
	}

	return pubkey.SerializeCompressed(), nil
}

func CompressedPubKey2UNCompressedPubKey(compressedPubKey string) (unCompressedPubKey []byte, err error) {
	compressedPubKeyHex, err := hex.DecodeString(compressedPubKey)
	if err != nil {
		return nil, err
	}

	pubkey, err := btcsecp256k1.ParsePubKey(compressedPubKeyHex, btcsecp256k1.S256())
	if err != nil {
		return nil, err
	}

	return pubkey.SerializeUncompressed(), nil
}

func UNCompressedPubKey2Address(uncompressedPubKey string) (crypto.Address, error) {
	pubKey, err := UNCompressedPubKey2CompressedPubKey(uncompressedPubKey)
	if err != nil {
		return nil, err
	}

	hasherSHA256 := sha256.New()
	hasherSHA256.Write(pubKey[:])
	sha := hasherSHA256.Sum(nil)

	hasherRIPEMD160 := ripemd160.New()
	hasherRIPEMD160.Write(sha)
	return hasherRIPEMD160.Sum(nil), nil
}

func UNCompressedPubKey2AddressBech32(uncompressedPubKey string) (string, error) {
	addr, err := UNCompressedPubKey2Address(uncompressedPubKey)
	if err != nil {
		return "", err
	}

	return types.AccAddress(addr).String(), nil
}

func PubKeyHexString2AddressBech32(pubKeyStr string) (string, error) {
	if len(pubKeyStr) == 0 {
		return "", fmt.Errorf("pubkey invalid")
	}

	pubkeyHex, err := hex.DecodeString(pubKeyStr)
	if err != nil {
		return "", err
	}

	var pk secp256k1.PubKeySecp256k1
	copy(pk[:], pubkeyHex)
	addr := types.AccAddress(pk.Address().Bytes())

	return addr.String(), nil
}

func PubKeyHexString2Address(pubKeyStr string) (addr types.AccAddress, err error) {
	if len(pubKeyStr) == 0 {
		return addr, fmt.Errorf("pubkey invalid")
	}

	pubkeyHex, err := hex.DecodeString(pubKeyStr)
	if err != nil {
		return addr, err
	}

	var pk secp256k1.PubKeySecp256k1
	copy(pk[:], pubkeyHex)
	addr = types.AccAddress(pk.Address().Bytes())
	return
}

func PubKey2AddressBech32(pubKey crypto.PubKey) (string, error) {
	return types.AccAddress(pubKey.Address().Bytes()).String(), nil
}

func GetAccAddressByBechAddress(bech32addr string) (types.AccAddress, error) {
	return types.AccAddressFromBech32(bech32addr)
}
