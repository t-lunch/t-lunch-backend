package tipo

import (
	"errors"
	"sync"

	"github.com/t-lunch/t-lunch-backend/internal/models"
)

var (
	ErrUserDoesNotExist = errors.New("user does not exist")
)

type Users struct {
	key  int64
	data map[int64]*models.User
	mu   sync.Mutex
}

func NewUsers() *Users {
	return &Users{data: make(map[int64]*models.User)}
}

func (u *Users) AddUser(name, surname, tg, login, password string, office int64) int64 {
	u.mu.Lock()
	defer u.mu.Unlock()

	user := &models.User{
		ID:       u.key,
		Name:     name,
		Surname:  surname,
		Tg:       tg,
		Office:   office,
		Login:    login,
		Password: password,
	}
	u.key++

	u.data[user.ID] = user
	return user.ID
}

func (u *Users) GetUser(id int64) (*models.User, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	user, ok := u.data[id]
	if !ok {
		return nil, ErrUserDoesNotExist
	}
	return user, nil
}

func (u *Users) ListUsers() []*models.User {
	u.mu.Lock()
	defer u.mu.Unlock()

	users := make([]*models.User, 0, len(u.data))
	for _, user := range u.data {
		users = append(users, user)
	}
	return users
}
