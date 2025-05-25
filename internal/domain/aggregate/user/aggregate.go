package user_aggregate

import (
	user_password_model "github.com/Roum1212/todo/internal/domain/model/user-password"
	user_login_model "github.com/Roum1212/todo/internal/domain/model/usre-login"
)

type User struct {
	login    user_login_model.UserLogin
	password user_password_model.UserPassword
}

func (x User) GetPassword() user_password_model.UserPassword {
	return x.password
}

func (x User) GetLogin() user_login_model.UserLogin {
	return x.login
}

func NewUser(login user_login_model.UserLogin, password user_password_model.UserPassword) User {
	return User{
		login:    login,
		password: password,
	}
}
