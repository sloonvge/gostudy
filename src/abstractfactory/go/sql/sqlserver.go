package sql

import "fmt"

type SqlServerUser struct {
}

func NewSqlServerUser() (s *SqlServerUser){
	s = &SqlServerUser{}
	return
}

func (s *SqlServerUser) Insert(u *User) {
	fmt.Println("Insert one user to SQL server!")
}

func (s *SqlServerUser) GetUser(i int) (u *User){
	fmt.Println("Get one user table from SQL server according to ID!")
	return
}

type SqlServerDepartment struct {
}

func NewSqlServerDepartment() (s *SqlServerDepartment){
	s = &SqlServerDepartment{}
	return
}

func (s *SqlServerDepartment) Insert(u *Department) {
	fmt.Println("Insert one Department to SQL server!")
}

func (s *SqlServerDepartment) GetDepartment(i int) (u *Department){
	fmt.Println("Get one Department table from SQL server according to ID!")
	return
}
