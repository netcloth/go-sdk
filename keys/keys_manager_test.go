package keys

import (
	"testing"

	"github.com/netcloth/go-sdk/config"

	"github.com/stretchr/testify/assert"
)

func TestNewKeyManager(t *testing.T) {
	if km, err := NewKeyManager(config.KeyStoreFileAbsPath, config.KeyStorePasswd); err != nil {
		t.Fatal(err)
	} else {
		msg := []byte("hello world")
		signature, err := km.GetPrivKey().Sign(msg)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(km.GetAddr().String())

		assert.Equal(t, km.GetPrivKey().PubKey().VerifyBytes(msg, signature), true)
	}
}
