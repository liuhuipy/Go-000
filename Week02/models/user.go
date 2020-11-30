package models

type User struct {
	ID int
	Name string
	Age int
}

func (user *User) TableName() string {
	return "user"
}