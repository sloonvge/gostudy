package main

import (
	"sort"
	"testing"
)

/** 
* Created by wanjx in 2019/5/13 23:31
**/

var track = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func TestByArtist(t *testing.T) {
	sort.Sort(byArtist(track))

	printTracks(track)
}