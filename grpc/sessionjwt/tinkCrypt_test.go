package sessionjwt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTinkCrypt_DecryptStr(t *testing.T) {
	tinkSecure := []byte("SecurekeyforCryptBytes")
	tk, err := NewTinkCrypt(tinkSecure)
	assert.NoError(t, err)
	assert.NotNil(t, tk)
	cryptStr := []byte("this xxxxxxxxxx  data needs to be encrypted")
	ct, er2 := tk.Encrypt(cryptStr)
	assert.NoError(t, er2)

	decryptStr, er3 := tk.Decrypt(ct)
	assert.NoError(t, er3)

	assert.Equal(t, cryptStr, decryptStr)
}

func TestTinkCrypt_DecryptString(t *testing.T) {
	tinkSecure := []byte("SecurekeyforCryptBytes")
	tk, err := NewTinkCrypt(tinkSecure)
	assert.NoError(t, err)
	assert.NotNil(t, tk)
	cryptStr := "this xxxxxxxxxx  data needs to be encrypted"
	ct, er2 := tk.EncryptString(cryptStr)
	assert.NoError(t, er2)

	decryptStr, er3 := tk.DecryptString(ct)
	assert.NoError(t, er3)

	assert.Equal(t, cryptStr, decryptStr)
}

func BenchmarkDecryptStr(b *testing.B) {
	b.ReportAllocs()

	tinkSecure := []byte("SecurekeyforCryptBytes")
	tink, _ := NewTinkCrypt(tinkSecure)

	cryptStr := []byte("this xxxxxxxxxx  data needs to be encrypted")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ct, _ := tink.Encrypt(cryptStr)

		tink.Decrypt(ct)

	}
}
