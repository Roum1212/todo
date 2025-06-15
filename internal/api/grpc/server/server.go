package server_grpc

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	create_reminder_rpc "github.com/Roum1212/todo/internal/api/grpc/rpc/create-reminder"
	delete_reminder_rpc "github.com/Roum1212/todo/internal/api/grpc/rpc/delete-reminder"
	get_all_reminders_rpc "github.com/Roum1212/todo/internal/api/grpc/rpc/get-all-reminders"
	get_reminder_by_id_rpc "github.com/Roum1212/todo/internal/api/grpc/rpc/get-reminder-by-id"
	reminder_v1 "github.com/Roum1212/todo/pkg/gen/reminder/v1"
)

type ReminderServiceServer struct {
	reminder_v1.UnimplementedReminderServiceServer

	createReminderRPC  create_reminder_rpc.CreateReminderRPC
	deleteReminderRPC  delete_reminder_rpc.DeleteReminderRPC
	getAllRemindersRPC get_all_reminders_rpc.GetAllRemindersRPC
	getReminderByIDRPC get_reminder_by_id_rpc.GetReminderByIDRPC
}

func (x ReminderServiceServer) CreateReminder(
	ctx context.Context,
	r *reminder_v1.CreateReminderRequest,
) (*reminder_v1.CreateReminderResponse, error) {
	return x.createReminderRPC.CreateReminder(ctx, r)
}

func (x ReminderServiceServer) DeleteReminder(
	ctx context.Context,
	r *reminder_v1.DeleteReminderRequest,
) (*emptypb.Empty, error) {
	return x.deleteReminderRPC.DeleteReminder(ctx, r)
}

func (x ReminderServiceServer) GetAllReminders(
	ctx context.Context,
	empty *emptypb.Empty,
) (*reminder_v1.GetAllRemindersResponse, error) {
	return x.getAllRemindersRPC.GetAllReminders(ctx, empty)
}

func (x ReminderServiceServer) GetReminderByID(
	ctx context.Context,
	r *reminder_v1.GetReminderByIDRequest,
) (*reminder_v1.GetReminderByIDResponse, error) {
	return x.getReminderByIDRPC.GetReminderByID(ctx, r)
}

func NewCreateReminderService(
	createReminderRPC create_reminder_rpc.CreateReminderRPC,
	deleteReminderRPC delete_reminder_rpc.DeleteReminderRPC,
	getAllRemindersRPC get_all_reminders_rpc.GetAllRemindersRPC,
	getReminderByIDRPC get_reminder_by_id_rpc.GetReminderByIDRPC,
) ReminderServiceServer {
	return ReminderServiceServer{
		createReminderRPC:  createReminderRPC,
		deleteReminderRPC:  deleteReminderRPC,
		getAllRemindersRPC: getAllRemindersRPC,
		getReminderByIDRPC: getReminderByIDRPC,
	}
}
