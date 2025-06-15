package delete_reminder_rpc

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	delete_reminder_command "github.com/Roum1212/todo/internal/app/command/delete-reminder"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
	reminder_v1 "github.com/Roum1212/todo/pkg/gen/reminder/v1"
)

type DeleteReminderRPC struct {
	commandHandler delete_reminder_command.CommandHandler
}

func (x DeleteReminderRPC) DeleteReminder(
	ctx context.Context,
	r *reminder_v1.DeleteReminderRequest,
) (*emptypb.Empty, error) {
	reminderID, err := reminder_id_model.NewReminderID(r.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid reminderID: %v", err)
	}

	if err = x.commandHandler.HandleCommand(
		ctx,
		delete_reminder_command.NewCommand(reminderID),
	); err != nil {
		switch {
		case errors.Is(err, delete_reminder_command.ErrReminderNotFound):
			return nil, status.Errorf(codes.NotFound, "reminder not found: %v", err)
		default:
			return nil, status.Errorf(codes.Internal, "failed to delete reminder: %v", err)
		}
	}

	return &emptypb.Empty{}, nil
}

func NewDeleteReminderRPC(commandHandler delete_reminder_command.CommandHandler) DeleteReminderRPC {
	return DeleteReminderRPC{
		commandHandler: commandHandler,
	}
}
