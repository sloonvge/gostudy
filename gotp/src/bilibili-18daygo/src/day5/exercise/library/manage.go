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
}

func ChangeStudent(s *Student, name string) {
	s.Name = name
} 