package main

import (
	"fmt"
)

// channels pipeline: gen -> sqrt -> print
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// generator goroutine
	go func() {
		for x := 0; x <= 100; x++ {
			naturals <- x
			// time.Sleep(time.Second)
		}
		close(naturals)
	}()

	// squarer goroutine
	go func() {
		for {
			x, ok := <-naturals
			if !ok { // naturals channel was closed and drained and will send 0 values
				break
			}
			squares <- x * x
		}
		close(squares)
	}()

	// printer in main gouroutine
	for x := range squares { // another way to detect closed channel with for range loop
		fmt.Println(x)
	}
}
