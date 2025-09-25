package repository

import "context"

type UserRepository interface {
	CreateUser(ctx context.Context, userID string) error
}

type URLRepository interface {
	SaveURL(ctx context.Context, fullURL, shortURL, userID string) error
	GetURL(ctx context.Context, shortURL string) (string, error)
	GetURLsByUserID(ctx context.Context, userID string) ([]string, error)
}
