
package main

import (
    "log"
    "net"

	"github.com/spikebike/probob2/matrix" // Import the generated Protobuf code
	"google.golang.org/protobuf/proto"

)

func main() {
    conn, err := net.Dial("tcp", "localhost:13080")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    // Create your data ([][]byte)
    data := [][]byte{
        []byte("row1-data"),
        []byte("row2-data"),
    }

    // Construct the protobuf message
    matrix := &matrix.Matrix{Rows: data}

    // Encode the protobuf message
    encodedData, err := proto.Marshal(matrix)
    if err != nil {
        log.Fatal("Marshal error:", err)
    }

    // Send encoded data to server
    _, err = conn.Write(encodedData)
    if err != nil {
        log.Fatal("Write error:", err)
    }

    // ... (Add code to receive response from the server if needed)
}

