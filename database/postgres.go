package database

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/cristiangar0398/leal-rewards/internal/models"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db}, nil
}

func (repo *PostgresRepository) InsertUser(ctx context.Context, user *models.User) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO users (id , email , password) VALUES ($1, $2 ,$3)", user.Id, user.Email, user.Password)
	return err
}

func (repo *PostgresRepository) Close() error {
	return repo.db.Close()
}
