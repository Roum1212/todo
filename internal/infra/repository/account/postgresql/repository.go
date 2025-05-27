package postgresql_account_repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"

	account_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/account"
)

const table = "accounts"

const (
	fieldLogin    = "login"
	fieldPassword = "password"
)

type Repository struct {
	client *pgxpool.Pool
}

func (x Repository) LogInAccount(ctx context.Context, account account_aggregate.Account) error {
	sql, args, err := squirrel.
		

	return nil
}

func NewRepository(client *pgxpool.Pool) Repository {
	return Repository{
		client: client,
	}
}
