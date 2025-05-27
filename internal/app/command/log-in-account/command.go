package log_in_account_command

import (
	account_login_model "github.com/Roum1212/todo/internal/domain/model/account-login"
	account_password_model "github.com/Roum1212/todo/internal/domain/model/account-password"
)

type Command struct {
	Login    account_login_model.AccountLogin
	Password account_password_model.AccountPassword
}

func NewCommand(
	login account_login_model.AccountLogin,
	password account_password_model.AccountPassword,
) Command {
	return Command{
		Login:    login,
		Password: password,
	}
}
