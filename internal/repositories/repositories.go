package repositories

import (
	"context"
	"log/slog"
	"main/internal/database"
	"main/internal/models"

)

type repository struct {
	db     database.DB
	logger *slog.Logger
}

type Repository interface {
	CreateNews(ctx context.Context, params models.News) (models.News, error)
}

func NewUserRepo(db database.DB, logger *slog.Logger) Repository {
	return &repository{
		db:     db,
		logger: logger,
	}
}
