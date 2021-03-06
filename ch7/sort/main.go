// Derivative of "The Go Programming Language" © 2016 examples by
// Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

// Track type
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

// convert time string into nanoseconds
func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // flush all buffered data
	fmt.Println()
}

// sort by Artist. add sorting methods Len, Less, Swap (sort.Interface parameter for sort.Sort function)
type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// sort by Duration
type byDuration []*Track

func (x byDuration) Len() int           { return len(x) }
func (x byDuration) Less(i, j int) bool { return x[i].Length < x[j].Length }
func (x byDuration) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// customized sorting. accepts custom less function logic
type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

// customized comparison function with multiple conditions
func customLess(x, y *Track) bool {
	if x.Title != y.Title {
		return x.Title < y.Title
	}
	if x.Year != y.Year {
		return x.Year < y.Year
	}
	if x.Length != y.Length {
		return x.Length < y.Length
	}
	return false
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func main() {
	fmt.Println("not sorted:")
	printTracks(tracks)

	fmt.Println("sort byArtist:")
	sort.Sort(byArtist(tracks))
	printTracks(tracks)

	fmt.Println("sort byDuration:")
	sort.Sort(byDuration(tracks))
	printTracks(tracks)

	fmt.Println("sort Custom:")
	sort.Sort(customSort{tracks, customLess})
	printTracks(tracks)
}
