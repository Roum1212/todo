package account_aggregate

import (
	account_login_model "github.com/Roum1212/todo/internal/domain/model/account-login"
	account_password_model "github.com/Roum1212/todo/internal/domain/model/account-password"
)

type Account struct {
	login    account_login_model.AccountLogin
	password account_password_model.AccountPassword
}

func (x Account) GetLogin() account_login_model.AccountLogin {
	return x.login
}

func (x Account) GetPassword() account_password_model.AccountPassword {
	return x.password
}

func NewAccount(login account_login_model.AccountLogin, password account_password_model.AccountPassword) Account {
	return Account{
		login:    login,
		password: password,
	}
}
