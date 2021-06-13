package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	fmt.Println("UDP Server")

	// READ is blocking

	socket, err := net.ListenPacket("udp", ":8001")
	if err != nil {
		log.Fatalf("socket error: %v", err)
	}
	defer socket.Close()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			buf := make([]byte, 512)
			read, raddr, err := socket.ReadFrom(buf)
			if err != nil {
				log.Fatalf("read: %v", err)
			}

			fmt.Println("received ", read, "bytes from", raddr)
			fmt.Printf("CLIENT %v", buf[:read])
			socket.WriteTo(buf[:read], raddr)
			fmt.Println("sending back to remote")
			time.Sleep(500 * time.Millisecond)
		}
	}()
	wg.Wait()
}
