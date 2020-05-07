package main

import (
  "fmt"
  "log"
  "net"
  "io"
)

const port = 81

func echo(conn net.Conn) {
  defer conn.Close()

  if _, err := io.Copy(conn, conn); err != nil {
    log.Fatalln("Unable to read/write data")
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

    go echo(conn)
  }
}
