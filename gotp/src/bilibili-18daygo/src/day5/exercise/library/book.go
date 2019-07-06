package library

import (
	"fmt"
)

type Book struct {
	Name string
	Copy int
	Author string
	Publication string
	WhoBorrow []*Student
}

type Library []*Book

func (l Library) Quary(name string) (res []*Book) {
	for _, book := range l {
		if book.Name == name {
			res = append(res, book)
		}
	}
	return res
}

func (l Library) String() string {
	var res string
	for _, v := range l {
		res = fmt.Sprintf("%s\n%v", res, *v)
	}
	return res
}
func (b Book) BorrowBy() []*Student {
	return b.WhoBorrow
}

