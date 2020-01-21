package main

import (
	// "errors"
	"fmt"
)

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f(arg int) (int, error) {
	if arg == 0 {
		return -1, &argError{arg, "inavlid"}
	}
	return arg, nil
}

func main() {
	if arg, e := f(0); e != nil {
		fmt.Println(arg)
	}

	_, e := f(0)
	if ea, ok := e.(*argError); ok {
		fmt.Println(ea.arg)
		fmt.Println(ea.prob)
	}
}
