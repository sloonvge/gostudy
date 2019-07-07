package main

import (
	"fmt"
	"math/rand"
	"bilibili-18daygo/src/day5/exercise/library"
)
/*
1. 实现一个图书管理系统，具有以下功能：
	a. 书籍录入功能，书籍信息包括书名、副本数、作者、出版日期 struct, slice
	b. 书籍查询功能，按照书名、作者、出版日期等条件检索 
	c. 学生信息管理系统，管理每个学生的姓名、年级、身份证、性别、借书信息等 
	d. 借书功能，学生可以查询想要的书籍，进行借出
	e. 书籍管理功能，可以看到每种书被哪些人借出了
*/
var lib library.Library
var stu0, stu1 *library.Student

func init() {
	nbook := 5
	for i := 0; i < nbook; i++ {
		book := library.Book{
			Name: fmt.Sprintf("book%d", i),
			Copy: rand.Intn(10),
			Author: fmt.Sprintf("author%d", i),
			Publication: fmt.Sprintf("publication%d", i),
		}
		lib = append(lib, &book)
	}

	nstu := 2
	stus := make([]*library.Student, 0)
	for i := 0; i < nstu; i++ {
		s := library.NewStudent()
		s.Name = fmt.Sprintf("stu%d", i)
		s.Sex = "male"
		s.Class = fmt.Sprintf("class%d", rand.Intn(6))
		s.IdCard = rand.Intn(9999)
		stus = append(stus, s)
	}
	stu0 = stus[0]
	stu1 = stus[1]
}

func main() {
	fmt.Println(lib.String())
	fmt.Printf("query name: %v\n", lib.Quary("name", "book1"))
	fmt.Printf("query name: %v\n", lib.Quary("name", "-1"))
	fmt.Printf("query author: %v\n", lib.Quary("author", "author1"))
	fmt.Printf("query author: %v\n", lib.Quary("author", "-1"))
	fmt.Printf("query pub: %v\n", lib.Quary("publication", "publication1"))
	fmt.Printf("query pub: %v\n", lib.Quary("publication", "-1"))

	stu0.Borrow(lib[0])
	stu1.Borrows(lib.Quary("name", "book2"))

	fmt.Println(lib.BorrowInfo())
}