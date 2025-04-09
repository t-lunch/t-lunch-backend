package tipo

import (
	"errors"
	"sync"

	"github.com/t-lunch/t-lunch-backend/internal/models"
)

var (
	ErrOfficeDoesNotExist = errors.New("office does not exist")
)

type Offices struct {
	key  int64
	data map[int64]*models.Office
	mu   sync.Mutex
}

func NewOffices() *Offices {
	return &Offices{data: make(map[int64]*models.Office)}
}

func (o *Offices) AddOffice(address string) int64 {
	o.mu.Lock()
	defer o.mu.Unlock()

	office := &models.Office{
		ID:      o.key,
		Address: address,
	}
	o.key++

	o.data[office.ID] = office
	return office.ID
}

func (o *Offices) GetOffice(id int64) (*models.Office, error) {
	o.mu.Lock()
	defer o.mu.Unlock()

	office, ok := o.data[id]
	if !ok {
		return nil, ErrOfficeDoesNotExist
	}
	return office, nil
}

func (o *Offices) ListOffices() []*models.Office {
	o.mu.Lock()
	defer o.mu.Unlock()

	offices := make([]*models.Office, 0, len(o.data))
	for _, office := range o.data {
		offices = append(offices, office)
	}
	return offices
}
