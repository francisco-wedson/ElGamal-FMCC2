package elgamalpackage

import (
	"crypto/rand"
	"math"
	"math/big"

	"../utils"
)

func geradorChavePublica(generator big.Int, p big.Int, intermediate big.Int) (*big.Int, *big.Int) {
	privateKey, _ := rand.Int(rand.Reader, &p)

	var publicKeyInt = utils.ExponenciacaoModular(int(intermediate.Int64()), int(privateKey.Int64()), int(p.Int64()))
	var cipherText1Int = utils.ExponenciacaoModular(int(generator.Int64()), int(privateKey.Int64()), int(p.Int64()))
	var publicKey = big.NewInt(int64(publicKeyInt))
	var cipherText1 = big.NewInt(int64(cipherText1Int))

	return publicKey, cipherText1
}

func encriptacao(generator *big.Int, keySize int, intermediate *big.Int, mensagem string) (*big.Int, []*big.Int) {
	prime, _ := rand.Int(rand.Reader, big.NewInt(Pow(2, keySize)))

	K, c1 := geradorChavePublica(*generator, *prime, *intermediate)

	var c2 []*big.Int
	for _, chr := range mensagem {
		//falta converter chr para int
		c2 = append(c2, codificaDigito(big.Int(chr), *K, *prime))
	}

	return c1, c2
}

func codificaDigito(digito big.Int, chavePublica big.Int, p big.Int) *big.Int {
	var mult big.Int
	mult.Mul(&digito, &chavePublica)

	return mult.Rem(&mult, &p)
}

func Pow(x int, y int) int64 {
	return int64(math.Pow(float64(x), float64(y)))
}
