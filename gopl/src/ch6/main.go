package main

import "fmt"

func main() {
	m := make(map[int]int, 0)
	f := func(m map[int]int) {
		m[1] = 1
	}
	f(m)
	fmt.Printf("%v\n", m)

	c := m
	c[2] = 2
	c = nil
	fmt.Printf("%v-%v\n", m, c)
}

