package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var (
	listen = flag.String("listen", "127.0.0.1:6397", "address to listen to")
)

func main() {
	flag.Parse()

	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() (err error) {
	l, err := net.Listen("tcp", *listen)
	if err != nil {
		return fmt.Errorf("listen: %w", err)
	}

	defer closeIt(l, &err, "close listener")

	c, err := l.Accept()
	if err != nil {
		return fmt.Errorf("accept: %w", err)
	}

	defer closeIt(c, &err, "close connection")

	buf := make([]byte, 128)

	_, err = c.Read(buf)
	if err != nil {
		return fmt.Errorf("read command: %w", err)
	}

	log.Printf("read command:\n%s", buf)

	_, err = c.Write([]byte("+PONG\r\n"))
	if err != nil {
		return fmt.Errorf("write response: %w", err)
	}

	return nil
}

func closeIt(c io.Closer, errp *error, msg string) {
	err := c.Close()
	if *errp == nil {
		*errp = fmt.Errorf("%v: %w", msg, err)
	}
}