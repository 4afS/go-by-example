package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(1, 2, 3, 4))
  s := []int{1, 2, 3, 4}
	fmt.Println(sum(s...))
}

func sum(nums ...int) int {
	n := 0
	for _, e := range nums {
		n += e
	}
	return n
}
