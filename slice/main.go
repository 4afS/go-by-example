package main

import "fmt"

func main() {
  s := make([]int, 10)
  fmt.Println(s)
  s = append(s, 1,4,6)
  fmt.Println(s)

  ss := make([]int, 10)
  fmt.Println(s)
  ss= append(ss, 1,4,6)
  fmt.Println(ss)

  for _, e := range ss {
    s = append(s, e)
  }

  fmt.Println(s)
}
