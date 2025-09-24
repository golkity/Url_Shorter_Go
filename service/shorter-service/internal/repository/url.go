package repository

import (
	"context"
)

type URLRepository interface {
	SaveURL(ctx context.Context, fullURL, shortURL string) error
	GetURL(ctx context.Context, shortURL string) (string, error)
}
