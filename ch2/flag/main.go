// Derivative of "The Go Programming Language" © 2016 examples by
// Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "no newline")
var s = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *s))
	if !*n {
		fmt.Println()
	}
}
