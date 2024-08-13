package service

import (
	bp "booking/genproto/booking"
	st "booking/storage/postgres"
	"context"
)

type ServiceService struct {
	storage st.Storage
	bp.UnimplementedServiceServiceServer
}

func NewServiceService(storage *st.Storage) *ServiceService {
	return &ServiceService{storage: *storage}
}

func (s *ServiceService) Create(ctx context.Context, req *bp.ServiceRes) (*bp.Void, error) {
	return s.storage.ServiceS.Create(req)
}
func (s *ServiceService) GetById(ctx context.Context, req *bp.ById) (*bp.ServiceGetByIdRes, error) {
	return s.storage.ServiceS.GetById(req)
}
func (s *ServiceService) GetAll(ctx context.Context, req *bp.ServiceGetAllReq) (*bp.ServiceGetAllRes, error) {
	return s.storage.ServiceS.GetAll(req)
}
func (s *ServiceService) Update(ctx context.Context, req *bp.ServiceUpdateReq) (*bp.Void, error) {
	return s.storage.ServiceS.Update(req)
}
func (s *ServiceService) Delete(ctx context.Context, req *bp.ById) (*bp.Void, error) {
	return s.storage.ServiceS.Delete(req)
}
