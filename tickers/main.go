package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case tt := <-t.C:
				fmt.Println("tick at", tt)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	t.Stop()
	done <- true
	fmt.Println("ticker stopped")
}
