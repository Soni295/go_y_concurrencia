package main

import (
	"fmt"
	"time"
)

var comienza = time.Now()

func sumaUno(num int) {
	time.Sleep(time.Millisecond * 500)
	tardo := time.Since(comienza)
	fmt.Printf("sumaUno: %v + 1 = %v\t[%v]\n", num, num+1, tardo)
}

func main() {
	for i := 0; i < 10; i++ {
		go func() { sumaUno(i) }()
	}
	fmt.Printf("func main termina: [%v]\n", time.Since(comienza))
}