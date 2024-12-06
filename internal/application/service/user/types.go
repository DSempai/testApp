package user

import (
	"errors"
	"test/internal/domain/entity"
)

// Result types
type UserResult struct {
	User  *entity.User
	Error error
}

// Validation
func (cmd CreateUserCommand) Validate() error {
	if cmd.ID == "" {
		return errors.New("user ID is required")
	}
	if cmd.Name == "" {
		return errors.New("user name is required")
	}
	return nil
}

func (cmd UpdateUserCommand) Validate() error {
	if cmd.ID == "" {
		return errors.New("user ID is required")
	}
	if cmd.Name == "" {
		return errors.New("user name is required")
	}
	return nil
}
