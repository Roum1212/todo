package account_login_model

type AccountLogin string

func NewAccountLogin(login string) AccountLogin {
	return AccountLogin(login)
}
