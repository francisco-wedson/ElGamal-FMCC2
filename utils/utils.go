package utils

/*
	Essa funcao calcula o resultado da exponenciacao modular, recebendo a base, o expoente e o modulo, respectivamente e retornando o valor do da exponenciacao modular
	Baseado em: https://en.wikipedia.org/wiki/Modular_exponentiation#Memory-efficient_method
*/
func ExponenciacaoModular(base, expoente, modulo int) int {
	if modulo == 1 {
		return 0
	}
	var expoenteModular int = 1

	for expoenteTeste := 0; expoenteTeste < expoente; expoenteTeste++ {
		expoenteModular = (base * expoenteModular) % modulo
	}

	return expoenteModular
} 

/*
	Funcao Thotiente de Euler que calcula a quantidade de coprimos de um determinado valor inteiro
*/
func Totiente(n int) int {
    resultado := n
    valor := n

    for i := 2; i * i <= valor; i++ {
        if valor % i == 0 {
            resultado -= resultado / i
            for valor % i == 0 {
                valor /= i
            }
        }
    }

    if valor > 1 {
        resultado -= resultado / valor
    }

    return resultado
}
