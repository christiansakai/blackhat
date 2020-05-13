package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

const dstPort = 81
const srcPort = 80

func handle(src net.Conn) {
	dst, err := net.Dial("tcp", fmt.Sprintf(":%d", dstPort))
	if err != nil {
		log.Fatalln("Unable to connect to our unreachable host")
	}

	defer dst.Close()

	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()

	// Copy dst's output back to src
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", srcPort))
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}

	log.Println(fmt.Sprintf("Listening on 0.0.0.0:%d", srcPort))

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}

		go handle(conn)
	}
}
