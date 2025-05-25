package user_login_model

type UserLogin string

func NewUserLogin(login string) UserLogin { return UserLogin(login) }
