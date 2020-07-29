package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func main() {
	var verbose = flag.Bool("v", false, "print progress") // -v flag enables progress print
	var tick <-chan time.Time                             // progress print delay
	var nfiles, nbytes int64                              // total files and bytes counters
	var n sync.WaitGroup                                  // count number of calls to walkDir that are still active
	fileSizes := make(chan int64)                         // filesizes unbuffered channel (pass size from anon goroutine to main goroutine)

	flag.Parse()         // parse commandline flags and arguments
	roots := flag.Args() // get list of dirs
	if len(roots) == 0 {
		roots = []string{"."} // or use current
	}

	for _, root := range roots { // traverse dirs recursively
		n.Add(1)                        // initial increment for number of calls to walkDir
		go walkDir(root, &n, fileSizes) // call walkDir concurrently
	}

	go func() {
		n.Wait()         // wait when counter drops to zero
		close(fileSizes) // close and drain channel at the end of traversal
	}()

	if *verbose { // if verbose flag set to true set tick channel to 500 ms (print delay)
		tick = time.Tick(500 * time.Millisecond)
	}

loop: // label. optional for 'break' and 'continue', mandatory for 'goto' statements
	for {
		select { // select action depending which channel receives event
		case size, ok := <-fileSizes:
			if !ok {
				break loop // break for loop when fileSizes was closed
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes) // printed only when -v flag is used. otherwise tick channel is nil and not selected
		}
	}
	printDiskUsage(nfiles, nbytes) // print final totals
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.2f MB\n", nfiles, float64(nbytes)/1e6)
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done() // decrement counter in the end in any case
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1) // increment counter
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes) // call walkDir concurrently
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	var sema = make(chan struct{}, 20) // buffered channel of empty structs used as counting semaphore for limiting concurrency
	sema <- struct{}{}                 // acquire token. will block execution when channel filled(synchronizer)
	defer func() { <-sema }()          // release token. will continue execution when there are empty slots available

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du2: %v\n", err)
		return nil
	}
	return entries
}

/*
$time go run du2/main.go c:/Users/ 2>/dev/null
64251 files  78206.86 MB

real    0m1.968s
*/
