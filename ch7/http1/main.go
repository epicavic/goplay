// Derivative of "The Go Programming Language" © 2016 examples by
// Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import (
	"fmt"
	"log"
	"net/http"
)

// define types
type dollars float64
type database map[string]dollars

// define methods to satisfy fmt.Stringer and http.Handler interfaces
func (d dollars) String() string { return fmt.Sprintf("$ %2.2f", d) }
func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func main() {
	db := database{"pants": 20, "shirts": 10}
	log.Fatal(http.ListenAndServe("localhost:8080", db))
}
