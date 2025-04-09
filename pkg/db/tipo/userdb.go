package tipo

import (
	"errors"
	"sync"

	"github.com/t-lunch/t-lunch-backend/internal/domain/models"
)

var (
	ErrUserDoesNotExist = errors.New("user does not exist")
)

type Users struct {
	key  int
	data map[int]*models.User
	mu   sync.Mutex
}

func NewUsers() *Users {
	return &Users{data: make(map[int]*models.User)}
}

func (u *Users) AddUser(name, surname, tg, login, password string, office int) int {
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

func (u *Users) GetUser(id int) (*models.User, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	user, ok := u.data[id]
	if !ok {
		return nil, ErrUserDoesNotExist
	}
	return user, nil
}

func (u *Users) UpdateUser(id int, name, surname, tg, login, password string, office int) (*models.User, error) {
	user, err := u.GetUser(id)
	if err != nil {
		return nil, err
	}

	u.mu.Lock()
	defer u.mu.Unlock()

	user.Name = name
	user.Surname = surname
	user.Tg = tg
	user.Office = office
	user.Login = login
	user.Password = password

	u.data[id] = user
	return user, nil
}

func (u *Users) DeleteUser(id int) (bool, error) {
	_, err := u.GetUser(id)
	if err != nil {
		return false, err
	}

	u.mu.Lock()
	defer u.mu.Unlock()

	delete(u.data, id)
	return true, nil
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
