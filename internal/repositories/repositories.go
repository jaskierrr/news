package repositories

import (
	"context"
	"log/slog"
	"main/internal/models"

	"gopkg.in/reform.v1"
)

type repository struct {
	db     *reform.DB
	logger *slog.Logger
}

type Repository interface {
	CreateNews(ctx context.Context, params models.News) (models.News, error)
}

func NewUserRepo(db *reform.DB, logger *slog.Logger) Repository {
	return &repository{
		db:     db,
		logger: logger,
	}
}
