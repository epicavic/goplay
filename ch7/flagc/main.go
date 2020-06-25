package main

import (
	"flag"
	"fmt"
)

type celsius float64
type fahrenheit float64
type celsiusFlag struct{ celsius }

// "fahrenheit" to "celsius" conversion
func fToC(f fahrenheit) celsius { return celsius((f - 32.0) * 5.0 / 9.0) }

// methods to satisfy flag.Value interface {String() string, Set(string) error}
func (c celsius) String() string { return fmt.Sprintf("%g°C", c) } // "String" method is embedded from "celsius" type.
func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) //parse string into vars
	switch unit {
	case "C", "°C":
		f.celsius = celsius(value) // f is a receiver (object)
		return nil
	case "F", "°F":
		f.celsius = fToC(fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// define flag function
func cf(name string, value celsius, usage string) *celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.celsius
}

func main() {
	var temp = cf("temp", 20.0, "the temperature")
	flag.Parse()
	fmt.Println(*temp)
}
