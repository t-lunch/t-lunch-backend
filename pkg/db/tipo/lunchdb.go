package tipo

import (
	"errors"
	"sync"
	"time"

	"github.com/t-lunch/t-lunch-backend/internal/models"
)

var (
	ErrLunchDoesNotExist = errors.New("lunch does not exist")
)

type Lunches struct {
	key  int64
	data map[int64]*models.Lunch
	mu   sync.Mutex
}

func NewLunches() *Lunches {
	return &Lunches{data: make(map[int64]*models.Lunch)}
}

func (l *Lunches) AddLunch(creator int64, time time.Duration, place, optional string, participants []int64) int64 {
	l.mu.Lock()
	defer l.mu.Unlock()

	lunch := &models.Lunch{
		ID:                   l.key,
		Creator:              creator,
		Time:                 time,
		Place:                place,
		Optional:             optional,
		Participants:         participants,
		NumberOfParticipants: int64(len(participants)),
	}
	l.key++

	l.data[lunch.ID] = lunch
	return lunch.ID
}

func (l *Lunches) GetLunch(id int64) (*models.Lunch, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	lunch, ok := l.data[id]
	if !ok {
		return nil, ErrLunchDoesNotExist
	}
	return lunch, nil
}

func (l *Lunches) UpdateLunch(id, creator int64, time time.Duration, place, optional string, participants []int64) (*models.Lunch, error) {
	lunch, err := l.GetLunch(id)
	if err != nil {
		return nil, err
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	lunch.Creator = creator
	lunch.Time = time
	lunch.Place = place
	lunch.Optional = optional
	lunch.Participants = participants
	lunch.NumberOfParticipants = int64(len(participants))

	l.data[id] = lunch
	return lunch, nil
}

func (l *Lunches) ListLunches() []*models.Lunch {
	l.mu.Lock()
	defer l.mu.Unlock()

	lunches := make([]*models.Lunch, 0, len(l.data))
	for _, lunch := range l.data {
		lunches = append(lunches, lunch)
	}
	return lunches
}
