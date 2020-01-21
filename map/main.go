package main

import "fmt"

func main() {
  m := make(map[string]int)

  m["m1"] = 1
  m["m2"] = 2

  fmt.Println(m)

  a, prs := m["m3"]
  fmt.Println(a, prs)
}
