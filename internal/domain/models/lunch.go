package models

type Lunch struct {
	ID                   int
	Creator              int
	Time                 string
	Place                string
	Optional             string
	Participants         []int
	NumberOfParticipants int
}
