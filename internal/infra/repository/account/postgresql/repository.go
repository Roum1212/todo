package postgresql_account_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgconn"
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

func (x Repository) SignUpAccount(
	ctx context.Context,
	account account_aggregate.Account,
) error {
	sql, args, err := squirrel.
		Insert(table).
		Columns(fieldLogin, fieldPassword).
		Values(string(account.GetLogin()), string(account.GetPassword())).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("faild to build sql: %w", err)
	}

	if _, err = x.client.Exec(ctx, sql, args...); err != nil { //nolint:nestif // OK.
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return account_aggregate.ErrAccountAlreadyExists
		}

		return fmt.Errorf("failed to sign up account: %w", err)
	}

	return nil
}

func NewRepository(client *pgxpool.Pool) Repository {
	return Repository{
		client: client,
	}
}
