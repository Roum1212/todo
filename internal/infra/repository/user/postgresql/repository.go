package postgresql_user_repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	user_password_model "github.com/Roum1212/todo/internal/domain/model/user-password"
	user_login_model "github.com/Roum1212/todo/internal/domain/model/usre-login"
)

type Repository struct {
	client *pgxpool.Pool
}

func (x Repository) UserRegistration(
	ctx context.Context,
	login user_login_model.UserLogin,
	password user_password_model.UserPassword,
) error {

	return nil
}

func NewRepository(client *pgxpool.Pool) Repository {
	return Repository{
		client: client,
	}
}
