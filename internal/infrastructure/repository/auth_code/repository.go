package authcode

import (
	"errors"
)

var ErrAuthCodeNotFound = errors.New("auth code not found")

type InMemoryRepository struct {
	authCodes map[string]int64
}

func New() *InMemoryRepository {
	return &InMemoryRepository{
		authCodes: make(map[string]int64),
	}
}

func (repository *InMemoryRepository) Insert(code string, userID int64) error {
	repository.authCodes[code] = userID

	return nil
}

func (repository *InMemoryRepository) GetUserID(code string) (int64, error) {
	userID, ok := repository.authCodes[code]
	if !ok {
		return 0, ErrAuthCodeNotFound
	}

	return userID, nil
}

func (repository *InMemoryRepository) Remove(code string) error {
	_, ok := repository.authCodes[code]
	if !ok {
		return ErrAuthCodeNotFound
	}

	delete(repository.authCodes, code)

	return nil
}
