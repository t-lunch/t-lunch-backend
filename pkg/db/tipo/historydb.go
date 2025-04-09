package tipo

import (
	"errors"
	"sync"

	"github.com/t-lunch/t-lunch-backend/internal/domain/models"
)

var (
	ErrHistoryDoesNotExist = errors.New("history does not exist")
)

type Histories struct {
	key  int
	data map[int]*models.History
	mu   sync.Mutex
}

func NewHistories() *Histories {
	return &Histories{data: make(map[int]*models.History)}
}

func (h *Histories) AddHistory(user, lunch int, isLiked bool) int {
	h.mu.Lock()
	defer h.mu.Unlock()

	history := &models.History{
		ID:      h.key,
		User:    user,
		Lunch:   lunch,
		IsLiked: isLiked,
	}
	h.key++

	h.data[history.ID] = history
	return history.ID
}

func (h *Histories) GetHistory(id int) (*models.History, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	history, ok := h.data[id]
	if !ok {
		return nil, ErrHistoryDoesNotExist
	}
	return history, nil
}

func (h *Histories) UpdateHistory(id, user, lunch int, isLiked bool) (*models.History, error) {
	history, err := h.GetHistory(id)
	if err != nil {
		return nil, err
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	history.User = user
	history.Lunch = lunch
	history.IsLiked = isLiked

	h.data[id] = history
	return history, nil
}

func (h *Histories) DeleteHistory(id int) (bool, error) {
	_, err := h.GetHistory(id)
	if err != nil {
		return false, err
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	delete(h.data, id)
	return true, nil
}

func (h *Histories) ListHistory() []*models.History {
	h.mu.Lock()
	defer h.mu.Unlock()

	histories := make([]*models.History, 0, len(h.data))
	for _, history := range h.data {
		histories = append(histories, history)
	}
	return histories
}
