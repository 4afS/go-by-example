package main

import (
  "fmt"
  "time"
)

func worker(done chan bool) {
  fmt.Print("working...")
  time.Sleep(time.Second)
  fmt.Print("done")

  done <- true
}

func main() {
  done := make(chan bool, 1)
  go worker(done)

  if <-done {
    print("done")
  }
}
