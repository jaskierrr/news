package usecases

import (
	"context"
	"main/internal/models"
)

func (u *usecase) CreateNews(ctx context.Context, params models.News) (models.News, error) {
	news, err := u.repo.CreateNews(ctx, params)
	if err != nil {
		return models.News{}, err
	}
	return news, nil
}
