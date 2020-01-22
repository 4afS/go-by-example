package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var (
		state    = make(map[int]int)
		mutex    = &sync.Mutex{}
		readOps  uint64
		writeOps uint64
	)

	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				mutex.Lock()
				total += state[rand.Intn(5)]
				mutex.Unlock()

				atomic.AddUint64(&readOps, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				mutex.Lock()
				state[rand.Intn(10)] = rand.Intn(100)
				mutex.Unlock()

				atomic.AddUint64(&writeOps, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}
	time.Sleep(10 * time.Second)

	fmt.Println("readOps: ", atomic.LoadUint64(&readOps))
	fmt.Println("writeOps: ", atomic.LoadUint64(&writeOps))

	mutex.Lock()
	fmt.Println("state: ", state)
	mutex.Unlock()
}
