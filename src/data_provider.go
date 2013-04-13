// You can edit this code!
// Click here and start typing.
package main

import (
	"btccharts"
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

	for true {
		line, _ := reader.ReadString('\n')
		m := btccharts.GetMessage(line)
		fmt.Printf("%s\n", m)

	}

	defer conn.Close()

}
