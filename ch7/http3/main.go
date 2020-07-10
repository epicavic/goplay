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

// define http handler methods
func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("q")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func main() {
	db := database{"pants": 20, "shirts": 10}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
