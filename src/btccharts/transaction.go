// Package provides functions and structs for working with data from telnet
// btccharts.com
package btccharts

import "regexp"
import "strings"
import "io"
import "log"
import "fmt"
import "btccharts/neuroph"
import "encoding/json"
import "math/big"

// Represents the a json messsage row from btccharts
type Message struct {
	Symbol, Volume, Id, Timestamp, Price string
}

// This map hardcodes a trader string to a float:
var TraderIdMap = map[string]int64{
	"bit2cILS":    1,
	"bitfloorUSD": 2,
	"bitstampUSD": 3,
	"btcdeEUR":    4,
	"btceRUR":     5,
	"btceUSD":     6,
	"btcnCNY":     7,
	"mtgoxEUR":    8,
	"mtgoxPLN":    9,
	"mtgoxUSD":    10,
	"rippleXRP":   11,
	"virwoxSLL":   12,
	"virtexCAD":   13,
}

// The traderId hardcoded in TraderIdMap
func (m Message) GetTraderId() int64 {
	return TraderIdMap[m.Symbol]
}

// Returns the normalized string for neuroph
func (m Message) String() string {
	return fmt.Sprintf("%g,%s,%s,%s", m.GetTraderId(), m.Timestamp, m.Volume, m.Price)
}

// Parses a string from btccharts to json. We have to do some hacky reggex to
// make all values strings (fuck float64) before parsing it with std json go lib
func GetMessage(data string) Message {
	tmp := ConvertAllJsonValuesToString(data)
	dec := json.NewDecoder(strings.NewReader(tmp))
	var m Message
	for {
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}

	return m
}

func (m Message) Normalize() []*big.Rat {
	symbol := neuroph.NormalizeSymbol(big.NewRat(m.GetTraderId(), 100))
	price := neuroph.NormalizePrice(neuroph.RatFromString(m.Price))
	volume := neuroph.NormalizeVolume(neuroph.RatFromString(m.Volume))
	timestamp := neuroph.NormalizeTimestamp(neuroph.RatFromString(m.Timestamp))
	output1 := neuroph.NormalizePrice(neuroph.NewRat())
	result := []*big.Rat{symbol, price, volume, timestamp, output1, output1}
	return result
}

// Converts all json values to string to avoid handling with floats
func ConvertAllJsonValuesToString(data string) string {
	tmp := regexp.MustCompile(`: `).ReplaceAllString(data, ": \"")
	tmp = regexp.MustCompile(`[,]`).ReplaceAllString(tmp, "\",")
	tmp = regexp.MustCompile(`[}]`).ReplaceAllString(tmp, "\"}")
	tmp = regexp.MustCompile(`["]{2}`).ReplaceAllString(tmp, "\"")
	return tmp
}
