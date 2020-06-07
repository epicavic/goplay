package main

import "fmt"

func main() {
	const BoilingF = 212.0
	var f = BoilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
}
