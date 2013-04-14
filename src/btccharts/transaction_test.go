package btccharts

import "testing"

// Example json input row from btccharts
const input = `{"volume": 4.0, "timestamp": 1365812301, "price": 114.0, "symbol": "mtgoxUSD", "id": 21913359}`

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
