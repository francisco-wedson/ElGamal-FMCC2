package testes

import (
	"elgamalapp/elgamal"
	utils "elgamalapp/utilspackage"
	"encoding/base64"
	"errors"
	"math/big"
	"testing"
	"testing/cryptotest"
)

func SetUp(t *testing.T) (private *elgamal.PrivateKey) {
	cryptotest.SetGlobalRandom(t, 1)

	priv := &elgamal.PrivateKey{
		PublicKey: elgamal.PublicKey{
			G: big.NewInt(11),
			P: utils.GerarPrimo(512),
		},
		X: big.NewInt(62555),
	}

	return priv
}

func TestGeradorChavePublica(t *testing.T) {
	priv := SetUp(t)

	// teste com g=11, privateKeyB=5
	priv.Y, _ = elgamal.GeradorChavePublica(*priv)

	gotKey64 := base64.StdEncoding.EncodeToString(priv.Y.Bytes())

	expectedKey := "Yhpwcg/+95tAy128feAc3dA3KOFv0ebZbiyj4iqmCkzyZlH5e95WL5e5pzVumyM2SN6cVDfOn8OjQzysqNC8Fw=="

	if gotKey64 != expectedKey {
		// A conversão só é feita nos testes para facilitar a leitura, no sistema, todos são tratados como big.Int
		t.Errorf("publicKey = %v; want publicKey = %s", gotKey64, expectedKey)
	}
}

func TestGeraPublicKeyArgumentosInvalidos(t *testing.T) {
	priv := SetUp(t)

	priv.X = big.NewInt(1)

	_, err := elgamal.GeradorChavePublica(*priv)

	if err == errors.New("PrivateKey X must be: X > 1") {
		t.Errorf("got %s", err)
	}
	priv.X = big.NewInt(62555)
	priv.G = big.NewInt(1)

	_, err2 := elgamal.GeradorChavePublica(*priv)

	if err2 == errors.New("Cannot generate group elements with G=1") {
		t.Errorf("got %s", err2)
	}
}
