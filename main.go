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

// m.Post("/login", binding.Json(User{}), userLogin)

func GenerateDeterministicKeyPair(ctx *macaron.Context) {
	res := CipherData{}
	res.Seed = make([]byte, 23)
	rand.Read(res.Seed)
	p, s := cipher.GenerateDeterministicKeyPair(res.Seed)
	res.PubKey = []byte(p[:])
	res.SecKey = []byte(s[:])
	ctx.JSON(http.StatusOK, res)
}

/*
func PostECDH(ctx *macaron.Context, data CipherData) {
	c := db.C(ExpenseCollectionName)
	id := bson.ObjectIdHex(ctx.Params("id"))

	filter := bson.M{"_id": id}
	if ok, _ := inArray("admin", user.Roles); !ok {
		filter["owner"] = user.ID
	}
	var exp Expense
	err := c.Find(filter).One(&exp)
	if err != nil {
		sendError(ctx, http.StatusNotFound, err)
		return
	}

	err = c.Remove(filter)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err)
	} else {
		ctx.JSON(http.StatusOK, exp)
	}
}
*/
