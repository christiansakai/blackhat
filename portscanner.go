package main

import (
  "sort"
  "fmt"
  "net"
)

const workerCount = 10000
const maxPortNum = 6335

func worker(ports <-chan int, results chan<- int) {
  for p := range ports {
    address := fmt.Sprintf("127.0.0.1:%d", p)
    fmt.Println(address)

    conn, err := net.Dial("tcp", address)
    if err != nil {
      results <- 0
      continue
    }

    conn.Close()
    results <- p
  }
}

func main() {
  ports := make(chan int, workerCount)
  results := make(chan int)

  var openports []int

  // Create worker
  for i := 0; i < cap(ports); i++ {
    go worker(ports, results)
  }

  // Goroutine to start sending work to worker
  go func() {
    for i := 1; i <= maxPortNum; i++ {
      ports <- i
    }

    close(ports)
  }()

  // Main goroutine receive results synchronously
  for i := 1; i <= maxPortNum; i++ {
    port := <- results
    if port != 0 {
      openports = append(openports, port)
    }
  }

  close(results)

  // Sort and print
  sort.Ints(openports)
  for _, port := range openports {
    fmt.Printf("%d open\n", port)
  }
}
