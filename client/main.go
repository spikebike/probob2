package main

import (
	"log"
	"net"
	"github.com/golang/protobuf/proto"
	"sum"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:50051")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Create ByteArrays struct and marshal it to bytes
	msg := &sum.ByteArrays{
		Arrays: [][]byte{{0, 1, 2}, {3, 4, 5}},
	}

	data, err := proto.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Write(data)
	if err != nil {
		log.Fatal(err)
	}

	// Then, read echoed data
	echoData := make([]byte, 1024)
	n, err := conn.Read(echoData)
	if err != nil {
		log.Fatal(err)
	}
	
	// Unmarshal echoed data
	echoMsg := &sum.ByteArrays{}
	err = proto.Unmarshal(echoData[:n], echoMsg)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Read from server: ", echoMsg)
}
