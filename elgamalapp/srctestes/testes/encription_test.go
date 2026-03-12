package testes

import (
	"elgamalapp/elgamal"
	"elgamalapp/utils"
	"testing"
	"testing/cryptotest"
)

func SetUp(t *testing.T) {
	cryptotest.SetGlobalRandom(t, uint64(32767))
}

func TestGeradorChavePublica(t *testing.T) {
	SetUp(t)

	p, _ := utils.GerarPrimo(32)
	// teste com g=11, privateKeyB=5
	gotKey, _ := elgamal.GeradorChavePublica(11, p.Int64(), 55)

	if gotKey.Int64() != 1519133793 {
		t.Errorf("publicKey = %v; want publicKey = 1519133793", gotKey.Int64())
	}
}
