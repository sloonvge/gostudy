package main

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

/** 
* Created by wanjx in 2019/5/13 23:20
**/
 
type Track struct {
	Title string
	Artist string
	Album string
	Year int
	Length time.Duration
}


func length(s string) time.Duration{
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track)  {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8,
		2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

type byArtist []*Track

func (t byArtist) Len() int { return len(t) }
func (t byArtist) Less(i, j int) bool { return t[i].Artist < t[j].Artist }
func (t byArtist) Swap(i, j int) { t[i], t[j] = t[j], t[i]}