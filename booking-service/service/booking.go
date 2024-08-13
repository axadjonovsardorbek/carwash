package service

import (
	bp "booking/genproto/booking"
	st "booking/storage/postgres"
	"context"
)

type BookingService struct {
	storage st.Storage
	bp.UnimplementedBookingServiceServer
}

func NewBookingService(storage *st.Storage) *BookingService {
	return &BookingService{storage: *storage}
}

func (s *BookingService) Create(ctx context.Context, req *bp.BookingRes) (*bp.Void, error) {
	return s.storage.BookingS.Create(req)
}
func (s *BookingService) GetById(ctx context.Context, req *bp.ById) (*bp.BookingGetByIdRes, error) {
	return s.storage.BookingS.GetById(req)
}
func (s *BookingService) GetAll(ctx context.Context, req *bp.BookingGetAllReq) (*bp.BookingGetAllRes, error) {
	return s.storage.BookingS.GetAll(req)
}
func (s *BookingService) Update(ctx context.Context, req *bp.BookingUpdateReq) (*bp.Void, error) {
	return s.storage.BookingS.Update(req)
}
func (s *BookingService) Delete(ctx context.Context, req *bp.ById) (*bp.Void, error) {
	return s.storage.BookingS.Delete(req)
}
