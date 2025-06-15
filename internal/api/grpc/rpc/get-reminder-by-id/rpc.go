package get_reminder_by_id_rpc

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	get_reminder_by_id_query "github.com/Roum1212/todo/internal/app/query/get-reminder-by-id"
	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
	reminder_v1 "github.com/Roum1212/todo/pkg/gen/reminder/v1"
)

type GetReminderByIDRPC struct {
	queryHandler get_reminder_by_id_query.QueryHandler
}

func (x GetReminderByIDRPC) GetReminderByID(
	ctx context.Context,
	r *reminder_v1.GetReminderByIDRequest,
) (*reminder_v1.GetReminderByIDResponse, error) {
	reminderID, err := reminder_id_model.NewReminderID(r.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid reminder id: %v", err)
	}

	reminder, err := x.queryHandler.HandleQuery(
		ctx,
		get_reminder_by_id_query.NewQuery(reminderID),
	)
	if err != nil {
		switch {
		case errors.Is(err, get_reminder_by_id_query.ErrReminderNotFound):
			return nil, status.Errorf(codes.NotFound, "reminder not found: %v", err)
		default:
			return nil, status.Errorf(codes.Internal, "failed to get reminder by id: %v", err)
		}
	}

	return &reminder_v1.GetReminderByIDResponse{
		Reminder: newReminderDTO(reminder),
	}, nil
}

func NewGetReminderByIDRPC(queryHandler get_reminder_by_id_query.QueryHandler) GetReminderByIDRPC {
	return GetReminderByIDRPC{
		queryHandler: queryHandler,
	}
}

func newReminderDTO(reminder reminder_aggregate.Reminder) *reminder_v1.Reminder {
	return &reminder_v1.Reminder{
		Id:          int64(reminder.GetID()),
		Title:       string(reminder.GetTitle()),
		Description: string(reminder.GetDescription()),
	}
}
