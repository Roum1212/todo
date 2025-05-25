package user_password_model

type UserPassword string

func NewUserPassword(password string) UserPassword { return UserPassword(password) }
