package btccharts

import "testing"

// Example json input row from btccharts
const input = `{"volume": 4.0, "timestamp": 1365812301, "price": 114.0, "symbol": "virtexCAD", "id": 21913359}`

func TestGetMessage(t *testing.T) {
	m := GetMessage(input)
	if m.Volume != "4.0" || m.Timestamp != "1365812301" || m.Price != "114.0" || m.Symbol != "virtexCAD" || m.Id != "21913359" {
		t.Errorf("Fail %#v", m)
	}
}

func TestConvertAllJsonValuesToString(t *testing.T) {
	const expected = `{"volume": "4.0", "timestamp": "1365812301", "price": "114.0", "symbol": "virtexCAD", "id": "21913359"}`
	result := ConvertAllJsonValuesToString(input)
	checkStrings(expected, result, t)
}

func checkStrings(expected, result string, t *testing.T) {
	if expected != result {
		t.Errorf("Expected: %s\nResult: %s\n", expected, result)
	}
}
