package main

import "fmt"

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
	dog := Dog{"little pig"}
	var pet Pet = &dog
	dog.SetName("monster")
	fmt.Println(pet.Name())

}