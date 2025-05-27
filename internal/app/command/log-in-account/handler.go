package log_in_account_command

import (
	"context"
	"fmt"

	account_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/account"
)

type Handler struct {
	repository account_aggregate.AccountRepository
}

func (x Handler) Handle(ctx context.Context, c Command) error {

	if err := x.repository.LogInAccount(
		ctx,
		account_aggregate.NewAccount(c.Login, c.Password),
	); err != nil {
		return fmt.Errorf("failed to log in account: %w", err)
	}

	return nil
}

func NewHandler(repository account_aggregate.AccountRepository) Handler {
	return Handler{
		repository: repository,
	}
}
