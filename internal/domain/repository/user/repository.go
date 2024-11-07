package user

import "github.com/AnthonyViniciusMuller/OAuth/internal/domain/entity"

type Repository interface {
	GetByUserName(username string) (entity.User, error)
}
