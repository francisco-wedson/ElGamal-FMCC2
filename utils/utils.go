package utils

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
