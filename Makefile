.PHONY: all clean proto

all: server client

clean: 
	rm -f matrix/matrix.pb.go server client

proto:
	protoc --go_out=. --go_opt=paths=source_relative  matrix/matrix.proto

#	protoc --go_out=. matrix.proto

server: proto
	go build -o server server.go

client: proto
	go build -o client client.go
