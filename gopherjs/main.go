package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/skycoin/skycoin/src/cipher"
)

func main() {
	js.Global.Set("Cipher", map[string]interface{}{
		"GenerateDeterministicKeyPair": cipher.GenerateDeterministicKeyPair,
		"ECDH": cipher.ECDH,
		"BitcoinAddressFromPubkey": cipher.BitcoinAddressFromPubkey,
		"AddressFromSecKey":        cipher.AddressFromSecKey,
	})
}
