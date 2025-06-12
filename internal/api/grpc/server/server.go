package server_grpc

import (
	"context"

	create_reminder_rpc "github.com/Roum1212/todo/internal/api/grpc/rpc/create-reminder"
	create_reminder_command "github.com/Roum1212/todo/internal/app/command/create-reminder"
	reminder_v1 "github.com/Roum1212/todo/pkg/gen/reminder/v1"
)

type ReminderServiceServer struct {
	reminder_v1.UnimplementedReminderServiceServer

	CreateReminderRPC create_reminder_rpc.CreateReminderRPC
}

func (x ReminderServiceServer) CreateReminder(
	ctx context.Context,
	r *reminder_v1.CreateReminderRequest,
) (*reminder_v1.CreateReminderResponse, error) {
	createReminderResponse, err := x.CreateReminderRPC.CreateReminder(ctx, r)
	if err != nil {
		return nil, err
	}

	return createReminderResponse, nil
}

func NewCreateReminderService(commandHandler create_reminder_command.CommandHandler) ReminderServiceServer {
	return ReminderServiceServer{
		CreateReminderRPC: create_reminder_rpc.NewCreateReminderRPC(commandHandler),
	}
}
