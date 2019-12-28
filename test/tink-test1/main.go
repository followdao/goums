package main

import (
	"fmt"
	"log"

	"github.com/google/tink/go/aead"
	"github.com/google/tink/go/keyset"
	"github.com/savsgio/gotils"
	"github.com/shengdoushi/base58"
)

func main() {

	kh, err := keyset.NewHandle(aead.AES256GCMKeyTemplate())
	if err != nil {
		log.Fatal(err)
	}

	a, err := aead.New(kh)
	if err != nil {
		log.Fatal(err)
	}

	k := []byte("associated data")
	ct, err := a.Encrypt([]byte("this data needs to be encrypted"), k)
	if err != nil {
		log.Fatal(err)
	}

	pt, err := a.Decrypt(ct, k)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Cipher text: %s\nPlain text: %s\n", base58.Encode(ct, base58.RippleAlphabet), gotils.B2S(pt))

}
