package library

type Student struct {
	Name string `json:"name`
	Class string `json:"class"`
	IdCard int	`json:"id_card"`
	Sex string	`json:"sex"`
	BorrowBook []*Book `json:"borrow_book"`
}

func NewStudent() *Student {
	return &Student{
		IdCard: -1,
		BorrowBook: make([]*Book, 0),
	}
}

func (s *Student) Borrow(b *Book) {
	s.BorrowBook = append(s.BorrowBook, b)
	b.WhoBorrow = append(b.WhoBorrow, s)
}

func (s *Student) Borrows(b []*Book) {
	for _, v := range b {
		s.Borrow(v)
	}
}




