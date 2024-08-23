package main

import (
	"fmt"
	"sync"
)

var (
	countV1 = 0
	countV2 = 0
	wg      = sync.WaitGroup{}
	limit   = 1_000
)

func sumOneV1() { countV1++ }

func sumOneV2() {
	defer wg.Done()
	countV2++
}

func main() {
	for i := 0; limit > i; i++ {
		sumOneV1()
	}

	wg.Add(limit)
	for i := 0; limit > i; i++ {
		go sumOneV2()
	}
	wg.Wait()

	fmt.Printf("countV1 [%v]\n", countV1)
	fmt.Printf("countV2 [%v]\n", countV2)
}