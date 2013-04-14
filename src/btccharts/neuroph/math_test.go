package neuroph

import "testing"
import "math/big"

func TestAdd(t *testing.T) {
	x := big.NewRat(24, 1)
	y := big.NewRat(6, 1)
	z := big.NewRat(30, 1)
	result := Add(x, y)

	if z.Cmp(result) != 0 {
		t.Errorf("Expected: %s\nResult: %s\n", z, result)
	}

}

func TestDiv(t *testing.T) {
	x := big.NewRat(24, 1)
	y := big.NewRat(6, 1)
	z := big.NewRat(4, 1)
	result := Div(x, y)

	if z.Cmp(result) != 0 {
		t.Errorf("Expected: %s\nResult: %s\n", z, result)
	}

}

func TestMul(t *testing.T) {
	x := big.NewRat(4, 1)
	y := big.NewRat(6, 1)
	z := big.NewRat(24, 1)
	result := Mul(x, y)

	if z.Cmp(result) != 0 {
		t.Errorf("Expected: %s\nResult: %s\n", z, result)
	}

}

func TestSub(t *testing.T) {
	x := big.NewRat(24, 1)
	y := big.NewRat(6, 1)
	z := big.NewRat(18, 1)
	result := Sub(x, y)

	if z.Cmp(result) != 0 {
		t.Errorf("Expected: %s\nResult: %s\n", z, result)
	}

}

func TestMaxMin(t *testing.T) {
	a := big.NewRat(3, 1)
	minA := big.NewRat(1, 1)
	maxA := big.NewRat(10, 1)
	expected := big.NewRat(2, 9)
	result := MaxMin(maxA, minA, a)
	if expected.Cmp(result) != 0 {
		t.Errorf("Expected: %s\nResult: %s\n", expected, result)
	}
}
