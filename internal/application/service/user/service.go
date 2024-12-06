package user

import (
	"context"
	"test/internal/domain/entity"
	"test/internal/domain/repository"
)

type UserService struct {
	commandHandler *UserCommandHandler
	queryHandler   *UserQueryHandler
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		commandHandler: NewUserCommandHandler(repo),
		queryHandler:   NewUserQueryHandler(repo),
	}
}

// Query methods
func (s *UserService) GetUser(id string) (*entity.User, error) {
	query := GetUserByIDQuery{ID: id}
	return s.queryHandler.HandleGetByID(context.Background(), query)
}

// Command methods
func (s *UserService) SaveUser(user *entity.User) error {
	cmd := CreateUserCommand{
		ID:   user.ID,
		Name: user.Name,
	}
	return s.commandHandler.HandleCreate(context.Background(), cmd)
}

func (s *UserService) UpdateUser(user *entity.User) error {
	cmd := UpdateUserCommand{
		ID:   user.ID,
		Name: user.Name,
	}
	return s.commandHandler.HandleUpdate(cmd)
}
