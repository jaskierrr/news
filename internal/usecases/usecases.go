package usecases

import (
	"context"
	"log/slog"
	"main/internal/models"
	"main/internal/repositories"
)

type usecase struct {
	logger *slog.Logger
	repo   repositories.Repository
}

type Usecases interface {
	CreateNews(ctx context.Context, params models.News) (models.News, error)
}

func New(repo repositories.Repository, logger *slog.Logger) Usecases {
	return &usecase{
		repo:   repo,
		logger: logger,
	}
}
