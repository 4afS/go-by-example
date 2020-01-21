package main

import "fmt"

func main() {
  fmt.Println("switch")

  whatType := func(someType interface{}) {
    switch typeName := someType.(type) {
    case bool:
      fmt.Println("bool")
    case string:
      fmt.Println("string")
    default:
      fmt.Printf("dont know %T\n", typeName)
    }
  }

  whatType("hello")
}
