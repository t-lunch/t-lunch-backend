package models

type User struct {
	ID             int64
	Name           string
	Surname        string
	Tg             string
	Office         string //int64
	Emoji          string
	Email          string
	HashedPassword string
}
