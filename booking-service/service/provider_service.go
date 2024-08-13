package service

import (
	bp "booking/genproto/booking"
	st "booking/storage/postgres"
	"context"
)

type ProviderServiceService struct {
	storage st.Storage
	bp.UnimplementedProviderServiceServiceServer
}

func NewProviderServiceService(storage *st.Storage) *ProviderServiceService {
	return &ProviderServiceService{storage: *storage}
}

func (s *ProviderServiceService) Create(ctx context.Context, req *bp.ProviderServiceRes) (*bp.Void, error) {
	return s.storage.ProviderServiceS.Create(req)
}
func (s *ProviderServiceService) GetById(ctx context.Context, req *bp.ById) (*bp.ProviderServiceGetByIdRes, error) {
	return s.storage.ProviderServiceS.GetById(req)
}
func (s *ProviderServiceService) GetAll(ctx context.Context, req *bp.ProviderServiceGetAllReq) (*bp.ProviderServiceGetAllRes, error) {
	return s.storage.ProviderServiceS.GetAll(req)
}
func (s *ProviderServiceService) Delete(ctx context.Context, req *bp.ById) (*bp.Void, error) {
	return s.storage.ProviderServiceS.Delete(req)
}
