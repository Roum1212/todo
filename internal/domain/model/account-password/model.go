package account_password_model

type AccountPassword string

func NewAccountPassword(s string) AccountPassword {
	return AccountPassword(s)
}
