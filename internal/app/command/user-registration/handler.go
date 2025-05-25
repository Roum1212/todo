package user_registration_command

import (
	"context"

	user_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/user"
)

type Handler struct {
	repository user_aggregate.UserRepository
}

func (x Handler) Handle(ctx context.Context, command Command) error {

	return nil
}

func NewHandler(repository user_aggregate.UserRepository) Handler {
	return Handler{
		repository: repository,
	}
}
