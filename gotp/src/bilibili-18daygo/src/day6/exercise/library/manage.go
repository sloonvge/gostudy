package library

import (
	"fmt"
	"encoding/json"
)

func ManageStudent(s *Student, value interface{}) {
	var stu *Student
	bytes, err := json.Marshal(value)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(bytes, stu)
	if err != nil {
		fmt.Println(err)
	}
	if stu.Name != "" {
		s.Name = stu.Name
	}
	if stu.IdCard != -1 {
		s.IdCard = stu.IdCard
	}
	if stu.Class != "" {
		s.Class = stu.Class
	}
	if stu.Sex != "" {
		s.Sex = stu.Sex
	}

}


func (l Library) BorrowInfo() string {
	var res string
	for _, b := range l {
		res = fmt.Sprintf("%s\n%+v %v", res, b, b.BorrowBy())
	} 
	return res
}