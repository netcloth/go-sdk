package keys

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/netcloth/go-sdk/util/constant"
)

func TestNewKeyManager(t *testing.T) {
	if km, err := NewKeyManager(constant.KeyStoreFileAbsPath, "12345678"); err != nil {
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
