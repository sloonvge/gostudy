package main

type Student struct {
	Name string
	Class string
	IdCard int
	Sex string
	BorrowBook []*Book
}

func (s *Student) Borrow(b *Book) {
	s.BorrowBook = append(s.BorrowBook, b)
	b.WhoBorrow = append(b.WhoBorrow, s)
}




