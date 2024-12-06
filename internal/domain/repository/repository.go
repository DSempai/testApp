package repository

import "test/internal/domain/entity"

type UserRepository interface {
	GetUserByID(id string) (*entity.User, error)
	SaveUser(user *entity.User) error
}
