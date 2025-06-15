package get_all_reminders_rpc

import (
	"crypto/rand"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	get_all_reminders_query "github.com/Roum1212/todo/internal/app/query/get-all-reminders"
	get_all_reminders_query_mock "github.com/Roum1212/todo/internal/app/query/get-all-reminders/mock"
	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
	reminder_description_model "github.com/Roum1212/todo/internal/domain/model/reminder-description"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
	reminder_title_model "github.com/Roum1212/todo/internal/domain/model/reminder-title"
)

func TestGetAllRemindersPRC_GetAllReminders(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	reminderTitleFirst, err := reminder_title_model.NewReminderTitle(rand.Text())
	require.NoError(t, err)

	reminderDescriptionFirst, err := reminder_description_model.NewReminderDescription(rand.Text())
	require.NoError(t, err)

	reminderFirst := reminder_aggregate.NewReminder(
		reminder_id_model.GenerateReminderID(),
		reminderTitleFirst,
		reminderDescriptionFirst,
	)

	reminderTitleSecond, err := reminder_title_model.NewReminderTitle(rand.Text())
	require.NoError(t, err)

	reminderDescriptionSecond, err := reminder_description_model.NewReminderDescription(rand.Text())
	require.NoError(t, err)

	reminderSecond := reminder_aggregate.NewReminder(
		reminder_id_model.GenerateReminderID(),
		reminderTitleSecond,
		reminderDescriptionSecond,
	)

	reminders := []reminder_aggregate.Reminder{reminderFirst, reminderSecond}

	handleQueryMock := get_all_reminders_query_mock.NewQueryHandlerMock(mc).
		HandleQueryMock.
		Expect(minimock.AnyContext).
		Return(reminders, nil)

	getAllRemindersRPC := NewGetAllRemindersRPC(handleQueryMock)

	getAllRemindersResponse, err := getAllRemindersRPC.GetAllReminders(t.Context(), &emptypb.Empty{})
	require.NoError(t, err)
	require.Equal(t, newReminderDTOs(reminders), getAllRemindersResponse.GetReminders())
}

func TestGetAllRemindersRPC_GetAllReminders_ErrReminderNotFound(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	handleQueryMock := get_all_reminders_query_mock.NewQueryHandlerMock(mc).
		HandleQueryMock.
		Expect(minimock.AnyContext).
		Return(nil, get_all_reminders_query.ErrReminderNotFound)

	getAllRemindersRPC := NewGetAllRemindersRPC(handleQueryMock)

	getAllRemindersResponse, err := getAllRemindersRPC.GetAllReminders(t.Context(), &emptypb.Empty{})
	require.Error(t, err)
	require.Nil(t, getAllRemindersResponse)

	pbStatus, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.NotFound, pbStatus.Code())
}

func TestGetAllRemindersRPC_GetAllReminders_Internal(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	handleQueryMock := get_all_reminders_query_mock.NewQueryHandlerMock(mc).
		HandleQueryMock.
		Expect(minimock.AnyContext).
		Return(nil, assert.AnError)

	getAllRemindersRPC := NewGetAllRemindersRPC(handleQueryMock)

	getAllRemindersResponse, err := getAllRemindersRPC.GetAllReminders(t.Context(), &emptypb.Empty{})
	require.Error(t, err)
	require.Nil(t, getAllRemindersResponse)

	pbStatus, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.Internal, pbStatus.Code())
}
