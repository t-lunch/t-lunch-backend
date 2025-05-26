package errors

import (
	"errors"
	"fmt"
)

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrMissingAuthToken    = errors.New("missing auth token")
)

var (
	ErrConfigIsNil      = errors.New("config is nil")
	ErrDBIsNil          = errors.New("db is nil")
	ErrUnknownTokenType = errors.New("unknown token type")
	ErrInvalidToken     = errors.New("invalid token")
)

var (
	ErrInvalidRequest  = errors.New("invalid request")
	ErrInvalidPassword = errors.New("invalid password")
	ErrTokenExpired    = errors.New("token expired")
)

type ErrRepository struct {
	Repo   string
	Method string
	Err    error
}

func NewErrRepository(repo, method string, err error) ErrRepository {
	return ErrRepository{
		Repo:   repo,
		Method: method,
		Err:    err,
	}
}

func (e ErrRepository) Error() string {
	return fmt.Sprintf("error in %s, method %s: %v", e.Repo, e.Method, e.Err)
}

type ErrUserWithEmailAlreadyExists struct {
	Email string
}

func NewErrUserWithEmailAlreadyExists(email string) ErrUserWithEmailAlreadyExists {
	return ErrUserWithEmailAlreadyExists{
		Email: email,
	}
}

func (e ErrUserWithEmailAlreadyExists) Error() string {
	return fmt.Sprintf("user with email [%s] already exists", e.Email)
}

type ErrUserAndOwnerAreDifferent struct {
	User  int64
	Owner int64
}

func NewErrUserAndOwnerAreDifferent(user, owner int64) ErrUserAndOwnerAreDifferent {
	return ErrUserAndOwnerAreDifferent{
		User:  user,
		Owner: owner,
	}
}

func (e ErrUserAndOwnerAreDifferent) Error() string {
	return fmt.Sprintf("user [%d] and owner [%d] are different", e.User, e.Owner)
}
