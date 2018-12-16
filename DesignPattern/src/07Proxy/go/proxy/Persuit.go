package proxy

import "fmt"

type Persuit struct {
	MM *Girl
}

func NewPersuit(mm *Girl) (p *Persuit) {
	p = &Persuit{MM: mm}
	return
}

func (p *Persuit) GiveDolls() {
	fmt.Println("give dolls to " + p.MM.Name)
}

func (p *Persuit) GiveFlowers() {
	fmt.Println("give flowers to "+ p.MM.Name)
}

func (p *Persuit) GiveChocolate() {
	fmt.Println("give chocolate to "+ p.MM.Name)
}