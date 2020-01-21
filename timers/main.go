package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.NewTimer(2 * time.Second)

	<-t1.C

	fmt.Println("t1")

	t2 := time.NewTimer(time.Second)

	go func() {
		<-t2.C
		fmt.Println("t2 fired")
	}()

	s2 := t2.Stop()
	if s2 {
		fmt.Println("t2 stopped")
	}

	time.Sleep(2 * time.Second)
}
