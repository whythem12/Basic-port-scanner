package main

import (
	"fmt"
	"net"
	"time"
)

func (qwe *Scanner) Scan() bool {
	address := qwe.addr
	protocol := qwe.protocol
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

type Scanner struct {
	protocol string
	addr     string
}

func main() {
	println("Scanning is started.")
	asdf := Scanner{
		protocol: "tcp",
		addr: "google.com:3000",
	}
	open := asdf.Scan()
	fmt.Printf("Port Open: %t\n", open)
}
