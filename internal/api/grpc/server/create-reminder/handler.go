package create_reminder_grpc_server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	create_reminder_command "github.com/Roum1212/todo/internal/app/command/create-reminder"
	reminder_description_model "github.com/Roum1212/todo/internal/domain/model/reminder-description"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
	reminder_title_model "github.com/Roum1212/todo/internal/domain/model/reminder-title"
	reminder_v1 "github.com/Roum1212/todo/pkg/gen/reminder/v1"
)

type Server struct {
	reminder_v1.UnimplementedReminderServiceServer
	commandHandler create_reminder_command.CommandHandler
}

func (x *Server) CreateReminder(
	ctx context.Context,
	r *reminder_v1.CreateReminderRequest,
) (*reminder_v1.CreateReminderResponse, error) {
	reminderID := reminder_id_model.GenerateReminderID()

	reminderTitle, err := reminder_title_model.NewReminderTitle(r.GetTitle())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid reminderTitle: %v", err)
	}

	reminderDescription, err := reminder_description_model.NewReminderDescription(r.GetDescription())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid reminderDescription: %v", err)
	}

	if err = x.commandHandler.HandleCommand(
		ctx,
		create_reminder_command.NewCommand(reminderID, reminderTitle, reminderDescription),
	); err != nil {
		return nil, status.Errorf(codes.Internal, "internal error: %v", err)
	}

	return &reminder_v1.CreateReminderResponse{
		Id: int64(reminderID),
	}, nil
}

func NewServer(commandHandler create_reminder_command.CommandHandler) Server {
	return Server{
		commandHandler: commandHandler,
	}
}
