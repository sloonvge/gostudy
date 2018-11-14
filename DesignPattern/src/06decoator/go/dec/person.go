package dec

import "fmt"

type IShow interface {
	Show()
}

type Person struct {
	Name string
}

func NewBlankPerson() (p *Person){
	p = &Person{}
	return
}

func NewPerson(name string) (p *Person) {
	p = &Person{Name: name}
	return
}

func (p *Person) Show() {
	fmt.Printf("%s", p.Name)
}
