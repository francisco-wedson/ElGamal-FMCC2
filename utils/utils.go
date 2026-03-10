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
	Baseado em: https://en.wikipedia.org/wiki/Euler%27s_totient_function#Computing_Euler's_totient_function
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

/*
	Funçao que retorna o inverso modular de determinado valor em determinado modulo
	Baseado em: https://en.wikipedia.org/wiki/Modular_multiplicative_inverse#Computation
*/
func InversoModular(valor, modulo int) int {
	if mdc(valor, modulo) != 1 {
		return -1
	}
	return ExponenciacaoModular(valor, Totiente(modulo) - 1, modulo)
}

/*
	Funcao auxiliar para calcular o mdc entre dois inteiros
*/
func mdc(a, b int) int {
	var resto int

	for a % b > 0 {
		resto = a % b
		a = b
		b = resto
	}

	return b
}
