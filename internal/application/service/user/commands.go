package user

import (
	"context"
	"time"

	"test/internal/domain/entity"
	"test/internal/domain/repository"
	"test/internal/infrastructure/metrics"

	"github.com/rs/zerolog/log"
)

// Commands
type CreateUserCommand struct {
	ID   string
	Name string
}

type UpdateUserCommand struct {
	ID   string
	Name string
}

// Command Handlers
type UserCommandHandler struct {
	repo repository.UserRepository
}

func NewUserCommandHandler(repo repository.UserRepository) *UserCommandHandler {
	return &UserCommandHandler{repo: repo}
}

func (h *UserCommandHandler) HandleCreate(ctx context.Context, cmd CreateUserCommand) error {
	logger := log.With().
		Str("command", "CreateUser").
		Str("user_id", cmd.ID).
		Logger()

	startTime := time.Now()
	logger.Info().Msg("Starting CreateUser command")

	var err error
	defer func() {
		duration := time.Since(startTime).Seconds()
		status := "success"
		if err != nil {
			status = "error"
		}
		metrics.CommandDuration.WithLabelValues("CreateUser", status).Observe(duration)
		metrics.CommandCounter.WithLabelValues("CreateUser", status).Inc()
	}()

	if err := cmd.Validate(); err != nil {
		logger.Error().Err(err).Msg("Command validation failed")
		return err
	}

	user := &entity.User{
		ID:   cmd.ID,
		Name: cmd.Name,
	}

	if err := h.repo.SaveUser(user); err != nil {
		logger.Error().Err(err).Msg("Failed to save user")
		return err
	}

	logger.Info().Msg("User created successfully")
	return nil
}

func (h *UserCommandHandler) HandleUpdate(cmd UpdateUserCommand) error {
	user := &entity.User{
		ID:   cmd.ID,
		Name: cmd.Name,
	}
	return h.repo.SaveUser(user)
}
