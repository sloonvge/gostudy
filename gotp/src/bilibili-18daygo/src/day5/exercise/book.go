package main

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
func (b Book) BorrowBy() []*Student {
	return b.WhoBorrow
}

