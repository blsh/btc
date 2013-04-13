// Package provides functions and structs for working with data from telnet
// btccharts.com
package btccharts

import "regexp"
import "strings"
import "io"
import "log"
import "encoding/json"

type Message struct {
	Symbol, Volume, Id, Timestamp, Price string
}

/*func (m Message) GetCSV() string interface {*/

/*}*/

/*const input = "{\"volume\": 4.0, \"timestamp\": 1365812301, \"price\": 114.0, \"symbol\": \"virtexCAD\", \"id\": 21913359}"*/
func GetMessage(data string) Message {
    var tmp  string
    tmp = regexp.MustCompile(`: `).ReplaceAllString(data,": \"")
    tmp = regexp.MustCompile(`[,]`).ReplaceAllString(tmp,"\",")
    tmp = regexp.MustCompile(`[}]`).ReplaceAllString(tmp,"\"}")
    tmp = regexp.MustCompile(`["]{2}`).ReplaceAllString(tmp,"\"")
    log.Println(tmp)
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
