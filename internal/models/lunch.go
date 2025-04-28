package models

import "time"

type Lunch struct {
	ID                   int64
	CreatorName          string
	CreatorSurname       string
	Place                string
	Time                 time.Duration
	NumberOfParticipants int64
	Description          *string
	Participants         []*User
}
