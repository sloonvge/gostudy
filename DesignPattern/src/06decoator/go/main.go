package main

import "./dec"

func main(){
	p := dec.NewPerson("Tom")

	t := dec.NewTShirts(p)
	b := dec.NewBigTruser(t)

	b.Show()
}
