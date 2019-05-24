package main

import (
	"testing"
	"math/rand"
	"fmt"
	"time"
)

func TestSelect(t *testing.T) {
	intChans := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1)}

	index := rand.Intn(3)
	fmt.Println("The index: ", index)
	intChans[index] <- index

	select {
	case <-intChans[0]:
		fmt.Println("Select First")
	case <-intChans[1]:
		fmt.Println("Select Second")
	case <-intChans[2]:
		fmt.Println("Select Third")
	default:
		fmt.Println("Select None")
	}

	for i := 0; i < 5; i++ {
		fmt.Println("T", rand.Intn(3))
	}

}

func TestChannelClosed(t *testing.T) {
	intChan := make(chan int, 1)
	time.AfterFunc(1 * time.Second, func() {
		close(intChan)
	})
	// intChan <- 1
	select {
	case _, ok := <-intChan:
		if !ok {
			fmt.Println("Closed Channel!")
			break
		}
		fmt.Println("Select this case")
	}
}

func TestSelectFor(t *testing.T) {
	num := 5
	ch := make(chan int, num)
	for i := 0; i < num; i++ {
		ch <- i
	}
	close(ch)

	for i := 0; i < num; i++ {
		select {
		case c := <- ch:
			if c == 0{
				break
			}
			fmt.Printf("%v\n", c)
		}
		//fmt.Printf("%v\n", v)
	}
}

func TestSelectClosed(t *testing.T) {
	intChan1 := make(chan int, 1)
	intChan2 := make(chan int, 1)

	close(intChan1)
	intChan2 <- 1
	select {
	case _, ok := <- intChan1:
		if !ok {
			fmt.Printf("closed chan1\n")
			intChan1 = nil
			break
		}
		fmt.Printf("value1: \n")
	case _, ok := <- intChan2:
		if !ok {
			fmt.Printf("closed chan2\n")
		}
		fmt.Printf("value2: \n")

	}
}

// 通道用法1. 打印log
func TestChannelUse(t *testing.T) {
	tic := time.NewTicker(2 * time.Second)
	defer tic.Stop()

	for v := range tic.C {
		fmt.Printf("[INFO] %v s\n", v)
	}
}

func TestChannelUse2(t *testing.T) {
	// ratePerSec := 10
	var dur time.Duration = 1e9
	chRate := time.Tick(dur)
	for i := 0; i < 5; i++ {
		<- chRate
		go func() {
			fmt.Printf("RPC requests:\n")
		}()
	}
}



// https://www.kancloud.cn/kancloud/the-way-to-go/165094
// 简单超时模式
