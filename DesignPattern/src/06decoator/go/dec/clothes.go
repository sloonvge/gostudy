package dec

import "fmt"

type Finery struct {
	Person
	Component *Person
}

func NewFinery() (f *Finery){
	f = &Finery{}
	return
}

func (f *Finery) Decorate(p *Person) {
	f.Component = p
}

func (f *Finery) Show() {
	if f.Component != nil {
		f.Component.Show()
	}
}


type TShirts struct {
	Finery
}

func NewTShirts() (t *TShirts) {
	t = &TShirts{}
	return
}

func (t *TShirts) Decorate(f *Finery) {
	t.Component = f.Component
}

func (t *TShirts) Show() {
	fmt.Println("T shirts")
	t.Component.Show()
}

