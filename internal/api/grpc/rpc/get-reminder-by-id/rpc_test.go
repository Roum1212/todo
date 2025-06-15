package get_reminder_by_id_rpc

import (
	"crypto/rand"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	get_reminder_by_id_query "github.com/Roum1212/todo/internal/app/query/get-reminder-by-id"
	get_reminder_by_id_query_mock "github.com/Roum1212/todo/internal/app/query/get-reminder-by-id/mock"
	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
	reminder_description_model "github.com/Roum1212/todo/internal/domain/model/reminder-description"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
	reminder_title_model "github.com/Roum1212/todo/internal/domain/model/reminder-title"
	reminder_v1 "github.com/Roum1212/todo/pkg/gen/reminder/v1"
)

func TestGetReminderByIDRPC_GetReminderByID(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	reminderID := reminder_id_model.GenerateReminderID()

	request := reminder_v1.GetReminderByIDRequest{
		Id: int64(reminderID),
	}

	reminderTitle, err := reminder_title_model.NewReminderTitle(rand.Text())
	require.NoError(t, err)

	reminderDescription, err := reminder_description_model.NewReminderDescription(rand.Text())
	require.NoError(t, err)

	reminder := reminder_aggregate.NewReminder(reminderID, reminderTitle, reminderDescription)

	handleQueryMock := get_reminder_by_id_query_mock.NewQueryHandlerMock(mc).
		HandleQueryMock.
		Expect(minimock.AnyContext, get_reminder_by_id_query.NewQuery(reminderID)).
		Return(reminder, nil)

	getReminderByIDPRC := NewGetReminderByIDRPC(handleQueryMock)

	getReminderByIDResponse, err := getReminderByIDPRC.GetReminderByID(t.Context(), &request)
	require.NoError(t, err)
	require.Equal(t, newReminderDTO(reminder), getReminderByIDResponse.GetReminder())
}

func TestGetReminderByIDRPC_GetReminderByID_ErrReminderNotFound(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	reminderID := reminder_id_model.GenerateReminderID()

	request := reminder_v1.GetReminderByIDRequest{
		Id: int64(reminderID),
	}

	handleQueryMock := get_reminder_by_id_query_mock.NewQueryHandlerMock(mc).
		HandleQueryMock.
		Expect(minimock.AnyContext, get_reminder_by_id_query.NewQuery(reminderID)).
		Return(reminder_aggregate.Reminder{}, get_reminder_by_id_query.ErrReminderNotFound)

	getReminderByIDRPC := NewGetReminderByIDRPC(handleQueryMock)

	getReminderByIDResponse, err := getReminderByIDRPC.GetReminderByID(t.Context(), &request)
	require.Error(t, err)
	require.Nil(t, getReminderByIDResponse)

	pbStatus, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.NotFound, pbStatus.Code())
}

func TestGetReminderByIDRPC_GetReminderByID_Internal(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	reminderID := reminder_id_model.GenerateReminderID()

	request := reminder_v1.GetReminderByIDRequest{
		Id: int64(reminderID),
	}

	handleQueryMock := get_reminder_by_id_query_mock.NewQueryHandlerMock(mc).
		HandleQueryMock.
		Expect(minimock.AnyContext, get_reminder_by_id_query.NewQuery(reminderID)).
		Return(reminder_aggregate.Reminder{}, assert.AnError)

	getReminderByIDRPC := NewGetReminderByIDRPC(handleQueryMock)

	getReminderByIDResponse, err := getReminderByIDRPC.GetReminderByID(t.Context(), &request)
	require.Error(t, err)
	require.Nil(t, getReminderByIDResponse)

	pbStatus, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.Internal, pbStatus.Code())
}

func TestGetReminderByIDRPC_GetReminderBy_InvalidArgument(t *testing.T) {
	t.Parallel()

	getReminderByIDRPC := NewGetReminderByIDRPC(nil)

	getReminderByIDResponse, err := getReminderByIDRPC.GetReminderByID(t.Context(), &reminder_v1.GetReminderByIDRequest{})
	require.Error(t, err)
	require.Nil(t, getReminderByIDResponse)
}
