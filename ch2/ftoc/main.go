package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, ftoc(freezingF)) // "32°F = 0°C"
	fmt.Printf("%g°F = %g°C\n", boilingF, ftoc(boilingF))   // "212°F = 100°C"
}

func ftoc(f float64) float64 {
	return (f - 32) * 5 / 9
}
