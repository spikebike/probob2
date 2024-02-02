package main

import (
	"io/ioutil"
	"log"
	"net"
	"github.com/golang/protobuf/proto"
	"sum"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	
	data, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Println(err)
		return
	}
	
	msg := &sum.ByteArrays{}
	err = proto.Unmarshal(data, msg)
	if err != nil {
		log.Println(err)
		return
	}
	
	// Write back data
	_, err = conn.Write(data)
	if err != nil {
		log.Println(err)
		return
	}
}
