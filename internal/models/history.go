package models

type History struct {
	ID      int  `json:"id"`
	User    int  `json:"user"`  // Foreign key to Users
	Lunch   int  `json:"lunch"` // Foreign key to Lunches
	IsLiked bool `json:"is_liked"`
}
