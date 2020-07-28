package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// traverse dirs synchronously and print the result at the end of execution
func main() {
	roots := os.Args[1:] // get directories as args
	if len(roots) == 0 {
		roots = []string{"."} // or process current
	}

	fileSizes := make(chan int64) // unbuffered channel to send entries size to

	go func() { // async goroutine to traverse directories
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes) // close and drain channel at the end of traversing
	}()

	var nfiles, nbytes int64
	for size := range fileSizes { // range over unbuffered channel until it's closed. go magic !!!
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes) // print results
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.2f MB\n", nfiles, float64(nbytes)/1e6)
}

func walkDir(dir string, fileSizes chan<- int64) { // traverse directories and send filesizes to channel
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size() // send each entry size in bytes to unbuffered channel
		}
	}
}

func dirents(dir string) []os.FileInfo { // return list of dir entries
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

/*
$time go run ch8/du/main.go c:/Users/vb
du: open c:\Users\vb\AppData\Local\ElevatedDiagnostics: Access is denied.
du: open c:\Users\vb\AppData\Local\Temp\CR_B4861.tmp: Access is denied.
du: open c:\Users\vb\AppData\Local\Temp\msdtadmin: Access is denied.
63729 files  78002.27 MB

real    0m2.655s
*/
