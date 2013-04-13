// You can edit this code!
// Click here and start typing.
package main

import (
	"btccharts"
	"btccharts/logger"
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "bitcoincharts.com:27007")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	reader := bufio.NewReader(conn)
	lw := logger.NewLogWriter()
	for true {
		line, _ := reader.ReadString('\n')
		lw.Write(line)
		m := btccharts.GetMessage(line)
		if m.Symbol == "mtgoxUSD" {
			fmt.Printf("%s,%s\n", m, m.Price)
		}

	}

	defer conn.Close()

}
