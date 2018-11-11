package sql

import "fmt"

type AccessUser struct {

}

func NewAccessUser() (s *AccessUser){
	s = &AccessUser{}
	return
}

func (s *AccessUser) Insert(u *User) {
	fmt.Println("Insert one user to AccessUser server!")
}

func (s *AccessUser) GetUser(i int) (u *User){
	fmt.Println("Get one user table from AccessUser server according to ID!")
	return
}

type AccessDepartment struct {

}

func NewAccessDepartment() (s *AccessDepartment){
	s = &AccessDepartment{}
	return
}

func (s *AccessDepartment) Insert(u *Department) {
	fmt.Println("Insert one Department to AccessUser server!")
}

func (s *AccessDepartment) GetDepartment(i int) (u *Department){
	fmt.Println("Get one Department table from AccessUser server according to ID!")
	return
}
