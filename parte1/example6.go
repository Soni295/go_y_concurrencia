package main

import (
	"fmt"
	"sync"
	"time"
)

var comienza = time.Now()

func sumaUno(num int) {
	time.Sleep(time.Millisecond * 500)
	if num%2 == 0 {
		time.Sleep(time.Millisecond * 500)
	}
	tardo := time.Since(comienza)
	fmt.Printf("sumaUno: %v + 1 = %v\t[%v]\n", num, num+1, tardo)
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			sumaUno(j)
		}(i)
	}

	fmt.Printf("func main esperando las gorutinas: [%v]\n", time.Since(comienza))
	wg.Wait()
	fmt.Printf("func main termina: [%v]\n", time.Since(comienza))
}