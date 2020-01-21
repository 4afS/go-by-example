package main

import "fmt"

func main() {
  fmt.Println("array")
  var a [5]int
  fmt.Println(a, len(a), cap(a))

  for i := 0; i < len(a); i++ {
    fmt.Println(a[i])
  }
}
