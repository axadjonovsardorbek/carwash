package service

import (
	bp "booking/genproto/booking"
	st "booking/storage/postgres"
	"context"
)

type PaymentService struct {
	storage st.Storage
	bp.UnimplementedPaymentServiceServer
}

func NewPaymentService(storage *st.Storage) *PaymentService {
	return &PaymentService{storage: *storage}
}

func (s *PaymentService) Create(ctx context.Context, req *bp.PaymentRes) (*bp.Void, error) {
	return s.storage.PaymentS.Create(req)
}
func (s *PaymentService) GetById(ctx context.Context, req *bp.ById) (*bp.PaymentGetByIdRes, error) {
	return s.storage.PaymentS.GetById(req)
}
func (s *PaymentService) GetAll(ctx context.Context, req *bp.PaymentGetAllReq) (*bp.PaymentGetAllRes, error) {
	return s.storage.PaymentS.GetAll(req)
}
func (s *PaymentService) Update(ctx context.Context, req *bp.PaymentUpdateReq) (*bp.Void, error) {
	return s.storage.PaymentS.Update(req)
}
func (s *PaymentService) Delete(ctx context.Context, req *bp.ById) (*bp.Void, error) {
	return s.storage.PaymentS.Delete(req)
}
func (s *PaymentService) GetBookingId(ctx context.Context, req *bp.ById) (*bp.ById, error){
	return s.storage.PaymentS.GetBookingId(req)
}
func (s *PaymentService) GetBookingAmount(ctx context.Context, req *bp.ById) (*bp.GetAmountRes, error){
	return s.storage.PaymentS.GetBookingAmount(req)
}
