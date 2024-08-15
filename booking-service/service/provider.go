package service

import (
	bp "booking/genproto/booking"
	st "booking/storage/postgres"
	"context"
)

type ProviderService struct {
	storage st.Storage
	bp.UnimplementedProviderServiceServer
}

func NewProviderService(storage *st.Storage) *ProviderService {
	return &ProviderService{storage: *storage}
}

func (s *ProviderService) Create(ctx context.Context, req *bp.ProviderRes) (*bp.Void, error) {
	return s.storage.ProviderS.Create(req)
}
func (s *ProviderService) GetById(ctx context.Context, req *bp.ById) (*bp.ProviderGetByIdRes, error) {
	return s.storage.ProviderS.GetById(req)
}
func (s *ProviderService) GetAll(ctx context.Context, req *bp.ProviderGetAllReq) (*bp.ProviderGetAllRes, error) {
	return s.storage.ProviderS.GetAll(req)
}
func (s *ProviderService) Update(ctx context.Context, req *bp.ProviderUpdateReq) (*bp.Void, error) {
	return s.storage.ProviderS.Update(req)
}
func (s *ProviderService) Delete(ctx context.Context, req *bp.ById) (*bp.Void, error) {
	return s.storage.ProviderS.Delete(req)
}

func (s *ProviderService) GetProviderId(ctx context.Context, req *bp.ById) (*bp.ById, error) {
	return s.storage.ProviderS.GetProviderId(req)
}
