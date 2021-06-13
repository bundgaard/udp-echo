package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	fmt.Println("ECHO UDP CLIENT")

	raddr, err := net.ResolveUDPAddr("udp", ":8001")
	if err != nil {
		log.Fatalf("resolve %v", err)
	}

	socket, err := net.DialUDP("udp",nil, raddr)
	if err != nil {
		log.Fatalf("dial: %v", err)
	}


	fmt.Fprintf(socket, "Hello, World! " + time.Now().Format(time.RFC3339))


}
