package sing_up_account_command

import (
	"context"
	"fmt"

	account_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/account"
)

type Handler struct {
	repository account_aggregate.AccountRepository
}

func (x Handler) Handle(ctx context.Context, command Command) error {
	account := account_aggregate.NewAccount(command.accountLogin, command.accountPassword)

	if err := x.repository.SingUpAccount(ctx, account); err != nil {
		return fmt.Errorf("faild to sing up account %w", err)
	}

	return nil
}

func NewHandler(repository account_aggregate.AccountRepository) Handler {
	return Handler{
		repository: repository,
	}
}
