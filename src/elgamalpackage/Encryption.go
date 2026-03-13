package elgamalpackage

import "crypto/rand"

func geradorChavePublica(p int, elemGerador int) [3]int {
	var privateKey int = rand.Int()

	var h int = ExponenciacaoModular(elemGerador, privateKey, p)

	var publicKey [3]int = [3]int{p, elemGerador, h}

	return publicKey
}

func encriptacao(keyLength int, elemGerador int, Message int) {}
