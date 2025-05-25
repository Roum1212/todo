package user_registration_command

import (
	user_password_model "github.com/Roum1212/todo/internal/domain/model/user-password"
	user_login_model "github.com/Roum1212/todo/internal/domain/model/usre-login"
)

type Command struct {
	userLogin    user_login_model.UserLogin
	userPassword user_password_model.UserPassword
}

func NewCommand(userLogin user_login_model.UserLogin, userPassword user_password_model.UserPassword) Command {
	return Command{
		userLogin:    userLogin,
		userPassword: userPassword,
	}
}
