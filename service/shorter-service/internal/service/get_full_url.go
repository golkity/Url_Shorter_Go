package service

import "context"

func (s *urlService) GetFullURL(ctx context.Context, shortURL string) (string, error) {
	return s.repo.GetURL(ctx, shortURL)
}
