package sign_up_account_command

import (
	account_login_model "github.com/Roum1212/todo/internal/domain/model/account-login"
	account_password_model "github.com/Roum1212/todo/internal/domain/model/account-password"
)

type Command struct {
	login    account_login_model.AccountLogin
	password account_password_model.AccountPassword
}

func NewCommand(login account_login_model.AccountLogin, password account_password_model.AccountPassword) Command {
	return Command{
		login:    login,
		password: password,
	}
}
