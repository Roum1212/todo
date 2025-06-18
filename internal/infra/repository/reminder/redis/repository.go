package redis_reminder_repository

import (
	"context"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"github.com/redis/rueidis"
	"google.golang.org/protobuf/proto"

	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
	reminder_v1 "github.com/Roum1212/todo/pkg/gen/reminder/v1"
)

const cursorCount = 100

const keyPrefix = "reminder:"

type Repository struct {
	client rueidis.Client
}

func (x Repository) DeleteReminder(ctx context.Context, reminderID reminder_id_model.ReminderID) error {
	if err := x.client.Do(
		ctx,
		x.client.B().Del().
			Key(newKey(reminderID)).
			Build(),
	).Error(); err != nil {
		return fmt.Errorf("failed to delete reminder: %w", err)
	}

	return nil
}

func (x Repository) GetAllReminders(ctx context.Context) ([]reminder_aggregate.Reminder, error) {
	var reminders []reminder_aggregate.Reminder

	var cursor uint64 = 0

	for {
		resp := x.client.Do(
			ctx,
			x.client.B().Scan().Cursor(cursor).Count(cursorCount).Build(),
		)

		scanRes, err := resp.AsScanEntry()
		if err != nil {
			return nil, fmt.Errorf("failed to scan reminder: %w", err)
		}

		cursor = scanRes.Cursor
		keys := scanRes.Elements

		var reminder reminder_aggregate.Reminder
		var reminderID reminder_id_model.ReminderID

		for _, key := range keys {
			reminderID, err = reminder_id_model.NewReminderIDFromString(strings.TrimPrefix(key, keyPrefix))
			if err != nil {
				return nil, fmt.Errorf("failed to parse reminder: %w", err)
			}

			reminder, err = x.GetReminderByID(ctx, reminderID)
			if err != nil {
				return nil, err
			}

			reminders = append(reminders, reminder)
		}

		if cursor == 0 {
			break
		}
	}

	return reminders, nil
}

func (x Repository) GetReminderByID(
	ctx context.Context, reminderID reminder_id_model.ReminderID,
) (reminder_aggregate.Reminder, error) {
	v, err := x.client.Do(
		ctx,
		x.client.B().Get().
			Key(newKey(reminderID)).
			Build(),
	).ToString()
	if err != nil {
		switch {
		case rueidis.IsRedisNil(err):
			return reminder_aggregate.Reminder{}, reminder_aggregate.ErrReminderNotFound
		default:
			return reminder_aggregate.Reminder{}, fmt.Errorf("failed to get reminder: %w", err)
		}
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
			newKey(reminder.GetID())).
			Value(encode).
			Build(),
	).Error(); err != nil {
		return fmt.Errorf("failed to save reminder: %w", err)
	}

	return nil
}

func NewRepository(client rueidis.Client) reminder_aggregate.ReminderRepository {
	return Repository{
		client: client,
	}
}

func newKey(reminderID reminder_id_model.ReminderID) string {
	return keyPrefix + strconv.FormatInt(int64(reminderID), 10)
}
