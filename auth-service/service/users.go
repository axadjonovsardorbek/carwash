package service

import (
	ap "auth/genproto/auth"
	st "auth/storage/postgres"
	"context"
)

type UsersService struct {
	storage st.Storage
	ap.UnimplementedUserServiceServer
}

func NewUsersService(storage *st.Storage) *UsersService {
	return &UsersService{storage: *storage}
}

func (u *UsersService) Register(ctx context.Context, req *ap.UserCreateReq) (*ap.Void, error) {
	return u.storage.UserS.Register(req)
}

func (u *UsersService) Login(ctx context.Context, req *ap.UserLoginReq) (*ap.TokenRes, error) {
	return u.storage.UserS.Login(req)
}

func (u *UsersService) Profile(ctx context.Context, req *ap.ById) (*ap.UserRes, error) {
	return u.storage.UserS.Profile(req)
}

func (u *UsersService) UpdateProfile(ctx context.Context, req *ap.UserUpdateReq) (*ap.Void, error) {
	return u.storage.UserS.UpdateProfile(req)
}

func (u *UsersService) DeleteProfile(ctx context.Context, req *ap.ById) (*ap.Void, error) {
	return u.storage.UserS.DeleteProfile(req)
}

func (u *UsersService) RefreshToken(ctx context.Context, req *ap.ById) (*ap.TokenRes, error) {
	return u.storage.UserS.RefreshToken(req)
}

func (u *UsersService) ForgotPassword(ctx context.Context, req *ap.UsersForgotPassword) (*ap.Void, error) {
	return u.storage.UserS.ForgotPassword(req)
}

func (u *UsersService) ResetPassword(ctx context.Context, req *ap.UsersResetPassword) (*ap.Void, error) {
	return u.storage.UserS.ResetPassword(req)
}

func (u *UsersService) ChangePassword(ctx context.Context, req *ap.UsersChangePassword) (*ap.Void, error) {
	return u.storage.UserS.ChangePassword(req)
}

func (u *UsersService) CheckEmail(ctx context.Context, req *ap.CheckEmailReq) (*ap.ById, error){
	return u.storage.UserS.CheckEmail(req)
}

func (u *UsersService) GetAllUsers(ctx context.Context, req *ap.GetAllUsersReq) (*ap.GetAllUsersRes, error){
	return u.storage.UserS.GetAllUsers(req)
}