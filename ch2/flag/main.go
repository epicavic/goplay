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
