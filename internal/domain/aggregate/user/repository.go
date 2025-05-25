package user_aggregate

import (
	"context"

	user_password_model "github.com/Roum1212/todo/internal/domain/model/user-password"
	user_login_model "github.com/Roum1212/todo/internal/domain/model/usre-login"
)

type UserRepository interface {
	UserRegistration(
		ctx context.Context,
		login user_login_model.UserLogin,
		password user_password_model.UserPassword,
	) error
}
