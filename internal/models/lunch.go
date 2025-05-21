package models

import (
	"time"

	"github.com/lib/pq"
)

type UpdateAction string

const (
	Add UpdateAction = "array_append"
	Del UpdateAction = "array_remove"
)

type Lunch struct {
	ID                   int64         `gorm:"primaryKey;column:id"`
	CreatorID            int64         `gorm:"column:creator_id;not null"`
	Creator              User          `gorm:"foreignKey:CreatorID;references:ID"`
	Place                string        `gorm:"column:place;not null"`
	Time                 time.Time     `gorm:"column:time;not null"`
	NumberOfParticipants int64         `gorm:"column:number_of_participants;not null"`
	Description          string        `gorm:"column:description"`
	Participants         pq.Int64Array `gorm:"type:bigint[];column:participants;not null"`
	LikedBy              pq.Int64Array `gorm:"type:bigint[];column:liked_by;not null"`
}

type LunchFeedback struct {
	Lunch   *Lunch
	IsLiked bool
}
