package dec

import "fmt"

type TShirts struct {
	Clothes IClothes
}

func NewTShirts(c IClothes) (t *TShirts) {
	t = &TShirts{Clothes: c}
	return
}

func (t *TShirts) Decorate(c IClothes) {
	t.Clothes = c
}

func (t *TShirts) Show() {
	fmt.Println("T shirts")
	t.Clothes.Show()
}

type BigTrouser struct {
	Clothes IClothes
}

func NewBigTruser(c IClothes) (b *BigTrouser) {
	b = &BigTrouser{Clothes: c}
	return
}

func (b *BigTrouser) Decorate(c IClothes) {
	b.Clothes = c
}

func (b *BigTrouser) Show() {
	fmt.Println("Big triuser!")
	b.Clothes.Show()
}
