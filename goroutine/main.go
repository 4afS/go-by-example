package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	go f(s)
	go f(s)
	go f(s)
  fmt.Scanln()
  fmt.Println("done")
}

func f(s []int) {
	for _, e := range s {
		fmt.Println(e)
	}
}
