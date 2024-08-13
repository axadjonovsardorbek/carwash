package service

import (
	bp "booking/genproto/booking"
	st "booking/storage/postgres"
	"context"
)

type ReviewService struct {
	storage st.Storage
	bp.UnimplementedReviewServiceServer
}

func NewReviewService(storage *st.Storage) *ReviewService {
	return &ReviewService{storage: *storage}
}

func (s *ReviewService) Create(ctx context.Context, req *bp.ReviewRes) (*bp.Void, error) {
	return s.storage.ReviewS.Create(req)
}
func (s *ReviewService) GetById(ctx context.Context, req *bp.ById) (*bp.ReviewGetByIdRes, error) {
	return s.storage.ReviewS.GetById(req)
}
func (s *ReviewService) GetAll(ctx context.Context, req *bp.ReviewGetAllReq) (*bp.ReviewGetAllRes, error) {
	return s.storage.ReviewS.GetAll(req)
}
func (s *ReviewService) Update(ctx context.Context, req *bp.ReviewUpdateReq) (*bp.Void, error) {
	return s.storage.ReviewS.Update(req)
}
func (s *ReviewService) Delete(ctx context.Context, req *bp.ById) (*bp.Void, error) {
	return s.storage.ReviewS.Delete(req)
}
