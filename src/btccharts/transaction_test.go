package btccharts

import "testing"

import "math/big"
import "btccharts/neuroph"

// Example json input row from btccharts
const input = `{"volume": 4.0, "timestamp": 1365812301, "price": 114.0, "symbol": "mtgoxUSD", "id": 21913359}`

// Example strings like you get from the telnet bitcoincharts stream
var TradeStringsMap = []string{
	`{"volume": 4.0, "timestamp": 1365812301, "price": 114.0, "symbol": "mtgoxUSD", "id": 21913359}`,
	`{"volume": 3.2, "timestamp": 1365812329, "price": 115.0, "symbol": "btceUSD", "id": 21913360}`,
	`{"volume": 2.4, "timestamp": 1365812421, "price": 114.5, "symbol": "mtgoxUSD", "id": 21913361}`,
}

/*expected := []*big.Rat{symbol, price, volume, timestamp}*/

var ExpectedNeurophResults = []string{
	`0.10,0,000113999998,0`,
}

func TestGetMessage(t *testing.T) {
	m := GetMessage(input)
	if m.Volume != "4.0" || m.Timestamp != "1365812301" || m.Price != "114.0" || m.Symbol != "mtgoxUSD" || m.Id != "21913359" {
		t.Errorf("Fail %#v", m)
	}
}

func TestConvertAllJsonValuesToString(t *testing.T) {
	const expected = `{"volume": "4.0", "timestamp": "1365812301", "price": "114.0", "symbol": "mtgoxUSD", "id": "21913359"}`
	result := ConvertAllJsonValuesToString(input)
	checkStrings(expected, result, t)
}

func TestGetTraderId(t *testing.T) {
	var expected int64
	expected = 10
	result := GetMessage(input).GetTraderId()
	if expected != result {
		t.Errorf("Expected: %d\nResult: %d\n", expected, result)
	}
}

func TestMessageString(t *testing.T) {
	t.Logf("TODO: REIMPLEMENT TestMessageString")
	/*expected := "0.1,1365812301,4.0,114.0"*/
	/*result := fmt.Sprintf("%s", GetMessage(input))*/
	/*checkStrings(expected, result, t)*/
}

func TestNormalize(t *testing.T) {
	// Some magic constants! Carefull! Here shall be dragons!
	// 
	// We do here the assumption that there will be no trade for more than 1M
	// and less 1/1M bitcoin
	volumeMax := big.NewRat(1000000, 1)
	volumeMin := big.NewRat(1, 1000000)

	// We do the assumption that first trade was made on some sunny day in may
	// 2005 and the last will be in may 2042 HAHAHA >:D
	timestampMax := big.NewRat(2284438026, 1)
	timestampMin := big.NewRat(1116828426, 1)

	// We predict that at the high point of the hype a BTC will be worth
	// 1M $CURRENCY (RUB perhaps? :D) and the lowest point will be at 
	// 1/1M $CURRENCY
	priceMax := big.NewRat(1000000, 1)
	priceMin := big.NewRat(1, 1000000)
	symbol := big.NewRat(1, 10)
	price := neuroph.MaxMin(priceMax, priceMin, big.NewRat(114, 1))
	volume := neuroph.MaxMin(volumeMax, volumeMin, big.NewRat(4, 1))
	timestamp := neuroph.MaxMin(timestampMax, timestampMin, big.NewRat(1365812301, 1))

	expected := []*big.Rat{symbol, price, volume, timestamp}
	result := GetMessage(input).Normalize()
	for k, v := range expected {
		if result[k].Cmp(v) != 0 {
			t.Errorf("Expected: %s\nResult: %s\n", v, result[k])
		}
	}

}

func checkStrings(expected, result string, t *testing.T) {
	if expected != result {
		t.Errorf("Expected: %s\nResult: %s\n", expected, result)
	}
}

// Checks if our id map is in limits 1 - 99
func TestTraderIdMap(t *testing.T) {
	for k, v := range TraderIdMap {
		if v > 99 || v < 1 {
			t.Errorf("Trader %s has id %i out of limit 1-99", k, v)
		}
	}
}

// This is one of the most important test. It takes the strings from TradeStringsMap
// and tries to normalize it for neuroph. It checks the results with
// ExpectedNeurophResults
func TestMessageToNeuroph(t *testing.T) {

}
