package elgamal

import (
	"errors"
	"math/big"
)

type PublicKey struct {
	G, P, Y *big.Int
}

type PrivateKey struct {
	PublicKey
	X *big.Int
}

func GeradorChavePublica(priv PrivateKey) (*big.Int, error) {

	if priv.X.Cmp(big.NewInt(1)) <= 0 {
		return nil, errors.New("PrivateKey X must be: X > 1")
	} else if priv.G.Cmp(big.NewInt(1)) <= 0 {
		return nil, errors.New("Cannot generate group elements with G=1")
	} else {
		Y := new(big.Int).Exp(priv.G, priv.X, priv.P)
		return Y, nil
	}
}

/*
func Encriptacao(generator big.Int, keySize int, prime big.Int, intermediate big.Int, mensagem string) (*big.Int, []*big.Int) {

		K, c1 := GeradorChavePublica(generator, prime, intermediate)

		var c2 []*big.Int

		mensagemBytes := StringToInt(mensagem)

		for _, letra := range mensagemBytes {
			//falta converter chr para int

			c2 = append(c2, CodificaDigito(*big.NewInt(letra), *K, prime))
		}

		return c1, c2
	}
*/
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
