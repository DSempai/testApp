package user

import (
	"context"
	"time"

	"test/internal/domain/entity"
	"test/internal/domain/repository"
	"test/internal/infrastructure/metrics"

	"github.com/rs/zerolog/log"
)

// Queries
type GetUserByIDQuery struct {
	ID string
}

// Query Handlers
type UserQueryHandler struct {
	repo repository.UserRepository
}

func NewUserQueryHandler(repo repository.UserRepository) *UserQueryHandler {
	return &UserQueryHandler{repo: repo}
}

func (h *UserQueryHandler) HandleGetByID(ctx context.Context, query GetUserByIDQuery) (*entity.User, error) {
	logger := log.With().
		Str("query", "GetUserByID").
		Str("user_id", query.ID).
		Logger()

	startTime := time.Now()
	logger.Info().Msg("Starting GetUserByID query")

	var err error
	defer func() {
		duration := time.Since(startTime).Seconds()
		status := "success"
		if err != nil {
			status = "error"
		}
		metrics.QueryDuration.WithLabelValues("GetUserByID", status).Observe(duration)
		metrics.QueryCounter.WithLabelValues("GetUserByID", status).Inc()
	}()

	user, err := h.repo.GetUserByID(query.ID)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get user")
		return nil, err
	}

	logger.Info().Msg("User retrieved successfully")
	return user, nil
}
