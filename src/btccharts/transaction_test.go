package btccharts

import "testing"

func TestGetMessage(t *testing.T){
    const input = `{"volume": 4.0, "timestamp": 1365812301, "price": 114.0, "symbol": "virtexCAD", "id": 21913359}`
    m := GetMessage(input)
    if m.Volume != "4.0" || m.Timestamp != "1365812301" || m.Price != "114.0" || m.Symbol != "virtexCAD" || m.Id != "21913359" {
        t.Errorf("Fail %#v", m)
    }
}
