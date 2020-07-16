package main

import (
	"fmt"
	"time"
)

func main() {
	const n, delay = 45, 200
	go spinner(delay * time.Millisecond) // executed independently. terminated when main exits
	fibN := fib(n)                       // slow calculation
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
