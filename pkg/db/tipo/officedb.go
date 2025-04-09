package tipo

import (
	"errors"
	"sync"

	"github.com/t-lunch/t-lunch-backend/internal/domain/models"
)

var (
	ErrOfficeDoesNotExist = errors.New("office does not exist")
)

type Offices struct {
	key  int
	data map[int]*models.Office
	mu   sync.Mutex
}

func NewOffices() *Offices {
	return &Offices{data: make(map[int]*models.Office)}
}

func (o *Offices) AddOffice(address string) int {
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

func (o *Offices) GetOffice(id int) (*models.Office, error) {
	o.mu.Lock()
	defer o.mu.Unlock()

	office, ok := o.data[id]
	if !ok {
		return nil, ErrOfficeDoesNotExist
	}
	return office, nil
}

func (o *Offices) UpdateOffice(id int, address string) (*models.Office, error) {
	office, err := o.GetOffice(id)
	if err != nil {
		return nil, err
	}

	o.mu.Lock()
	defer o.mu.Unlock()

	office.Address = address
	o.data[id] = office
	return office, nil
}

func (o *Offices) DeleteOffice(id int) (bool, error) {
	_, err := o.GetOffice(id)
	if err != nil {
		return false, err
	}

	o.mu.Lock()
	defer o.mu.Unlock()

	delete(o.data, id)
	return true, nil
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
