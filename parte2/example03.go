package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	countV2 = 0
	wg      = sync.WaitGroup{}
	limit   = 10
	reads   = 1_000
	mutex   = sync.RWMutex{} // < mutex
)

func readCount() {
	defer wg.Done()
	mutex.RLock() // bloqueo
	fmt.Println(countV2)
	mutex.RUnlock() // desbloqueo
}

func sumOneV2() {
	defer wg.Done()
	mutex.RLock() // bloqueo
	countV2++
	time.Sleep(time.Second * 1)
	mutex.RUnlock() // desbloqueo
}

func main() {

	wg.Add(limit + reads)
	for i := 0; limit > i; i++ {
		go sumOneV2()
	}

	for i := 0; reads > i; i++ {
		go readCount()
	}

	wg.Wait()

	fmt.Printf("countV2 [%v]\n", countV2)
}