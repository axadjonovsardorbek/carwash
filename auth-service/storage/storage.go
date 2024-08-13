package storage

import (
	ap "auth/genproto/auth"
)

type UserI interface {
	Register(*ap.UserCreateReq) (*ap.Void, error)
	Login(*ap.UserLoginReq) (*ap.TokenRes, error)
	Profile(*ap.ById) (*ap.UserRes, error)
	UpdateProfile(*ap.UserUpdateReq) (*ap.Void, error)
	DeleteProfile(*ap.ById) (*ap.Void, error)
	RefreshToken(*ap.ById) (*ap.TokenRes, error)
	ForgotPassword(*ap.UsersForgotPassword) (*ap.Void, error)
	ResetPassword(*ap.UsersResetPassword) (*ap.Void, error)
	ChangePassword(*ap.UsersChangePassword) (*ap.Void, error)
	CheckEmail(*ap.CheckEmailReq) (*ap.ById, error)
	GetAllUsers(*ap.GetAllUsersReq) (*ap.GetAllUsersRes, error)
}
