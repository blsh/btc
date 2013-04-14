package logger

import "testing"
import "math/big"
import "btccharts/neuroph"
import "btccharts"

const input = `{"volume": 4.0, "timestamp": 1365812301, "price": 114.0, "symbol": "mtgoxUSD", "id": 21913359}`

func TestAddMisisingValues(t *testing.T) {
	expected := getNormalizedInput()
	result := AddMissingValues(btccharts.GetMessage(input).Normalize())
	for k, v := range expected {
		if result[k].Cmp(v) != 0 {
			t.Errorf("Expected: %s\nResult: %s\n", v, result[k])
		}
	}

}

// This is an example of a row of data for neuroph
func ExampleNormalizedInput1() []*big.Rat {
	volumeMax := big.NewRat(1000000, 1)
	volumeMin := big.NewRat(1, 1000000)

	timestampMax := big.NewRat(2284438026, 1)
	timestampMin := big.NewRat(1116828426, 1)

	priceMax := big.NewRat(1000000, 1)
	priceMin := big.NewRat(1, 1000000)

	symbol := big.NewRat(1, 10)
	price := neuroph.MaxMin(priceMax, priceMin, big.NewRat(113, 1))
	volume := neuroph.MaxMin(volumeMax, volumeMin, big.NewRat(4, 1))
	timestamp := neuroph.MaxMin(timestampMax, timestampMin, big.NewRat(1365812180, 1))
	output1 := big.NewRat(0, 1)
	output2 := big.NewRat(0, 1)
	output3 := big.NewRat(1, 1)

	return []*big.Rat{symbol, price, volume, timestamp}
}
func ExampleNormalizedInput2() []*big.Rat {
	volumeMax := big.NewRat(1000000, 1)
	volumeMin := big.NewRat(1, 1000000)

	timestampMax := big.NewRat(2284438026, 1)
	timestampMin := big.NewRat(1116828426, 1)

	priceMax := big.NewRat(1000000, 1)
	priceMin := big.NewRat(1, 1000000)

	symbol := big.NewRat(1, 10)
	price := neuroph.MaxMin(priceMax, priceMin, big.NewRat(114, 1))
	volume := neuroph.MaxMin(volumeMax, volumeMin, big.NewRat(4, 1))
	timestamp := neuroph.MaxMin(timestampMax, timestampMin, big.NewRat(1365812301, 1))
	output1 := big.NewRat(0, 1)
	output2 := big.NewRat(0, 1)
	output3 := big.NewRat(1, 1)

	return []*big.Rat{symbol, price, volume, timestamp}
}
