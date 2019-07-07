package model

import (
	"fmt"
	"errors"
)

var (
	ErrBookNotFound = errors.New("book Not found")
)

type Student struct {
	Name string `json:"name`
	Class string `json:"class"`
	IdCard int	`json:"id_card"`
	Sex string	`json:"sex"`
	BorrowBook []*BorrowItem `json:"borrow_book"`
	BorrowHistory []*BorrowHistoryItem `json:"borrow_history"`
}

type BorrowItem struct {
	book *Book
	num int
}

type BorrowHistoryItem struct {
	book *Book
	time string
}

func CreateStudent(name, class string, idCard int, sex string) *Student{
	stu:= &Student{
		Name: name,
		Class: class,
		IdCard: idCard,
		Sex: sex,
	}
	return stu
}

func (s *Student) AddBook(b *BorrowItem) {
	s.BorrowBook = append(s.BorrowBook, b)
	s.BorrowHistory = append(s.BorrowHistory, &BorrowHistoryItem{
		book: b.book,
		time: "today",
	})
}

func (s *Student) DeleteBook(b *BorrowItem) (err error) {
	for i := 0; i < len(s.BorrowBook); i++ {
		if s.BorrowBook[i].book.Name == b.book.Name {
			if b.num == s.BorrowBook[i].num {
				s.BorrowBook = append(s.BorrowBook[0:i], s.BorrowBook[i+1:]...)
				return
			}
			s.BorrowBook[i].num -= b.num
			return
		}
	}
	err = ErrBookNotFound
	return
}

func (s *Student) History() (res string) {
	for _, item := range s.BorrowHistory {
		res = fmt.Sprintf("%s\name: %s time: %s", res, item.book, item.time)
	}
	return
}







// func NewStudent() *Student {
// 	return &Student{
// 		IdCard: -1,
// 		BorrowBook: make([]*Book, 0),
// 	}
// }

// func (s *Student) Borrow(b *Book) {
// 	s.BorrowBook = append(s.BorrowBook, b)
// 	b.WhoBorrow = append(b.WhoBorrow, s)
// }

// func (s *Student) Borrows(b []*Book) {
// 	for _, v := range b {
// 		s.Borrow(v)
// 	}
// }




