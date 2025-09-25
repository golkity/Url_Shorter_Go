package repository

import (
	"context"
	"errors"
	"url-shortener/internal/custom_errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) SaveURL(ctx context.Context, fullURL, shortURL, userID string) error {
	query := `INSERT INTO urls (full_url, short_url, user_id) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(ctx, query, fullURL, shortURL, userID)
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

func (r *PostgresRepository) CreateUser(ctx context.Context, userID string) error {
	query := `INSERT INTO users (id) VALUES ($1) ON CONFLICT (id) DO NOTHING`
	_, err := r.db.Exec(ctx, query, userID)
	return err
}

func (r *PostgresRepository) GetURLsByUserID(ctx context.Context, userID string) ([]string, error) {
	return nil, nil // TODO: Доделать до 29-го
}
