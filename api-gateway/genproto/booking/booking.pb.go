// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: booking.proto

package booking

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BookingRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProviderId    string `protobuf:"bytes,3,opt,name=provider_id,json=providerId,proto3" json:"provider_id,omitempty"`
	ServiceId     string `protobuf:"bytes,4,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	Status        string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	ScheduledTime string `protobuf:"bytes,6,opt,name=scheduled_time,json=scheduledTime,proto3" json:"scheduled_time,omitempty"`
	Location      string `protobuf:"bytes,7,opt,name=location,proto3" json:"location,omitempty"`
	TotalPrice    int64  `protobuf:"varint,8,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
}

func (x *BookingRes) Reset() {
	*x = BookingRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingRes) ProtoMessage() {}

func (x *BookingRes) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingRes.ProtoReflect.Descriptor instead.
func (*BookingRes) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{0}
}

func (x *BookingRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BookingRes) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *BookingRes) GetProviderId() string {
	if x != nil {
		return x.ProviderId
	}
	return ""
}

func (x *BookingRes) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *BookingRes) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *BookingRes) GetScheduledTime() string {
	if x != nil {
		return x.ScheduledTime
	}
	return ""
}

func (x *BookingRes) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *BookingRes) GetTotalPrice() int64 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

type BookingCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProviderId    string `protobuf:"bytes,2,opt,name=provider_id,json=providerId,proto3" json:"provider_id,omitempty"`
	ServiceId     string `protobuf:"bytes,3,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	Status        string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	ScheduledTime string `protobuf:"bytes,5,opt,name=scheduled_time,json=scheduledTime,proto3" json:"scheduled_time,omitempty"`
	Location      string `protobuf:"bytes,6,opt,name=location,proto3" json:"location,omitempty"`
	TotalPrice    int64  `protobuf:"varint,7,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
}

func (x *BookingCreateReq) Reset() {
	*x = BookingCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingCreateReq) ProtoMessage() {}

func (x *BookingCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingCreateReq.ProtoReflect.Descriptor instead.
func (*BookingCreateReq) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{1}
}

func (x *BookingCreateReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *BookingCreateReq) GetProviderId() string {
	if x != nil {
		return x.ProviderId
	}
	return ""
}

func (x *BookingCreateReq) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *BookingCreateReq) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *BookingCreateReq) GetScheduledTime() string {
	if x != nil {
		return x.ScheduledTime
	}
	return ""
}

func (x *BookingCreateReq) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *BookingCreateReq) GetTotalPrice() int64 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

type BookingGetByIdRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Booking *BookingRes `protobuf:"bytes,1,opt,name=booking,proto3" json:"booking,omitempty"`
}

func (x *BookingGetByIdRes) Reset() {
	*x = BookingGetByIdRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingGetByIdRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingGetByIdRes) ProtoMessage() {}

func (x *BookingGetByIdRes) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingGetByIdRes.ProtoReflect.Descriptor instead.
func (*BookingGetByIdRes) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{2}
}

func (x *BookingGetByIdRes) GetBooking() *BookingRes {
	if x != nil {
		return x.Booking
	}
	return nil
}

type BookingGetAllRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bookings []*BookingRes `protobuf:"bytes,1,rep,name=bookings,proto3" json:"bookings,omitempty"`
}

func (x *BookingGetAllRes) Reset() {
	*x = BookingGetAllRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingGetAllRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingGetAllRes) ProtoMessage() {}

func (x *BookingGetAllRes) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingGetAllRes.ProtoReflect.Descriptor instead.
func (*BookingGetAllRes) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{3}
}

func (x *BookingGetAllRes) GetBookings() []*BookingRes {
	if x != nil {
		return x.Bookings
	}
	return nil
}

type BookingUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *BookingUpdateReq) Reset() {
	*x = BookingUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingUpdateReq) ProtoMessage() {}

func (x *BookingUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingUpdateReq.ProtoReflect.Descriptor instead.
func (*BookingUpdateReq) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{4}
}

func (x *BookingUpdateReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BookingUpdateReq) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type BookingGetAllReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string  `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProviderId string  `protobuf:"bytes,2,opt,name=provider_id,json=providerId,proto3" json:"provider_id,omitempty"`
	Filter     *Filter `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *BookingGetAllReq) Reset() {
	*x = BookingGetAllReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingGetAllReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingGetAllReq) ProtoMessage() {}

