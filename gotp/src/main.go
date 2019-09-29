package main

import (
	"fmt"
	"time"
)

type Dog struct {
	name string
}

func (d *Dog) Name() string {
	return  d.name
}

func (d *Dog) SetName(s string) {
	d.name = s
}

type Pet interface {
	SetName(name string)
	Name() string
}


func main() {
	// dog := Dog{"little pig"}
	// var pet Pet = &dog
	// dog.SetName("monster")
	// fmt.Println(pet.Name())
	// fmt.Printf("%x\n", uint8(255))
	fmt.Println(1)
	ch := make(chan int, 1)
	go foo(ch)
	fmt.Println(2)
	for i := range ch {
		fmt.Println(i)
	}
	close(ch)
}

func foo(ch chan int)  {
	time.Sleep(5*time.Second)
	ch <- 1
}