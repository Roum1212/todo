package get_all_reminders_rpc

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	get_all_reminders_query "github.com/Roum1212/todo/internal/app/query/get-all-reminders"
	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
	reminder_v1 "github.com/Roum1212/todo/pkg/gen/reminder/v1"
)

type GetAllRemindersRPC struct {
	queryHandler get_all_reminders_query.QueryHandler
}

func (x GetAllRemindersRPC) GetAllReminders(
	ctx context.Context,
	_ *emptypb.Empty,
) (*reminder_v1.GetAllRemindersResponse, error) {
	reminders, err := x.queryHandler.HandleQuery(ctx)
	if err != nil {
		switch {
		case errors.Is(err, get_all_reminders_query.ErrReminderNotFound):
			return nil, status.Errorf(codes.NotFound, "reminders not found: %v", err)
		default:
			return nil, status.Errorf(codes.Internal, "failed to get all reminders: %v", err)
		}
	}

	return &reminder_v1.GetAllRemindersResponse{
		Reminders: NewReminderDTOs(reminders),
	}, nil
}

func NewReminderDTOs(reminders []reminder_aggregate.Reminder) []*reminder_v1.Reminder {
	pbReminders := make([]*reminder_v1.Reminder, len(reminders))
	for i := range reminders {
		pbReminders[i] = &reminder_v1.Reminder{
			Id:          int64(reminders[i].GetID()),
			Title:       string(reminders[i].GetTitle()),
			Description: string(reminders[i].GetDescription()),
		}
	}

	return pbReminders
}

func NewGetAllRemindersRPC(queryHandler get_all_reminders_query.QueryHandler) GetAllRemindersRPC {
	return GetAllRemindersRPC{
		queryHandler: queryHandler,
	}
}
