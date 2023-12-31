package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

// To sort a custom type, we define a new type
// with the methods required by sorting interface
// like Len, Less and Swap
// Instead of defining new type of []Track, we are
// using a struct that combines type Track with
// required functions
type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (c customSort) Len() int {
	return len(c.t)
}

func (c customSort) Less(i, j int) bool {
	return c.less(c.t[i], c.t[j])
}

func (c customSort) Swap(i, j int) {
	c.t[i], c.t[j] = c.t[j], c.t[i]
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTrack(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "-----", "-----", "-----", "-----")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

func main() {
	var tracks = []*Track{
		{"Go", "Delliah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "martin Solveig", "Smash", 2011, length("4m24s")},
	}
	// Before sorting
	printTrack(tracks)

	// After sort
	// Only comparison function (less) will change for
	// different types of Track properties. So, we only
	// need to provide custom less function when doing
	// sorting using customSort
	sort.Sort(customSort{tracks, func(x, y *Track) bool {
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
	}})
	printTrack(tracks)
}

// Output:
// Before sort:
// Title       Artist          Album              Year   Length
// -----       -----           -----              -----  -----
// Go          Delliah         From the Roots Up  2012   3m38s
// Go          Moby            Moby               1992   3m37s
// Go Ahead    Alicia Keys     As I Am            2007   4m36s
// Ready 2 Go  martin Solveig  Smash              2011   4m24s
//
// After sort:
// Title       Artist          Album              Year   Length
// -----       -----           -----              -----  -----
// Go          Moby            Moby               1992   3m37s
// Go          Delliah         From the Roots Up  2012   3m38s
// Go Ahead    Alicia Keys     As I Am            2007   4m36s
// Ready 2 Go  martin Solveig  Smash              2011   4m24s
