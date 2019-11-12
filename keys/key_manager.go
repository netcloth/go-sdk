package keys

import (
	"fmt"
	"io/ioutil"

	btcsecp256k1 "github.com/btcsuite/btcd/btcec"
	"github.com/netcloth/go-sdk/types/tx"
	"github.com/netcloth/netcloth-chain/crypto/keys/mintkey"
	"github.com/netcloth/netcloth-chain/modules/auth"
	"github.com/netcloth/netcloth-chain/types"
	ctypes "github.com/netcloth/netcloth-chain/types"
	"github.com/pkg/errors"
	"github.com/tendermint/tendermint/crypto"
)

type KeyManager interface {
	Sign(msg tx.StdSignMsg) ([]byte, error)
	SignBytes(msg []byte) ([]byte, error)
	GetPrivKey() crypto.PrivKey
	GetAddr() types.AccAddress

	GetUCPubKey() (UCPubKey []byte, err error)
}

type keyManager struct {
	privKey crypto.PrivKey
	addr    types.AccAddress
}

func NewKeyManager(file string, auth string) (KeyManager, error) {
	k := keyManager{}
	err := k.ImportKeystore(file, auth)
	return &k, err
}

func (k *keyManager) Sign(msg tx.StdSignMsg) ([]byte, error) {
	sig, err := k.makeSignature(msg)
	if err != nil {
		return nil, err
	}

	newTx := auth.NewStdTx(msg.Msgs, msg.Fee, []auth.StdSignature{sig}, msg.Memo)
	bz, err := tx.Cdc.MarshalBinaryLengthPrefixed(newTx)
	if err != nil {
		return nil, err
	}

	return bz, nil
}

func (k *keyManager) SignBytes(msg []byte) ([]byte, error) {
	return k.privKey.Sign(msg)
}

func (k *keyManager) GetPrivKey() crypto.PrivKey {
	return k.privKey
}

func (k *keyManager) GetAddr() types.AccAddress {
	return k.addr
}

func (k *keyManager) GetUCPubKey() (UCPubKey []byte, err error) {
	pubkey, err := btcsecp256k1.ParsePubKey(k.GetPrivKey().PubKey().Bytes()[5:], btcsecp256k1.S256())
	if err != nil {
		return nil, err
	}

	return pubkey.SerializeUncompressed(), nil
}

func (k *keyManager) makeSignature(msg tx.StdSignMsg) (sig auth.StdSignature, err error) {
	if err != nil {
		return
	}
	sigBytes, err := k.privKey.Sign(msg.Bytes())
	if err != nil {
		return
	}
	return auth.StdSignature{
		PubKey:    k.privKey.PubKey(),
		Signature: sigBytes,
	}, nil
}

func (k *keyManager) ImportKeystore(keystoreFile string, passphrase string) error {
	if passphrase == "" {
		return fmt.Errorf("Password is missing ")
	}

	armor, err := ioutil.ReadFile(keystoreFile)
	if err != nil {
		return err
	}

	privKey, err := mintkey.UnarmorDecryptPrivKey(string(armor), passphrase)
	if err != nil {
		return errors.Wrap(err, "couldn't import private key")
	}

	addr := ctypes.AccAddress(privKey.PubKey().Address())
	k.addr = addr
	k.privKey = privKey
	return nil
}
