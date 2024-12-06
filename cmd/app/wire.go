//go:build wireinject
// +build wireinject

package main

import (
	"test/internal/application/service/user"
	"test/internal/domain/repository"
	"test/internal/infrastructure/api"
	userHandler "test/internal/infrastructure/api/user"
	"test/internal/infrastructure/persistance"
	userRepo "test/internal/infrastructure/persistance/user"

	"github.com/google/wire"
)

func InitializeAPI(config persistance.Config) (*api.Router, error) {
	wire.Build(
		persistance.NewConnection,
		userRepo.NewUserRepository,
		wire.Bind(new(repository.UserRepository), new(*userRepo.UserRepository)),
		user.NewUserService,
		userHandler.NewHandler,
		api.NewRouter,
	)
	return nil, nil
}
