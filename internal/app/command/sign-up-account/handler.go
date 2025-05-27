package sign_up_account_command

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"

	account_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/account"
)

type Handler struct {
	repository account_aggregate.AccountRepository
}

func (x Handler) Handle(ctx context.Context, command Command) error {
	account := account_aggregate.NewAccount(command.login, command.password)

	if err := x.repository.SignUpAccount(ctx, account); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return account_aggregate.ErrAccountAlreadyExists
		}

		return fmt.Errorf("failed to sign up account: %w", err)
	}

	return nil
}

func NewHandler(repository account_aggregate.AccountRepository) Handler {
	return Handler{
		repository: repository,
	}
}
