
package main

import (
    "fmt"
//    "io"
    "log"
    "net"

    "github.com/spikebike/probob2/matrix"
	 "google.golang.org/protobuf/proto"

)

func handleConnection(conn net.Conn) {
    defer conn.Close()

    // Receive encoded protobuf data
    buf := make([]byte, 4096) // Adjust buffer size if needed
    n, err := conn.Read(buf)
    if err != nil {
        log.Println("Read error:", err)
        return
    }

    // Decode protobuf data
    var matrix matrix.Matrix
    err = proto.Unmarshal(buf[:n], &matrix)
    if err != nil {
        log.Println("Unmarshal error:", err)
        return
    }

    // Process received data ([][]byte)
    fmt.Println("Received matrix:", matrix.Rows)

    // ... (Add your server-side logic for processing and responding)
}

func main() {
    listener, err := net.Listen("tcp", ":13080")
    if err != nil {
        log.Fatal(err)
    }
    defer listener.Close()

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Println(err)
            continue
        }
        go handleConnection(conn)
    }
}

