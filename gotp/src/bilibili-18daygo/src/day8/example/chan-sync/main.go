package main

import (
	// "time"
	"fmt"
)

func send(ch chan int, sign chan struct{}) {
	for i :=0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	sign <- struct{}{}
	fmt.Println("send exit")
}

func recv(ch chan int, sign chan struct{}) {
	for {
		i, ok := <- ch
		if !ok {
			break
		}
		fmt.Println(i)
	}
	sign <- struct{}{}
	fmt.Println("recv exit")
}

func main() {
	ch := make(chan int, 10)
	sign := make(chan struct{}, 2)
	go send(ch, sign)
	go recv(ch, sign)
	// var i = 0
	// for _ = range sign {
	// 	i++
	// 	if i == 2 {
	// 		break
	// 	}
	// }
	// close(sign)
	for i := 0; i < 2; i++ {
		<-sign
	}
}