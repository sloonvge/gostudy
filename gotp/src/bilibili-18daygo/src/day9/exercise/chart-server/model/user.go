package model

const (
	UserStatusOnline = 1 
	UserStatusOffline = 0
)

type User struct {
	UserId int
	Passwd string
	Status int
	LastLogin string
}