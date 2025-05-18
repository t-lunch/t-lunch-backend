package models

type User struct {
	ID             int64  `gorm:"primaryKey;column:id"`
	Name           string `gorm:"column:name;not null"`
	Surname        string `gorm:"column:surname;not null"`
	Tg             string `gorm:"column:tg;unique;not null"`
	Office         string `gorm:"column:office;not null"` //int64
	Emoji          string `gorm:"column:emoji;not null"`
	Email          string `gorm:"column:email;unique;not null"`
	HashedPassword string `gorm:"column:password;not null"`
}

type UserResponse struct {
	ID      int64
	Name    string
	Surname string
	Tg      string
	Office  string
	Emoji   string
}
