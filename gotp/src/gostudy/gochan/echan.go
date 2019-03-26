package main

import "fmt"

func getIntChan() <-chan int {
	num := 5
	ch := make(chan int, num)
	for i := 0; i < num; i++ {
		ch <- i
	}
	close(ch)
	return ch
}

func main() {
	intChan := getIntChan()
	//fmt.Printf("%v\n", intChan)
	fmt.Printf("%T\n", intChan)
	for v := range intChan {
		fmt.Printf("%v\n", v)
	}
}

