package main

import (
	"math/rand"
	"fmt"
)

type linkedNode struct {
	title string
	time int
	next *linkedNode
}

func linkedList() {
	head  := &linkedNode{}
	head.title = "song"
	head.time = 10

	var tmp *linkedNode
	for i := 0; i < 10; i++ {
		s := &linkedNode{
			title: fmt.Sprintf("song%d", i),
			time: rand.Intn(100),
		}

		tmp = head
		s.next = head
		head.next = tmp
		fmt.Println(s)
	}
}

func main() {
	
}