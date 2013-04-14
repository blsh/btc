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

func RatFromString(s string) *big.Rat {
	rat, ok := NewRat().SetString(s)
	if !ok {
		panic(ok)
	}
	return rat
}

// Some magic constants! Carefull! Here shall be dragons!
// All NormalizeFoo() functions have constants for fooMin & fooMax. foo is
// reduced to a value between 0 and 1

// We predict that at the high point of the hype a BTC will be worth
// 1M in some weird $CURRENCY (RUB perhaps? :D) and the lowest point will be at 
// 1/1M $CURRENCY
func NormalizePrice(price *big.Rat) *big.Rat {
	priceMax := big.NewRat(1000000, 1)
	priceMin := big.NewRat(1, 1000000)
	return MaxMin(priceMax, priceMin, price)
}

func NormalizeSymbol(symbol *big.Rat) *big.Rat {
	symbolMax := big.NewRat(1, 1)
	symbolMin := big.NewRat(0, 1)
	return MaxMin(symbolMax, symbolMin, symbol)
}

// We do here the assumption that there will be no trade for more than 1M
// and less 1/1M bitcoin
func NormalizeVolume(volume *big.Rat) *big.Rat {
	volumeMax := big.NewRat(1000000, 1)
	volumeMin := big.NewRat(1, 1000000)
	return MaxMin(volumeMax, volumeMin, volume)
}

// We do the assumption that first trade was made on some sunny day in may
// 2005 and the last will be in may 2042 HAHAHA >:D
func NormalizeTimestamp(timestamp *big.Rat) *big.Rat {
	timestampMax := big.NewRat(2284438026, 1)
	timestampMin := big.NewRat(1116828426, 1)
	return MaxMin(timestampMax, timestampMin, timestamp)
}
