package main

import (
	"fmt"
)

type rectangle struct {
	width, height int
}

// func (r *rectangle) area() int {
func (r *rectangle) area() int {
  r.height = 1
  return r.width * r.height
}

func main() {
  r := rectangle{10, 20}
  fmt.Println(r.area())
  fmt.Println(r)
}
