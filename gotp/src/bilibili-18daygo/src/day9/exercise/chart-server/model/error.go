package model

import "errors"

var (
	ErrUserNotExist = errors.New("user not exist")
	ErrUserExist = errors.New("user exist")
	ErrInvalidPasswd = errors.New("")
	ErrInvalidParams = errors.New("")
)
