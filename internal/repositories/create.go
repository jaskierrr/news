package repositories

import (
	"context"
	"main/internal/models"
)

func (r *repository) CreateNews(ctx context.Context, params models.News) (models.News, error) {

	r.db.GetConn()
	return models.News{}, nil
}
