package main

import (
	"log"
	"net/http"

	"crypto/rand"

	"github.com/go-macaron/binding"
	"github.com/go-macaron/gzip"
	"github.com/skycoin/skycoin/src/cipher"
	"gopkg.in/macaron.v1"
)

type CipherData struct {
	Seed         []byte `json:"seed,omitempty"`
	PubKey       []byte `json:"pub,omitempty"`
	SecKey       []byte `json:"sec,omitempty"`
	SharedSecret []byte `json:"sharedSecret,omitempty"`
	Address      string `json:"address,omitempty"`
}

func main() {
	m := macaron.New()
	m.Use(macaron.Logger())
	m.Use(macaron.Recovery())
	m.Use(gzip.Gziper())
	m.Use(macaron.Renderer())

	//m.Get("/ECDH", getECDH)
	m.Get("/GenerateDeterministicKeyPair", GenerateDeterministicKeyPair)
	m.Post("/GenerateDeterministicKeyPair", binding.Json(CipherData{}), PostGenerateDeterministicKeyPair)
	m.Post("/ECDH", binding.Json(CipherData{}), PostECDH)
	m.Post("/BitcoinAddressFromPubkey", binding.Json(CipherData{}), BitcoinAddressFromPubkey)

	m.Use(macaron.Static("web"))
	m.Use(macaron.Static("gopherjs"))

	log.Println("Server is running...")
	log.Println(http.ListenAndServe("127.0.0.1:4000", m))
}

func PostGenerateDeterministicKeyPair(ctx *macaron.Context, data CipherData) {
	p, s := cipher.GenerateDeterministicKeyPair(data.Seed)
	data.PubKey = []byte(p[:])
	data.SecKey = []byte(s[:])
	ctx.JSON(http.StatusOK, data)
}

func GenerateDeterministicKeyPair(ctx *macaron.Context) {
	res := CipherData{}
	res.Seed = make([]byte, 23)
	rand.Read(res.Seed)
	p, s := cipher.GenerateDeterministicKeyPair(res.Seed)
	res.PubKey = []byte(p[:])
	res.SecKey = []byte(s[:])
	ctx.JSON(http.StatusOK, res)
}

func PostECDH(ctx *macaron.Context, data CipherData) {
	p := cipher.PubKey{}
	s := cipher.SecKey{}
	for i, v := range data.PubKey {
		p[i] = v
	}
	for i, v := range data.SecKey {
		s[i] = v
	}
	data.SharedSecret = cipher.ECDH(p, s)
	ctx.JSON(http.StatusOK, data)
}

func BitcoinAddressFromPubkey(ctx *macaron.Context, data CipherData) {
	p := cipher.PubKey{}
	for i, v := range data.PubKey {
		p[i] = v
	}
	data.Address = cipher.BitcoinAddressFromPubkey(p)
	ctx.JSON(http.StatusOK, data)
}
