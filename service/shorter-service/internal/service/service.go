package service

import (
	"context"
	"url-shortener/internal/repository"
)

type URLService interface {
	CreateShortURL(ctx context.Context, fullURL, userID string) (string, error)
	GetFullURL(ctx context.Context, shortURL string) (string, error)
}

type urlService struct {
	repo repository.URLRepository
}

func NewURLService(repo repository.URLRepository) *urlService {
	return &urlService{repo: repo}
}
