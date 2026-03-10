package elgamalpackage

import (
	"crypto/rand"
	"math/big"
)

func geradorChavePublica(generator big.Int, p big.Int, intermediate big.Int) (big.Int, big.Int) {
	privateKey, _ := rand.Int(rand.Reader, &p)

	var publicKey = ExponeciacaoModular(intermediate, privateKey, p)
	var h = ExponenciacaoModular(generator, privateKey, p)

	return publicKey, h
}

func encriptacao(generator big.Int, p big.Int, intermediate big.Int, mensagem string) (big.Int, []big.Int) {
	chavePublica, c1 := geradorChavePublica(generator, p, intermediate)

	var c2 []big.Int
	for _, chr := range mensagem {
		//falta converter chr para int
		c2 = append(c2, codificaDigito(big.Int(chr), chavePublica, p))
	}

	return c1, c2
}

// mul esta com parametros errados
func codificaDigito(digito big.Int, chavePublica big.Int, p big.Int) big.Int {
	return (big.Int.Mul(digito, chavePublica) % p)
}
