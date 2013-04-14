package neuroph

import "math/big"

// Calculates the MaxMin function to reduce numbers to some borders
// It's hardcoded to C = 0 & D =1
// B = (A - min⁡(A)) / (max⁡(A) - min⁡(A)) * ( D - C )+ C
//          x       / y                 *   z = 1 
func MaxMin(maxA, minA, a *big.Rat) *big.Rat {
	x := Sub(a, minA)
	y := Sub(maxA, minA)
	return Div(x, y)
}

// Returns a big.Rat with value 0
func NewRat() *big.Rat {
	return big.NewRat(0, 1)
}

// Wrapper for adding two values
func Add(x, y *big.Rat) *big.Rat {
	return NewRat().Add(x, y)
}

// Wrapper for dividing two values
func Div(x, y *big.Rat) *big.Rat {
	yinv := y.Inv(y)
	z := x.Mul(x, yinv)

	return z
}

// Wrapper for multiplying two values
func Mul(x, y *big.Rat) *big.Rat {
	return NewRat().Mul(x, y)
}

// Wraper for substracting two values
func Sub(x, y *big.Rat) *big.Rat {
	return NewRat().Sub(x, y)
}
