package main

import (
	"log"
	"net"
	"github.com/golang/protobuf/proto"
	"github.com/spikebike/probob2/sum"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:50051")
	if err != nil {
		log.Fatal(err)
	} else { 
		log.Println("dial worked")
   }
	defer conn.Close()

	// Create ByteArrays struct and marshal it to bytes
	msg := &sum.ByteArrays{
		Arrays: [][]byte{{0, 1, 2}, {3, 4, 5}},
	}

	data, err := proto.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	} else { 
		log.Println("dial worked")
   }

	_, err = conn.Write(data)
	if err != nil {
		log.Fatal(err)
	} else { 
		log.Println("write worked")
	}

	// Then, read echoed data
	echoData := make([]byte, 1024)
	n, err := conn.Read(echoData)
	if err != nil {
		log.Fatal(err)
	} else { 
		log.Println("read worked")
	}
	
	// Unmarshal echoed data
	echoMsg := &sum.ByteArrays{}
	err = proto.Unmarshal(echoData[:n], echoMsg)
	if err != nil {
		log.Fatal(err)
	} else { 
		log.Println("unmarshal worked")
	}

	log.Println("Read from server: ", echoMsg)
}
