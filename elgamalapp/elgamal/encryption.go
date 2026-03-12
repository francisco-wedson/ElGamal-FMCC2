package elgamal

import (
	"crypto/rand"
	"elgamalapp/utils"
	"math/big"
)

func GeradorChavePublica(generator int64, p int64, intermediate int64) (*big.Int, *big.Int) {

	privateKey, _ := rand.Int(rand.Reader, big.NewInt(p))

	var publicKey = big.NewInt(utils.ExponenciacaoModular(intermediate, privateKey.Int64(), p))
	var cipherText1 = big.NewInt(utils.ExponenciacaoModular(generator, privateKey.Int64(), p))

	return publicKey, cipherText1
}

func Encriptacao(generator *big.Int, keySize int, prime *big.Int, intermediate *big.Int, mensagem string) (*big.Int, []*big.Int) {

	K, c1 := GeradorChavePublica(generator.Int64(), prime.Int64(), intermediate.Int64())

	var c2 []*big.Int

	mensagemBytes := StringToInt(mensagem)

	for _, letra := range mensagemBytes {
		//falta converter chr para int

		c2 = append(c2, CodificaDigito(*big.NewInt(letra), *K, *prime))
	}

	return c1, c2
}

func CodificaDigito(digito big.Int, chavePublica big.Int, p big.Int) *big.Int {
	var mult big.Int
	mult.Mul(&digito, &chavePublica)

	return mult.Rem(&mult, &p)
}

func StringToInt(String string) []int64 {
	//1. Converter string para sequencia de bytes
	var StringBytes []int64
	for _, char := range String {
		StringBytes = append(StringBytes, int64(char))
	}

	return StringBytes
}
