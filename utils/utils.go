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

func InversoModular(valor, modulo int) int {
	
}
