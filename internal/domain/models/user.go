package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Tg       string `json:"tg"`
	Office   int    `json:"office"` // Foreign key to Offices
	Login    string `json:"login"`
	Password string `json:"password"`
}
