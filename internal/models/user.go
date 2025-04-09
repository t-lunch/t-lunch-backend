package models

type User struct {
	ID       int64
	Name     string
	Surname  string
	Tg       string
	Office   int64
	Login    string
	Password string
	Emoji    string
}
