// Logs the output of bitcoincharts.com to file and output
package main

import (
	"btccharts"
	"btccharts/logger"
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "bitcoincharts.com:27007")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	reader := bufio.NewReader(conn)
	logfile, err2 := os.OpenFile("/home/kalkin/projects/work/btc-trade/data/btccharts-data.log", os.O_APPEND|os.O_RDWR, 0666)
	if err2 != nil {
		panic(err2)
	}
	defer logfile.Close()
	lw := logger.NewLogWriter(logfile)
	for true {
		line, _ := reader.ReadString('\n')
		lw.Write(line)
		m := btccharts.GetMessage(line)
		timestamp, err3 := strconv.Atoi(m.Timestamp)

		if err3 != nil {
			panic(err3)
		}
		date := time.Unix(int64(timestamp), -1)
		volume, _ := strconv.ParseFloat(m.Volume, 64)
		price, _ := strconv.ParseFloat(m.Price, 64)
		fmt.Printf("[%02d:%02d:%02d]@%-15s TRADE: %-15f FOR %-15f \n",
			date.Hour(), date.Minute(), date.Second(),
			m.Symbol, volume, price)

	}

}
