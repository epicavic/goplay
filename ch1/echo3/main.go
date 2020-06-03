package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	defer timeMeasure(time.Now(), "main")
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[:])
	for i, arg := range os.Args[:] {
		fmt.Printf("%d %s\n", i, arg)
	}
}

func timeMeasure(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
