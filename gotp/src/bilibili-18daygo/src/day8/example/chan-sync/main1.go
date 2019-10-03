package main

import (
	"fmt"
	"time"
)

func calc(intChan chan int, resultChan chan int) {
	for v := range intChan {
		var flag = true
		if v % 2 == 0 {
			flag = false
		}
		if flag {
			resultChan <- v
		}
	}
}

func main() {
	intChan := make(chan int, 1000)
	resultChan := make(chan int, 1000)
	for i := 0; i < 100; i++ {
		intChan <- i
	}

	for i := 0; i< 8; i++ {
		go calc(intChan, resultChan)
	}

	go func() {
		for v := range resultChan {
			fmt.Println(v)
		}
	}()

	time.Sleep(time.Second)
}
