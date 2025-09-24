package service

import (
	"context"
	"github.com/teris-io/shortid"
)

func (s *urlService) CreateShortURL(ctx context.Context, fullURL string) (string, error) {
	shortURL, err := shortid.Generate()
	if err != nil {
		return "", err
	}

	err = s.repo.SaveURL(ctx, fullURL, shortURL)
	if err != nil {
		return "", err
	}

	return shortURL, nil
}
