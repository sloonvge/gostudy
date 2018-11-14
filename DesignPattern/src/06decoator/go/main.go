package main

import "./dec"

func main(){
	p := dec.NewPerson("Tom")
	f := dec.NewFinery()
	f.Decorate(p)
	t := dec.NewTShirts()
	t.Decorate(f)
	t.Show()
	//p.WearTShirts()
	//p.WearPaints()
	//p.WearShoes()
	//p.Show()
}
