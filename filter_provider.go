// You can edit this code!
// Click here and start typing.
package main

import (
    "fmt"
    "bufio"
    "net"
    "os"
    "strings"
    "io"
    "log"
    "encoding/json"
    )


func main() {
    conn, err := net.Dial("tcp", "bitcoincharts.com:27007")
    if err != nil {
        fmt.Println(err)
    os.Exit(1)
    }
    reader := bufio.NewReader(conn)

    type Message struct {
        Symbol string
        Volume, Id, Timestamp, Price float64
    }

    symbols := make(map[string]float64)
    i := 0.01
    for(true){
        line, _ := reader.ReadString('\n');
        dec := json.NewDecoder(strings.NewReader(line))
        for {
            var m Message
            if err := dec.Decode(&m); err == io.EOF {
                break
            } else if err != nil {
                log.Fatal(err)
            }
            if strings.Contains(m.Symbol, "USD")  {
                /*fmt.Printf("Zeile:  %+v", m)*/
                if _, ok := symbols[m.Symbol]; !ok {
                    i += 0.01
                    fmt.Printf("Adding %s as %f\n", m.Symbol, i)
                    symbols[m.Symbol] = i
                    fmt.Printf("Provider: %#v\n\n", symbols)
                }
            }

            /*fmt.Printf("%s: %f\n", m.Symbol, m.Price)*/
        }
    }
    
    defer conn.Close()                                                          
                 
}

