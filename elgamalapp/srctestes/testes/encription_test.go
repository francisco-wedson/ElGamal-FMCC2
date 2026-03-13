package testes

import (
	"elgamalapp/elgamal"
	utils "elgamalapp/utilspackage"
	"encoding/base64"
	"errors"
	"fmt"
	"math/big"
	"testing"
	"testing/cryptotest"
)

func SetUp(t *testing.T) (private *elgamal.PrivateKey) {
	cryptotest.SetGlobalRandom(t, 1)

	priv := &elgamal.PrivateKey{
		PublicKey: elgamal.PublicKey{
			// G é o elemento gerador da cifra, para simplificar os testes, foi utilizado um número pequeno
			G: big.NewInt(11),
			// P é um inteiro primo, neste caso, de 512 bits. O tamanho ideal para prevenir força bruta é de 1024 bits, porém, para simplificar os testes,
			// Adotamos P = 12717467131655530610417060098829663556118770331061190366206317456715657043637509246214103888453050204608792542575140380471386394942330068327343550542513527
			P: utils.GerarPrimo(512),
		},
		// G é a chave privada fixa da cifra, para simplificar os testes, foi utilizado um número pequeno
		X: big.NewInt(10),
	}

	return priv
}

func TestGeradorChavePublica(t *testing.T) {
	priv := SetUp(t)

	fmt.Println(priv.P)
	// Y é um valor calculado a partir da fórmula Y = g^x mod p, onde g é o gerador, x é a chave priv. fixa e p é um primo grande
	priv.Y, _ = elgamal.GeradorChavePublica(*priv)

	// utilizamos a codificação base64 para tornar o resultado curto e fácil de ser testado
	gotKey64 := base64.StdEncoding.EncodeToString(priv.Y.Bytes())

	expectedKey := "Bgn9sNk="

	if gotKey64 != expectedKey {
		t.Errorf("publicKey = %v; want publicKey = %s", gotKey64, expectedKey)
	}
}

func TestGeraPublicKeyArgumentosInvalidos(t *testing.T) {
	priv := SetUp(t)
	// Não deve ser permitido que a chave privada seja menor que ou igual a 1, pois g^1 mod p é congruente a g mod p, tornando a cifra muito simples, e
	// g^0 mod p é congruente a 1 mod p, tornando a cifra completamente ineficaz, visto que a cifragem é realizada com o produto do valor codificado
	// da mensagem, e valor_msg * 1 mod p é congruente à própria mensagem.
	priv.X = big.NewInt(1)

	_, err := elgamal.GeradorChavePublica(*priv)

	if err == errors.New("PrivateKey X must be: X > 1") {
		t.Errorf("got %s", err)
	}
	priv.X = big.NewInt(62555)
	// Assim como X, G não pode ser menor que ou igual a um, pois 1^X mod P é congruente a 1 mod P, e isso torna a cifra a própria mensagem;
	// Se G = 0, temos 0^X mod P, como a cifra é dada por um produto, seria impossível decodificar a mensagem, se seu valor cifrado fosse 0.
	priv.G = big.NewInt(1)

	_, err2 := elgamal.GeradorChavePublica(*priv)

	if err2 == errors.New("Cannot generate group elements with G=1") {
		t.Errorf("got %s", err2)
	}
}

func TestCriptografarMensagem(t *testing.T) {
	priv := SetUp(t)
	priv.Y, _ = elgamal.GeradorChavePublica(*priv)

	mensagem := "msg"
	c1, c2 := elgamal.Encriptacao(priv, mensagem)

	expected_c1 := "0WhV8iOPAeNaNTPfkcu2iqstEpuH/mVdoV9Kmquo9S65izQzSLz+AqiCVIRVAcyJHU1nJldpFuEqFaNYflcC7A=="

	expected_c2 := []string{"y6SP8ZfWJ55Br1PhRlEeAXB+25WvZ/FKwWC28n8kf4aUyMFUiMgR39kLdReEmrI8kLt7kzZgyZ7pUoY3NpPnbA==", "Z3e13+528Vt1ao2Upj63TucQ2ULvDlXDmTmL/ekiDyDG3HoX8SHxMUeceOcsOAMNhopJtKEODDyOVbaCma+Fng==", "PP+6ygBROeBIxAOpQ43UbAiCTdpldtpBBQgrvwaN4Vp2V0vUz6l4hlsRXC2RhFNAWbza3C27OFpXt+dq0w5Dww=="}

	if c1 != expected_c1 {
		t.Errorf("got %s; want %s", c1, expected_c1)
	}
	for i := 0; i < len(c2); i++ {
		if c2[i] != expected_c2[i] {
			t.Errorf("got %s; want %s", c2[i], expected_c2[i])
		}
	}
}
