package server_grpc

import (
	"context"

	create_reminder_rpc "github.com/Roum1212/todo/internal/api/grpc/rpc/create-reminder"
	reminder_v1 "github.com/Roum1212/todo/pkg/gen/reminder/v1"
)

type ReminderServiceServer struct {
	reminder_v1.UnimplementedReminderServiceServer

	createReminderRPC create_reminder_rpc.CreateReminderRPC
}

func (x ReminderServiceServer) CreateReminder(
	ctx context.Context,
	r *reminder_v1.CreateReminderRequest,
) (*reminder_v1.CreateReminderResponse, error) {
	return x.createReminderRPC.CreateReminder(ctx, r)
}

func NewCreateReminderService(createReminderRPC create_reminder_rpc.CreateReminderRPC) ReminderServiceServer {
	return ReminderServiceServer{
		createReminderRPC: createReminderRPC,
	}
}
