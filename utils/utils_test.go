package utils

import (
	"testing"
)

func TestExponenciacaoModularBasica(t *testing.T) {
	got := ExponenciacaoModular(2, 3, 5)
	want := 3
	if got != want {
		t.Errorf("Got %d, but want %d\n", got, want)
	}
}

func TestExponenciacaoModularPonteciaNula(t *testing.T) {
	got := ExponenciacaoModular(123, 0, 7)
	want := 1
	if got != want {
		t.Errorf("Got %d, but want %d\n", got, want)
	}
}

func TestExponenciacaoModularBaseMaiorQueModulo(t *testing.T) {
	got := ExponenciacaoModular(10, 2, 7)
	want := 2
	if got != want {
		t.Errorf("Got %d, but want %d\n", got, want)
	}
}

func TestExponenciacaoModularOverflow(t *testing.T) {
	got := ExponenciacaoModular(2, 10, 1000)
	want := 24
	if got != want {
		t.Errorf("Got %d, but want %d\n", got, want)
	}
}

func TestExponenciacaoModularPequenoTeoremaDeFermat(t *testing.T) {
	got := ExponenciacaoModular(5, 1_000_000_006, 1_000_000_007)
	want := 1
	if got != want {
		t.Errorf("Got %d, but want %d\n", got, want)
	}
}

func TestExponenciacaoModularBaseModuloIguais(t *testing.T) {
	got := ExponenciacaoModular(7, 10, 7)
	want := 0
	if got != want {
		t.Errorf("Got %d, but want %d\n", got, want)
	}
}

func TestInversoModularBasicoPrimos(t *testing.T) {
	got := InversoModular(3, 7)
	want := 5
	if got != want {
		t.Errorf("Got %d, but want %d\n", got, want)
	}
}

func TestInversoModularModuloComposto(t *testing.T) {
	got := InversoModular(3, 10)
	want := 7
	if got != want {
		t.Errorf("Got %d, but want %d\n", got, want)
	}
}


func TestInversoModularInexistente(t *testing.T) {
	got := InversoModular(2, 4)
	want := -1
	if got != want {
		t.Errorf("Got %d, but want %d\n", got, want)
	}
}

func TestInversoModularLimiar(t *testing.T) {
	got := InversoModular(12, 13)
	want := 12
	if got != want {
		t.Errorf("Got %d, but want %d\n", got, want)
	}
}

func TestInversoModularOverflow(t *testing.T) {
	got := InversoModular(2, 1_000_000_007)
	want := 500000004
	if got != want {
		t.Errorf("Got %d, but want %d\n", got, want)
	}
}
