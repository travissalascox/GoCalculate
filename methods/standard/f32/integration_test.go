package methods

import (
	"math"
	"testing"
)

func TestEuler1D(t *testing.T) {
	f := func(x, omega float32) float32 {
		return omega - x*x + 1
	}
	a := float32(0.0)
	b := float32(2.0)
	N := 10
	initValue := float32(0.5)
	result := Euler1D(a, b, N, initValue, f)
	if (result - 4.865784) > 0.000001 {
		t.Fail()
	}
}

func TestTrapezoidRule(t *testing.T) {
	f := func(x float32) float32 {
		return float32(math.Sin(float64(x)))
	}
	a := float32(0.0)
	b := float32(math.Pi / 4)
	result := TrapezoidRule(a, b, f)
	if (result - 0.2776801) > 0.0000001 {
		t.Fail()
	}
}

func TestSimpsonRule(t *testing.T) {
	f := func(x float32) float32 {
		return float32(math.Sin(float64(x)))
	}
	a := float32(0.0)
	b := float32(math.Pi / 4)
	result := SimpsonRule(a, b, f)
	if (result - 0.2929326) > 0.0000001 {
		t.Fail()
	}
}

func TestSimpson38Rule(t *testing.T) {
	f := func(x float32) float32 {
		return float32(math.Sin(float64(x)))
	}
	a := float32(0.0)
	b := float32(math.Pi / 4)
	result := Simpson38Rule(a, b, f)
	if (result - 0.2929107) > 0.0000001 {
		t.Fail()
	}
}

func TestBooleRule(t *testing.T) {
	f := func(x float32) float32 {
		return float32(math.Sin(float64(x)))
	}
	a := float32(0.0)
	b := float32(math.Pi / 4)
	result := BooleRule(a, b, f)
	if (result - 0.29289318) > 0.0000001 {
		t.Fail()
	}
}
