package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", "127.0.0.1:6379")
	if err != nil { 
		fmt.Println("Failed to connect to port 6379")
		os.Exit(1)	
	}

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Failed to accept connection: ", err.Error())
		os.Exit(1)
	}

	conn.Write([]byte("+PONG\r\n"))
}

