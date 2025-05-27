package postgresql_account_repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

const table = "accounts"

const (
	fieldLogin    = "login"
	fieldPassword = "password"
)

type Repository struct {
	client *pgxpool.Pool
}

func NewRepository(client *pgxpool.Pool) Repository {
	return Repository{
		client: client,
	}
}
