package main

import (
	"fmt"
	"math/rand"
	"time"
)

/** 
* Created by wanjx in 2019/7/7 22:02
**/
 
func read(ch chan int) {
	a := <- ch
	fmt.Printf("%d\t", a)
	time.Sleep(time.Second)
}

func write(ch chan int) {
	ch <- rand.Int()
}

func main() {
	ch := make(chan int, 10)
	go write(ch)
	go read(ch)

	time.Sleep(time.Second * 5)
}