package dec

import "fmt"

type IClothes interface {
	Show()
}

type Person struct {
	Name string
}

func NewPerson(name string) (p *Person) {
	p = &Person{Name: name}
	return
}

func (p *Person) Show() {
	fmt.Printf("%s", p.Name)
}
