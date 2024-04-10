package main

import (
	"fmt"
	"io"
	"net"
)

func process(client net.Conn) {
	target, err := net.Dial("tcp", "198.251.81.97:80")
	if err != nil {
		fmt.Printf("Connect failed: %v\n", err)
		return
	}
	forward := func(src, dest net.Conn) {
		defer src.Close()
		defer dest.Close()
		io.Copy(src, dest)
	}
	go forward(client, target)
	go forward(target, client)
}

func main() {
	server, err := net.Listen("tcp", ":2805")
	if err != nil {
		fmt.Printf("Listen failed: %v\n", err)
		return
	}

	for {
		client, err := server.Accept()
		if err != nil {
			fmt.Printf("Accept failed: %v\n", err)
			continue
		}
		go process(client)
	}
}
