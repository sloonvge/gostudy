package model

import (
	"errors"
	"fmt"
)

var (
	ErrUserExist = errors.New("user already exist.")
	ErrUsernameNotFound = errors.New("username not correct")
)

type User struct {
	Username string
	Password string
	StudentId int
}

type Users []*User

func (u Users) AddUser(user *User) {
	u = append(u, user)
}

func (u Users) isExist(user *User) bool {
	for _, v := range u {
		if v.StudentId == user.StudentId {
			return true
		}
	}
	return false
}

func CreateUser(username, password string, stuId int) (u *User) {
	u = &User{
		Username: username,
		Password: password,
		StudentId: stuId,
	}
	return
}

func (u Users) Register(username, password string, stuId int) (err error) {
	for _, v := range u {
		if u.isExist(v) {
			err = ErrUserExist
		}
		return
	}
	user := CreateUser(username, password, stuId)
	u.AddUser(user)
	return
}

func (u Users) Login(username, password string) (err error) {
	for _, v := range u {
		if v.Username == username && v.Password == password {
			fmt.Println("login success")
			return 
		}
	}
	err = ErrUsernameNotFound
	return
}