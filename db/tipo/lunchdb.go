package tipo

import (
	"errors"
	"sync"

	"github.com/t-lunch/t-lunch-backend/internal/models"
)

var (
	ErrLunchDoesNotExist = errors.New("lunch does not exist")
)

type Lunches struct {
	key  int
	data map[int]*models.Lunch
	mu   sync.Mutex
}

func NewLunches() *Lunches {
	return &Lunches{data: make(map[int]*models.Lunch)}
}

func (l *Lunches) AddLunch(creator int, time, place, optional string, participants []int) int {
	l.mu.Lock()
	defer l.mu.Unlock()

	lunch := &models.Lunch{
		ID:                   l.key,
		Creator:              creator,
		Time:                 time,
		Place:                place,
		Optional:             optional,
		Participants:         participants,
		NumberOfParticipants: len(participants),
	}
	l.key++

	l.data[lunch.ID] = lunch
	return lunch.ID
}

func (l *Lunches) GetLunch(id int) (*models.Lunch, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	lunch, ok := l.data[id]
	if !ok {
		return nil, ErrLunchDoesNotExist
	}
	return lunch, nil
}

func (l *Lunches) UpdateLunch(id, creator int, time, place, optional string, participants []int) (*models.Lunch, error) {
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
	lunch.NumberOfParticipants = len(participants)

	l.data[id] = lunch
	return lunch, nil
}

func (l *Lunches) DeleteLunch(id int) (bool, error) {
	_, err := l.GetLunch(id)
	if err != nil {
		return false, err
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	delete(l.data, id)
	return true, nil
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
