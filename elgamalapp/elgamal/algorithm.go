package elgamal

import (
	"crypto/rand"
	"encoding/base64"
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

func Encriptacao(priv *PrivateKey, mensagem string) (string, []string) {

	K, _ := rand.Int(rand.Reader, new(big.Int).Sub(priv.P, big.NewInt(1)))

	c1_bytes := new(big.Int).Exp(priv.G, K, priv.P).Bytes()

	var c2 []string

	mensagem_bytes := StringToInt(mensagem)

	for _, letra := range mensagem_bytes {
		encoded_digit := CodificaDigito(big.NewInt(letra), *priv, K).Bytes()
		c2 = append(c2, base64.StdEncoding.EncodeToString(encoded_digit))
	}

	c1 := base64.StdEncoding.EncodeToString(c1_bytes)

	return c1, c2
}

func CodificaDigito(digito *big.Int, priv PrivateKey, K *big.Int) *big.Int {

	mult := new(big.Int).Mul(digito, new(big.Int).Exp(priv.Y, K, priv.P))

	return mult.Rem(mult, priv.P)
}

func StringToInt(String string) []int64 {
	var StringBytes []int64
	for _, char := range String {
		StringBytes = append(StringBytes, int64(char))
	}

	return StringBytes
}

func Decriptacao(priv *PrivateKey, c1 string, c2 []string) string {
	c1_deco_base64, _ := base64.StdEncoding.DecodeString(c1)
	c1_deco_bytes := new(big.Int).SetBytes(c1_deco_base64)

	//Calcula o segredo S
	s := new(big.Int).Exp(c1_deco_bytes, priv.X, priv.P)
	s_mod_inverse := new(big.Int).ModInverse(s, priv.P)

	var decrypted_message []*big.Int

	for _, value := range c2 {
		value_deco_base64, _ := base64.StdEncoding.DecodeString(value)
		value_deco_bytes := new(big.Int).SetBytes(value_deco_base64)

		// Isola-se o M e multiplica o valor de c2 pelo inverso modular de S (mod p)
		M := new(big.Int).Mul(s_mod_inverse, value_deco_bytes)
		M.Mod(M, priv.P)

		decrypted_message = append(decrypted_message, M)
	}

	// Recuperando a mensagem, transformando cada unicode em string novamente
	message := IntToString(decrypted_message)

	return message
}

func IntToString(decrypted_message []*big.Int) string {
	message := ""
	for _, value := range decrypted_message {
		message += string(rune(value.Int64()))
	}
	return message
}
