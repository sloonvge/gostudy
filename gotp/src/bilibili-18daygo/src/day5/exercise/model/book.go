package model

import (
	"errors"
	"fmt"
)

var (
	ErrStockNotEnough = errors.New("stock is not enough")
)

type Book struct {
	Name string
	Total int
	Author string
	Publication string
	WhoBorrow []*Student
}

func CreateBook(name string, total int, author string, publication string) (b *Book) {
	b = &Book{
		Name: name,
		Total: total,
		Author: author,
		Publication: publication,
	}
	return
}

func (b *Book) canBorrow(c int) bool {
	return b.Total >= c
}

func (b *Book) Borrow(c int) (err error){
	if b.canBorrow(c) {
		err = ErrStockNotEnough
		return
	}
	b.Total -= c
	return
}

func (b *Book) Back(c int) (err error) {
	b.Total += c
	return
}

type Library []*Book

func (l Library) Quary(key string, value interface{}) (res []*Book){
	switch key {
	case "name":
		return l.QuaryName(value.(string))
	case "author":
		return l.QuaryAuthor(value.(string))
	case "publication":
		return l.QuaryPublication(value.(string))
	default:
		fmt.Println("useless search options.")
	}
	return
}

func (l Library) QuaryName(name string) (res []*Book) {
	for _, book := range l {
		if book.Name == name {
			res = append(res, book)
		}
	}
	return res
}

func (l Library) QuaryAuthor(author string) (res []*Book) {
	for _, book := range l {
		if book.Author == author {
			res = append(res, book)
		}
	}
	return res
}

func (l Library) QuaryPublication(publication string) (res []*Book) {
	for _, book := range l {
		if book.Publication == publication {
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

