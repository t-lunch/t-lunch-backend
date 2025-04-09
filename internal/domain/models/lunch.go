package models

import "time"

type Lunch struct {
	ID                   int64
	Creator              int64
	Time                 time.Duration
	Place                string
	Optional             string
	Participants         []int64
	NumberOfParticipants int64
}
