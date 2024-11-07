package user

import (
	"errors"

	"github.com/AnthonyViniciusMuller/OAuth/internal/domain/entity"
)

type InMemoryRepository struct {
	users map[string]entity.User
}

func New() *InMemoryRepository {
	return &InMemoryRepository{
		users: map[string]entity.User{
			"teste": {
				ID:       1,
				Username: "teste",
				Password: "teste_password",
			},
		},
	}
}

func (repository *InMemoryRepository) GetByUserName(username string) (entity.User, error) {
	user, ok := repository.users[username]
	if !ok {
		return entity.User{}, errors.New("user not found")
	}

	return user, nil
}
