package main

import (
  "fmt"
)

func main() {
  fmt.Print(plus(1, 3))
}

func plus(n int, m int) int {
  return n + m
}
