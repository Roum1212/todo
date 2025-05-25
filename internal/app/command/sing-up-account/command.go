package sing_up_account_command

import (
	account_login_model "github.com/Roum1212/todo/internal/domain/model/account-login"
	account_password_model "github.com/Roum1212/todo/internal/domain/model/account-password"
)

type Command struct {
	accountLogin    account_login_model.AccountLogin
	accountPassword account_password_model.UserPassword
}

func NewCommand(accountLogin account_login_model.AccountLogin, accountPassword account_password_model.UserPassword) Command {
	return Command{
		accountLogin:    accountLogin,
		accountPassword: accountPassword,
	}
}
