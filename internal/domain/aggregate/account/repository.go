package account_aggregate

import (
	"context"
	"errors"
)

var ErrAccountAlreadyExists = errors.New("account already exists")

type AccountRepository interface {
	SignUpAccount(ctx context.Context, account Account) error
}
