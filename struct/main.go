package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
  return &person{name: name, age: age}
}

func main() {
  fmt.Println(newPerson("me", 19))
}
