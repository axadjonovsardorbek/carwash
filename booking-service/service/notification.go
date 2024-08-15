package service

import (
	bp "booking/genproto/booking"
	st "booking/storage/mongo"
	"context"
)

type NotificationService struct {
	storage st.Storage
	bp.UnimplementedNotificationServiceServer
}

func NewNotificationService(storage *st.Storage) *NotificationService {
	return &NotificationService{storage: *storage}
}

func (s *NotificationService) Create(ctx context.Context, req *bp.NotificationRes) (*bp.Void, error) {
	return s.storage.NotificationS.Create(req)
}
func (s *NotificationService) GetById(ctx context.Context, req *bp.ById) (*bp.NotificationGetByIdRes, error) {
	return s.storage.NotificationS.GetById(req)
}
func (s *NotificationService) GetAll(ctx context.Context, req *bp.NotificationGetAllReq) (*bp.NotificationGetAllRes, error) {
	return s.storage.NotificationS.GetAll(req)
}
func (s *NotificationService) Update(ctx context.Context, req *bp.NotificationUpdateReq) (*bp.Void, error) {
	return s.storage.NotificationS.Update(req)
}
