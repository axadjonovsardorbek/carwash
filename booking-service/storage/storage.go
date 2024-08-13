package storage

import (
	bp "booking/genproto/booking"
)

type ProviderI interface {
	Create(*bp.ProviderRes) (*bp.Void, error)
	GetById(*bp.ById) (*bp.ProviderGetByIdRes, error)
	GetAll(*bp.ProviderGetAllReq) (*bp.ProviderGetAllRes, error)
	Update(*bp.ProviderUpdateReq) (*bp.Void, error)
	Delete(*bp.ById) (*bp.Void, error)
}

type ServiceI interface {
	Create(*bp.ServiceRes) (*bp.Void, error)
	GetById(*bp.ById) (*bp.ServiceGetByIdRes, error)
	GetAll(*bp.ServiceGetAllReq) (*bp.ServiceGetAllRes, error)
	Update(*bp.ServiceUpdateReq) (*bp.Void, error)
	Delete(*bp.ById) (*bp.Void, error)
}

type BookingI interface {
	Create(*bp.BookingRes) (*bp.Void, error)
	GetById(*bp.ById) (*bp.BookingGetByIdRes, error)
	GetAll(*bp.BookingGetAllReq) (*bp.BookingGetAllRes, error)
	Update(*bp.BookingUpdateReq) (*bp.Void, error)
	Delete(*bp.ById) (*bp.Void, error)
}

type PaymentI interface {
	Create(*bp.PaymentRes) (*bp.Void, error)
	GetById(*bp.ById) (*bp.PaymentGetByIdRes, error)
	GetAll(*bp.PaymentGetAllReq) (*bp.PaymentGetAllRes, error)
	Update(*bp.PaymentUpdateReq) (*bp.Void, error)
	Delete(*bp.ById) (*bp.Void, error)
}

type ReviewI interface {
	Create(*bp.ReviewRes) (*bp.Void, error)
	GetById(*bp.ById) (*bp.ReviewGetByIdRes, error)
	GetAll(*bp.ReviewGetAllReq) (*bp.ReviewGetAllRes, error)
	Update(*bp.ReviewUpdateReq) (*bp.Void, error)
	Delete(*bp.ById) (*bp.Void, error)
}

type ProviderServiceI interface {
	Create(*bp.ProviderServiceRes) (*bp.Void, error)
	GetById(*bp.ById) (*bp.ProviderServiceGetByIdRes, error)
	GetAll(*bp.ProviderServiceGetAllReq) (*bp.ProviderServiceGetAllRes, error)
	Delete(*bp.ById) (*bp.Void, error)
}
