package btccharts

type Message struct {
	Symbol, Volume, Id, Timestamp, Price string
}

/*func (m Message) GetCSV() string interface {*/

/*}*/

/*{"volume": 10.0, "timestamp": 1365812265, "price": 113.0, "symbol": "virtexCAD", "id": 21913358}*/
func GetMessage(string) *Message {
	return new(Message)
}
