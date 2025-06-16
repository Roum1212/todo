package redis_reminder_repository

import (
	"context"
	"encoding/base64"
	"fmt"
	"strconv"

	"github.com/redis/rueidis"
	"google.golang.org/protobuf/proto"

	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
	reminder_v1 "github.com/Roum1212/todo/pkg/gen/reminder/v1"
)

const keyPrefix = "reminder:"

type Repository struct {
	client rueidis.Client
}

func (x Repository) SaveReminder(ctx context.Context, reminder reminder_aggregate.Reminder) error {
	reminderDTO := NewReminderDTO(reminder)

	data, err := proto.Marshal(reminderDTO)
	if err != nil {
		return fmt.Errorf("failed to marshal reminder: %w", err)
	}

	encode := base64.StdEncoding.EncodeToString(data)

	if err = x.client.Do(
		ctx,
		x.client.B().Set().Key(
			NewKey(reminder.GetID())).
			Value(encode).
			Build(),
	).Error(); err != nil {
		return fmt.Errorf("failed to save reminder: %w", err)
	}

	return nil
}

func (x Repository) DeleteReminder(ctx context.Context, reminderID reminder_id_model.ReminderID) error {
	if err := x.client.Do(
		ctx,
		x.client.B().Del().
			Key(NewKey(reminderID)).
			Build(),
	).Error(); err != nil {
		return fmt.Errorf("failed to delete reminder: %w", err)
	}

	return nil
}

func (x Repository) GetReminderByID(
	ctx context.Context, reminderID reminder_id_model.ReminderID,
) (reminder_aggregate.Reminder, error) {
	v, err := x.client.Do(
		ctx,
		x.client.B().Get().
			Key(NewKey(reminderID)).
			Build(),
	).ToString()
	if err != nil {
		return reminder_aggregate.Reminder{}, fmt.Errorf("failed to get reminder: %w", err)
	}

	decode, err := base64.StdEncoding.DecodeString(v)
	if err != nil {
		return reminder_aggregate.Reminder{}, fmt.Errorf("failed to decode reminder: %w", err)
	}

	var reminderDTO reminder_v1.Reminder
	if err = proto.Unmarshal(decode, &reminderDTO); err != nil {
		return reminder_aggregate.Reminder{}, fmt.Errorf("failed to unmarshal reminder: %w", err)
	}

	reminder, err := ToReminder(&reminderDTO)
	if err != nil {
		return reminder_aggregate.Reminder{}, fmt.Errorf("failed to create reminder: %w", err)
	}

	return reminder, nil
}

func (x Repository) GetAllReminders(ctx context.Context) ([]reminder_aggregate.Reminder, error) {
	return nil, nil
}

func NewKey(reminderID reminder_id_model.ReminderID) string {
	return keyPrefix + strconv.FormatInt(int64(reminderID), 10)
}

func NewRepository(client rueidis.Client) reminder_aggregate.ReminderRepository {
	return Repository{
		client: client,
	}
}
