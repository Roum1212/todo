package account_password_model

type AccountPassword string

func NewAccountPassword(password string) AccountPassword {
	return AccountPassword(password)
}
