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
	mux.Handle("/list", http.HandlerFunc(db.list))   // mux.HandleFunc("/list", db.list)
	mux.Handle("/price", http.HandlerFunc(db.price)) // mux.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}

// db.list is a function that implements handler-like behavior, but since it has no methods,
// it doesn’t satisfy the http.Handler interface and can’t be passed directly to mux.Handle.
// The expression http.HandlerFunc(db.list) is a conversion, not a function call, since
// http.HandlerFunc is a type

// type HandlerFunc func(w ResponseWriter, r *Request)
// func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) { f(w, r) }

// for convenience, net/http provides a global ServeMux instance called DefaultServeMux
// and package-level functions called http.Handle and http.HandleFunc.
// To use DefaultServeMux as the server’s main handler, we needn’t pass it to ListenAndServe; nil will do.
// http.HandleFunc("/list", db.list)
// http.HandleFunc("/price", db.price)
// log.Fatal(http.ListenAndServe("localhost:8000", nil))

// web server invokes each handler in a new goroutine, so handlers must take precaution s such as locking
// when accessing variables that other goroutines, including other requests to the same handler
