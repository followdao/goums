package sessionjwt

import (
	"github.com/google/tink/go/aead"
	"github.com/google/tink/go/keyset"
	"github.com/google/tink/go/tink"
	"github.com/shengdoushi/base58"

	"github.com/tsingson/goums/pkg/vtils"
)

// TinkCrypt struct for tink crypt ahead
type TinkCrypt struct {
	secure []byte
	ahead  tink.AEAD
}

// NewTinkCrypt   new
func NewTinkCrypt(secure []byte) (*TinkCrypt, error) {
	kh, err := keyset.NewHandle(aead.AES256GCMKeyTemplate())
	if err != nil {
		return nil, err
	}

	a, er1 := aead.New(kh)
	if er1 != nil {
		return nil, er1
	}
	t := &TinkCrypt{
		secure: secure,
		ahead:  a,
	}
	return t, nil
}

func (t *TinkCrypt) Encrypt(source []byte) ([]byte, error) {
	return t.ahead.Encrypt(source, t.secure)
}

func (t *TinkCrypt) EncryptString(source string) (string, error) {
	s := vtils.S2B(source)

	outByte, err := t.ahead.Encrypt(s, t.secure)
	if err != nil {
		return "", err
	}
	return base58.Encode(outByte, base58.RippleAlphabet), nil
}

func (t *TinkCrypt) Decrypt(source []byte) ([]byte, error) {
	return t.ahead.Decrypt(source, t.secure)
}

func (t *TinkCrypt) DecryptString(source string) (string, error) {
	cryptByte, err := base58.Decode(source, base58.RippleAlphabet)
	if err != nil {
		return "", err
	}
	sourceByte, er1 := t.ahead.Decrypt(cryptByte, t.secure)
	if er1 != nil {
		return "", er1
	}

	return vtils.B2S(sourceByte), nil
}
