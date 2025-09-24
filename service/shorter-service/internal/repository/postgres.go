package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"url-shortener/internal/custom_errors"
)

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) SaveURL(ctx context.Context, fullURL, shortURL string) error {
	query := `INSERT INTO urls (full_url, short_url) VALUES ($1, $2)`
	_, err := r.db.Exec(ctx, query, fullURL, shortURL)
	return err
}

func (r *PostgresRepository) GetURL(ctx context.Context, shortURL string) (string, error) {
	var fullURL string
	query := `SELECT full_url FROM urls WHERE short_url = $1`
	err := r.db.QueryRow(ctx, query, shortURL).Scan(&fullURL)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", custom_errors.Error_URLNotFound
		}
		return "", err
	}

	return fullURL, nil
}
