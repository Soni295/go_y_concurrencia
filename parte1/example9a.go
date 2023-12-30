package main

import (
	"fmt"
	"sync"
	"time"
)

var comienza = time.Now()

func sumaUno(num int, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(time.Millisecond * 500)
	if num%2 == 0 {
		time.Sleep(time.Millisecond * 500)
	}

	tardo := time.Since(comienza)
	fmt.Printf("sumaUno: %v + 1 = %v\t[%v]\n", num, num+1, tardo)
}

func main() {
	wg := sync.WaitGroup{}
	gorutinas := 10
	wg.Add(gorutinas + 1) // agrego una de mas

	for i := 0; i < gorutinas; i++ {
		go sumaUno(i, &wg)
	}

	fmt.Printf("func main esperando las gorutinas: [%v]\n", time.Since(comienza))
	wg.Wait()
	fmt.Printf("func main termina: [%v]\n", time.Since(comienza))
}