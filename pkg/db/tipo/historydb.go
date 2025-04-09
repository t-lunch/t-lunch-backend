package tipo

import (
	"errors"
	"sync"

	"github.com/t-lunch/t-lunch-backend/internal/models"
)

var (
	ErrHistoryDoesNotExist = errors.New("history does not exist")
)

type Histories struct {
	key  int64
	data map[int64]*models.History
	mu   sync.Mutex
}

func NewHistories() *Histories {
	return &Histories{data: make(map[int64]*models.History)}
}

func (h *Histories) AddHistory(user, lunch int64, isLiked bool) int64 {
	h.mu.Lock()
	defer h.mu.Unlock()

	history := &models.History{
		ID:      h.key,
		UserID:  user,
		LunchID: lunch,
		IsLiked: isLiked,
	}
	h.key++

	h.data[history.ID] = history
	return history.ID
}

func (h *Histories) GetHistory(id int64) (*models.History, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	history, ok := h.data[id]
	if !ok {
		return nil, ErrHistoryDoesNotExist
	}
	return history, nil
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
