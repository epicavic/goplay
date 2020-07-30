// Derivative of "The Go Programming Language" © 2016 examples by
// Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4} //slice
	b := 1                 //int
	c := &b                //pointer
	d := map[string]int{   //map
		"alice": 1,
		"bob":   2,
	}

	// Arguments are passed by value, so the function receives a copy
	// of each argument, modifications to the copy do not affect the caller.
	// However, if the argument contains some kind of reference,
	// like a pointer, slice, map, function, or channel, then the caller
	// may be affected by any modifications the function makes to variables
	// indirectly referred to by the argument.

	adda(a, b)
	fmt.Println(a) // returns "[2 2 3 4]" : reference type - changed
	addb(b, b)
	fmt.Println(b) // returns "1" : integer type - not changed
	modc(c)
	fmt.Println(c, *c, b) // returns "0xc00009e068 2 2" : reference type - changed
	addd(d)
	fmt.Println(d) // returns "map[alice:2 bob:3]" : reference type - changed
}

func adda(aaa []int, b int) []int {
	aaa[0] += b
	return aaa
}

func addb(bbb, b int) int {
	bbb += b
	return bbb
}

func modc(ccc *int) *int {
	*ccc = 2
	return ccc
}

func addd(ddd map[string]int) map[string]int {
	for k := range ddd {
		ddd[k]++
	}
	return ddd
}
