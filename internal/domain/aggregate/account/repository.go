package account_aggregate

import (
	"context"
)

type AccountRepository interface {
	SingUpAccount(
		ctx context.Context,
		account Account,
	) error
}
