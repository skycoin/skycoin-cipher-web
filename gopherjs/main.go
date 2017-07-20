package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/skycoin/skycoin/src/cipher"
)

type Cipher struct {
}

func GenerateDeterministicKeyPair(seed []byte) (cipher.PubKey, cipher.SecKey) {
	return cipher.GenerateDeterministicKeyPair(seed)
}

func NewCipher() *js.Object {
	return js.MakeWrapper(&Cipher{})
}

func main() {
	js.Global.Set("Cipher", map[string]interface{}{
		"GenerateDeterministicKeyPair": GenerateDeterministicKeyPair,
	})
	/*
		p, s := GenerateDeterministicKeyPair()
		js.Global.Set("public_key", p)
		js.Global.Set("private_key", s)
		ecdh := cipher.ECDH(p, s)

	*/
}
