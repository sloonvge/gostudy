package main

import (
	"math/rand"
)

func testMap() {
	var a map[int]int
	a = make(map[int]int, 0)
	a[8] = 10
	a[3] = 10
	a[2] = 10
	a[1] = 10
	a[18] = 10
	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			b[8] = rand.Intn(100)
		}(a)
	}
}

func main() {
	testMap()
}