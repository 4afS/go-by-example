package main

import (
  "fmt"
)

func main() {
  nextInt := intSeq(100)

  fmt.Println(nextInt(1))
  fmt.Println(nextInt(2))
  fmt.Println(nextInt(3))
}

func intSeq(i int) func(int) int {
  return func(n int) int {
    i += n
    return i
  }
}