func (x *BookingGetAllReq) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingGetAllReq.ProtoReflect.Descriptor instead.
func (*BookingGetAllReq) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{5}
}

func (x *BookingGetAllReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *BookingGetAllReq) GetProviderId() string {
	if x != nil {
		return x.ProviderId
	}
	return ""
}

func (x *BookingGetAllReq) GetFilter() *Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

var File_booking_proto protoreflect.FileDescriptor

var file_booking_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x1a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf1, 0x01, 0x0a, 0x0a, 0x42, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0xe7, 0x01, 0x0a, 0x10,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x42, 0x0a, 0x11, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x07, 0x62, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x52, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x22, 0x43, 0x0a, 0x10, 0x42, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x12, 0x2f, 0x0a,
	0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x3a,
	0x0a, 0x10, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x75, 0x0a, 0x10, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x32, 0x90, 0x02, 0x0a, 0x0e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x13,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x1a, 0x0d, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f,
	0x69, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0d, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x1a, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x12, 0x3e, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x12, 0x19, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x19, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x0d, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x56, 0x6f, 0x69, 0x64, 0x42, 0x12, 0x5a, 0x10, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_booking_proto_rawDescOnce sync.Once
	file_booking_proto_rawDescData = file_booking_proto_rawDesc
)

func file_booking_proto_rawDescGZIP() []byte {
	file_booking_proto_rawDescOnce.Do(func() {
		file_booking_proto_rawDescData = protoimpl.X.CompressGZIP(file_booking_proto_rawDescData)
	})
	return file_booking_proto_rawDescData
}

var file_booking_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_booking_proto_goTypes = []interface{}{
	(*BookingRes)(nil),        // 0: booking.BookingRes
	(*BookingCreateReq)(nil),  // 1: booking.BookingCreateReq
	(*BookingGetByIdRes)(nil), // 2: booking.BookingGetByIdRes
	(*BookingGetAllRes)(nil),  // 3: booking.BookingGetAllRes
	(*BookingUpdateReq)(nil),  // 4: booking.BookingUpdateReq
	(*BookingGetAllReq)(nil),  // 5: booking.BookingGetAllReq
	(*Filter)(nil),            // 6: booking.Filter
	(*ById)(nil),              // 7: booking.ById
	(*Void)(nil),              // 8: booking.Void
}
var file_booking_proto_depIdxs = []int32{
	0, // 0: booking.BookingGetByIdRes.booking:type_name -> booking.BookingRes
	0, // 1: booking.BookingGetAllRes.bookings:type_name -> booking.BookingRes
	6, // 2: booking.BookingGetAllReq.filter:type_name -> booking.Filter
	0, // 3: booking.BookingService.Create:input_type -> booking.BookingRes
	7, // 4: booking.BookingService.GetById:input_type -> booking.ById
	5, // 5: booking.BookingService.GetAll:input_type -> booking.BookingGetAllReq
	4, // 6: booking.BookingService.Update:input_type -> booking.BookingUpdateReq
	7, // 7: booking.BookingService.Delete:input_type -> booking.ById
	8, // 8: booking.BookingService.Create:output_type -> booking.Void
	2, // 9: booking.BookingService.GetById:output_type -> booking.BookingGetByIdRes
	3, // 10: booking.BookingService.GetAll:output_type -> booking.BookingGetAllRes
	8, // 11: booking.BookingService.Update:output_type -> booking.Void
	8, // 12: booking.BookingService.Delete:output_type -> booking.Void
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_booking_proto_init() }
func file_booking_proto_init() {
	if File_booking_proto != nil {
		return
	}
	file_common2_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_booking_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_booking_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingCreateReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_booking_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingGetByIdRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_booking_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingGetAllRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_booking_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingUpdateReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_booking_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingGetAllReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_booking_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_booking_proto_goTypes,
		DependencyIndexes: file_booking_proto_depIdxs,
		MessageInfos:      file_booking_proto_msgTypes,
	}.Build()
	File_booking_proto = out.File
	file_booking_proto_rawDesc = nil
	file_booking_proto_goTypes = nil
	file_booking_proto_depIdxs = nil
}
