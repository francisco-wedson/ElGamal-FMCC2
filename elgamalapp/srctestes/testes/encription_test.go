package testes

import (
	"elgamalapp/elgamal"
	"elgamalapp/utils"
	"fmt"
	"math/big"
	"testing"
	"testing/cryptotest"
)

func SetUp(t *testing.T) {
	cryptotest.SetGlobalRandom(t, uint64(1))
}

func TestGeradorChavePublica(t *testing.T) {
	SetUp(t)

	p, _ := utils.GerarPrimo(1024)
	fmt.Print(p)
	// teste com g=11, privateKeyB=5
	gotKey, _ := elgamal.GeradorChavePublica(*big.NewInt(11), p, *big.NewInt(55))

	if gotKey.Int64() != 1519133793 {
		t.Errorf("publicKey = %v; want publicKey = 1519133793", gotKey)
	}
}
