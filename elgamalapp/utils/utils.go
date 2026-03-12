package utils

import (
	"crypto/rand"
	"math/big"
)

/*
Essa funcao calcula o resultado da exponenciacao modular, recebendo a base, o expoente e o modulo, respectivamente e retornando o valor do da exponenciacao modular
Baseado em: https://en.wikipedia.org/wiki/Modular_exponentiation#Memory-efficient_method
*/
func ExponenciacaoModular(base, expoente, modulo int64) int64 {
	if modulo == 1 {
		return 0
	}
	var expoenteModular int64 = 1

	for expoenteTeste := 0; int64(expoenteTeste) < expoente; expoenteTeste++ {
		expoenteModular = (base * expoenteModular) % modulo
	}

	return expoenteModular
}

/*
Funcao Thotiente de Euler que calcula a quantidade de coprimos de um determinado valor inteiro
Baseado em: https://en.wikipedia.org/wiki/Euler%27s_totient_function#Computing_Euler's_totient_function
*/
func Totiente(n int64) int64 {
	resultado := n
	valor := n

	for i := 2; int64(i*i) <= valor; i++ {
		if valor%int64(i) == 0 {
			resultado -= resultado / int64(i)
			for valor%int64(i) == 0 {
				valor /= int64(i)
			}
		}
	}

	if valor > 1 {
		resultado -= resultado / valor
	}

	return resultado
}

/*
Funçao que retorna o inverso modular de determinado valor em determinado modulo
Baseado em: https://en.wikipedia.org/wiki/Modular_multiplicative_inverse#Computation
*/
func InversoModular(valor, modulo int64) int64 {
	if mdc(valor, modulo) != 1 {
		return -1
	}
	return ExponenciacaoModular(valor, Totiente(modulo)-1, modulo)
}

/*
Funcao auxiliar para calcular o mdc entre dois inteiros
*/
func mdc(a, b int64) int64 {
	var resto int64

	for a%b > 0 {
		resto = a % b
		a = b
		b = resto
	}

	return b
}

/*
Função auxiliar para gerar um número primo pseudo-aleatório, criptograficamente seguro
*/
func GerarPrimo(keySize int) (big.Int, error) {
	n, _ := rand.Prime(rand.Reader, keySize)

	return *n, nil
}
