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
	empty *emptypb.Empty,
) (*reminder_v1.GetAllRemindersResponse, error) {
	reminders, err := x.queryHandler.HandleQuery(ctx)
	if err != nil {
		switch {
		case errors.Is(err, get_all_reminders_query.ErrReminderNotFound):
			return nil, status.Errorf(codes.NotFound, "reminders not found: %v", err)
		default:
			return nil, status.Errorf(codes.Internal, "internal error: %v", err)
		}
	}

	return &reminder_v1.GetAllRemindersResponse{
		Reminders: ToProtoReminders(reminders),
	}, nil
}

func ToProtoReminders(reminders []reminder_aggregate.Reminder) []*reminder_v1.Reminder {
	remindersProto := make([]*reminder_v1.Reminder, len(reminders))
	for i := range reminders {
		remindersProto[i] = &reminder_v1.Reminder{
			Id:          int64(reminders[i].GetID()),
			Title:       string(reminders[i].GetTitle()),
			Description: string(reminders[i].GetDescription()),
		}
	}

	return remindersProto
}

func NewGetAllRemindersRPC(queryHandler get_all_reminders_query.QueryHandler) GetAllRemindersRPC {
	return GetAllRemindersRPC{
		queryHandler: queryHandler,
	}
}
