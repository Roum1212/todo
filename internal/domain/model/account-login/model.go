package account_login_model

type AccountLogin string

func NewAccountLogin(s string) AccountLogin {
	return AccountLogin(s)
}
