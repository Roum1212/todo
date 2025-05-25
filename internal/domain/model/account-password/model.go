package account_password_model

type UserPassword string

func NewAccountPassword(password string) UserPassword { return UserPassword(password) }
