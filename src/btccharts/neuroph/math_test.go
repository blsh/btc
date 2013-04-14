package neuroph

import "testing"
import "math/big"

// Contains few example of price changes with expected normalization results
// Example price is going up 1CUR
// {
//		"priceOld": big.NewRat(100, 1),
//		"priceNew": big.NewRat(101, 1),
//		"output1":  big.NewRat(1, 50),
//		"output2":  big.NewRat(0, 1),
//		"output3":  big.NewRat(0, 1),
//	}
var PriceChangesMap = [][]*big.Rat{
	// Price is going up 1CUR
	{big.NewRat(100, 1), big.NewRat(101, 1), big.NewRat(1, 50), big.NewRat(0, 1), big.NewRat(0, 1)},
	// Price stayes same
	{big.NewRat(100, 1), big.NewRat(100, 1), big.NewRat(0, 1), big.NewRat(1, 1), big.NewRat(0, 1)},
	// Price is going down 1CUR
	{big.NewRat(101, 1), big.NewRat(100, 1), big.NewRat(0, 1), big.NewRat(0, 1), big.NewRat(1, 50)},
}

var Normalized

func TestAdd(t *testing.T) {
	x := big.NewRat(24, 1)
	y := big.NewRat(6, 1)
	z := big.NewRat(30, 1)
	result := Add(x, y)
	checkRats(z, result, t)

}

func TestDiv(t *testing.T) {
	x := big.NewRat(24, 1)
	y := big.NewRat(6, 1)
	z := big.NewRat(4, 1)
	result := Div(x, y)
	checkRats(z, result, t)

}

func TestMul(t *testing.T) {
	x := big.NewRat(4, 1)
	y := big.NewRat(6, 1)
	z := big.NewRat(24, 1)
	result := Mul(x, y)
	checkRats(z, result, t)

}

func TestSub(t *testing.T) {
	x := big.NewRat(24, 1)
	y := big.NewRat(6, 1)
	z := big.NewRat(18, 1)
	result := Sub(x, y)
	checkRats(z, result, t)

}

func TestMaxMin(t *testing.T) {
	a := big.NewRat(3, 1)
	minA := big.NewRat(1, 1)
	maxA := big.NewRat(10, 1)
	expected := big.NewRat(2, 9)
	result := MaxMin(maxA, minA, a)
	checkRats(expected, result, t)
}

// Check our Magic!
func TestNormalizePrice(t *testing.T) {
	price := big.NewRat(114, 1)
	priceMax := big.NewRat(1000000, 1)
	priceMin := big.NewRat(1, 1000000)
	expected := MaxMin(priceMax, priceMin, price)
	result := NormalizePrice(price)

	checkRats(expected, result, t)
}

func TestNormalizeSymbol(t *testing.T) {
	symbolMax := big.NewRat(1, 1)
	symbolMin := big.NewRat(0, 1)
	symbol := big.NewRat(10, 100)
	expected := MaxMin(symbolMax, symbolMin, symbol)
	result := NormalizeSymbol(symbol)
	checkRats(expected, result, t)
}

func TestNormalizeVolume(t *testing.T) {
	volumeMax := big.NewRat(1000000, 1)
	volumeMin := big.NewRat(1, 1000000)
	volume := big.NewRat(4, 1)
	expected := MaxMin(volumeMax, volumeMin, volume)
	result := NormalizeVolume(volume)
	checkRats(expected, result, t)
}

func TestNormalizeTimestamp(t *testing.T) {
	timestampMax := big.NewRat(2284438026, 1)
	timestampMin := big.NewRat(1116828426, 1)
	timestamp := big.NewRat(1365812301, 1)
	expected := MaxMin(timestampMax, timestampMin, timestamp)
	result := NormalizeTimestamp(timestamp)
	checkRats(expected, result, t)
}

func TestRatFromString(t *testing.T) {
	expected := big.NewRat(75, 80)
	result := RatFromString("0.9375")
	checkRats(expected, result, t)
}

func checkRats(expected *big.Rat, result *big.Rat, t *testing.T) {
	if expected.Cmp(result) != 0 {
		t.Errorf("Expected: %s\nResult: %s\n", expected, result)
	}
}
