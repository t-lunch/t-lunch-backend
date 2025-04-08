package models

type Lunch struct {
	ID                   int    `json:"id"`
	Creator              int    `json:"creator"` // Foreign key to Users
	Time                 string `json:"time"`
	Place                string `json:"place"`
	Optional             string `json:"optional,omitempty"`
	Participants         []int  `json:"participants"` // List of user IDs
	NumberOfParticipants int    `json:"number_of_participants"`
}
