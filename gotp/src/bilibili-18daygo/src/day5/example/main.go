package main

import (
	"math/rand"
	"fmt"
)

type song struct {
	title string
	time int
}

func (s1 song) show() {
	s1.title = "song1"
	s1.time = 327
	fmt.Printf("%p\n", &s1.title)
	fmt.Printf("%p\n", &s1.time)
}

type title struct {
	name string
}
type time struct {
	name int
}
type songLianxu struct {
	tit *title
	tim *time
}

func (s1 songLianxu) show() {
	s1.tit = &title{"song2"}
	s1.tim = &time{327}
	fmt.Printf("%p\n", &s1.tit)
	fmt.Printf("%p\n", &s1.tim)
}

type songLinked struct {
	title string
	time int
	next *songLinked
}

func linkedList() {
	head  := &songLinked{}
	head.title = "song"
	head.time = 10

	for i := 0; i < 10; i++ {
		s := &songLinked{
			title: fmt.Sprintf("song%d", i),
			time: rand.Intn(100),
		}
		
		head.next = s
		s.next = head
		
		fmt.Println(s)
	}
}

func main() {
	// s1:= new(song)
	// s1.show()

	// s2 := new(songLianxu)
	// s2.show()
	linkedList()

}