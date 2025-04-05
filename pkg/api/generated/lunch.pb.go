// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.28.3
// source: lunch.proto

package tlunch

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Surname       string                 `protobuf:"bytes,3,opt,name=surname,proto3" json:"surname,omitempty"`
	Tg            string                 `protobuf:"bytes,4,opt,name=tg,proto3" json:"tg,omitempty"`
	Office        string                 `protobuf:"bytes,5,opt,name=office,proto3" json:"office,omitempty"`
	Emoji         string                 `protobuf:"bytes,6,opt,name=emoji,proto3" json:"emoji,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_lunch_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_lunch_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_lunch_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *User) GetTg() string {
	if x != nil {
		return x.Tg
	}
	return ""
}

func (x *User) GetOffice() string {
	if x != nil {
		return x.Office
	}
	return ""
}

func (x *User) GetEmoji() string {
	if x != nil {
		return x.Emoji
	}
	return ""
}

type Lunch struct {
	state                protoimpl.MessageState `protogen:"open.v1"`
	Id                   int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Surname              string                 `protobuf:"bytes,3,opt,name=surname,proto3" json:"surname,omitempty"`
	Place                string                 `protobuf:"bytes,4,opt,name=place,proto3" json:"place,omitempty"`
	Time                 *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=time,proto3" json:"time,omitempty"`
	NumberOfParticipants int64                  `protobuf:"varint,6,opt,name=number_of_participants,json=numberOfParticipants,proto3" json:"number_of_participants,omitempty"`
	Description          *string                `protobuf:"bytes,7,opt,name=description,proto3,oneof" json:"description,omitempty"`
	Users                []*User                `protobuf:"bytes,8,rep,name=users,proto3" json:"users,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *Lunch) Reset() {
	*x = Lunch{}
	mi := &file_lunch_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Lunch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lunch) ProtoMessage() {}

func (x *Lunch) ProtoReflect() protoreflect.Message {
	mi := &file_lunch_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lunch.ProtoReflect.Descriptor instead.
func (*Lunch) Descriptor() ([]byte, []int) {
	return file_lunch_proto_rawDescGZIP(), []int{1}
}

func (x *Lunch) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Lunch) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Lunch) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *Lunch) GetPlace() string {
	if x != nil {
		return x.Place
	}
	return ""
}

func (x *Lunch) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Lunch) GetNumberOfParticipants() int64 {
	if x != nil {
		return x.NumberOfParticipants
	}
	return 0
}

func (x *Lunch) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *Lunch) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type LunchFeedback struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Lunch         *Lunch                 `protobuf:"bytes,1,opt,name=lunch,proto3" json:"lunch,omitempty"`
	IsLiked       bool                   `protobuf:"varint,2,opt,name=is_liked,json=isLiked,proto3" json:"is_liked,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LunchFeedback) Reset() {
	*x = LunchFeedback{}
	mi := &file_lunch_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LunchFeedback) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LunchFeedback) ProtoMessage() {}

func (x *LunchFeedback) ProtoReflect() protoreflect.Message {
	mi := &file_lunch_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LunchFeedback.ProtoReflect.Descriptor instead.
func (*LunchFeedback) Descriptor() ([]byte, []int) {
	return file_lunch_proto_rawDescGZIP(), []int{2}
}

func (x *LunchFeedback) GetLunch() *Lunch {
	if x != nil {
		return x.Lunch
	}
	return nil
}

func (x *LunchFeedback) GetIsLiked() bool {
	if x != nil {
		return x.IsLiked
	}
	return false
}

type UserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	mi := &file_lunch_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lunch_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_lunch_proto_rawDescGZIP(), []int{3}
}

func (x *UserRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type LunchRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Offset        int32                  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit         int32                  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LunchRequest) Reset() {
	*x = LunchRequest{}
	mi := &file_lunch_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LunchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LunchRequest) ProtoMessage() {}

func (x *LunchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lunch_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LunchRequest.ProtoReflect.Descriptor instead.
func (*LunchRequest) Descriptor() ([]byte, []int) {
	return file_lunch_proto_rawDescGZIP(), []int{4}
}

func (x *LunchRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *LunchRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *LunchRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetLunchesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Lunches       []*Lunch               `protobuf:"bytes,1,rep,name=lunches,proto3" json:"lunches,omitempty"`
	LunchId       *int64                 `protobuf:"varint,2,opt,name=lunch_id,json=lunchId,proto3,oneof" json:"lunch_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetLunchesResponse) Reset() {
	*x = GetLunchesResponse{}
	mi := &file_lunch_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLunchesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLunchesResponse) ProtoMessage() {}

func (x *GetLunchesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lunch_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLunchesResponse.ProtoReflect.Descriptor instead.
func (*GetLunchesResponse) Descriptor() ([]byte, []int) {
	return file_lunch_proto_rawDescGZIP(), []int{5}
}

func (x *GetLunchesResponse) GetLunches() []*Lunch {
	if x != nil {
		return x.Lunches
	}
	return nil
}

func (x *GetLunchesResponse) GetLunchId() int64 {
	if x != nil && x.LunchId != nil {
		return *x.LunchId
	}
	return 0
}

type CreateLunchRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Place         string                 `protobuf:"bytes,2,opt,name=place,proto3" json:"place,omitempty"`
	Time          *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	Description   string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateLunchRequest) Reset() {
	*x = CreateLunchRequest{}
	mi := &file_lunch_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateLunchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLunchRequest) ProtoMessage() {}

func (x *CreateLunchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lunch_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLunchRequest.ProtoReflect.Descriptor instead.
func (*CreateLunchRequest) Descriptor() ([]byte, []int) {
	return file_lunch_proto_rawDescGZIP(), []int{6}
}

func (x *CreateLunchRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateLunchRequest) GetPlace() string {
	if x != nil {
		return x.Place
	}
	return ""
}

func (x *CreateLunchRequest) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *CreateLunchRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type LunchResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Lunch         *Lunch                 `protobuf:"bytes,1,opt,name=lunch,proto3" json:"lunch,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LunchResponse) Reset() {
	*x = LunchResponse{}
	mi := &file_lunch_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LunchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LunchResponse) ProtoMessage() {}

func (x *LunchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lunch_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LunchResponse.ProtoReflect.Descriptor instead.
func (*LunchResponse) Descriptor() ([]byte, []int) {
	return file_lunch_proto_rawDescGZIP(), []int{7}
}

func (x *LunchResponse) GetLunch() *Lunch {
	if x != nil {
		return x.Lunch
	}
	return nil
}

type ActionLunchRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LunchId       int64                  `protobuf:"varint,1,opt,name=lunch_id,json=lunchId,proto3" json:"lunch_id,omitempty"`
	UserId        int64                  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ActionLunchRequest) Reset() {
	*x = ActionLunchRequest{}
	mi := &file_lunch_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ActionLunchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionLunchRequest) ProtoMessage() {}

func (x *ActionLunchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lunch_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionLunchRequest.ProtoReflect.Descriptor instead.
func (*ActionLunchRequest) Descriptor() ([]byte, []int) {
	return file_lunch_proto_rawDescGZIP(), []int{8}
}

func (x *ActionLunchRequest) GetLunchId() int64 {
	if x != nil {
		return x.LunchId
	}
	return 0
}

func (x *ActionLunchRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type DetailLunchRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LunchId       int64                  `protobuf:"varint,1,opt,name=lunch_id,json=lunchId,proto3" json:"lunch_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DetailLunchRequest) Reset() {
	*x = DetailLunchRequest{}
	mi := &file_lunch_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DetailLunchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailLunchRequest) ProtoMessage() {}

func (x *DetailLunchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lunch_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailLunchRequest.ProtoReflect.Descriptor instead.
func (*DetailLunchRequest) Descriptor() ([]byte, []int) {
	return file_lunch_proto_rawDescGZIP(), []int{9}
}

func (x *DetailLunchRequest) GetLunchId() int64 {
	if x != nil {
		return x.LunchId
	}
	return 0
}

type LunchHistoryResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Lunches       []*LunchFeedback       `protobuf:"bytes,1,rep,name=lunches,proto3" json:"lunches,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LunchHistoryResponse) Reset() {
	*x = LunchHistoryResponse{}
	mi := &file_lunch_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LunchHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LunchHistoryResponse) ProtoMessage() {}

func (x *LunchHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lunch_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LunchHistoryResponse.ProtoReflect.Descriptor instead.
func (*LunchHistoryResponse) Descriptor() ([]byte, []int) {
	return file_lunch_proto_rawDescGZIP(), []int{10}
}

func (x *LunchHistoryResponse) GetLunches() []*LunchFeedback {
	if x != nil {
		return x.Lunches
	}
	return nil
}

type RateLunchRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	LunchId       int64                  `protobuf:"varint,2,opt,name=lunch_id,json=lunchId,proto3" json:"lunch_id,omitempty"`
	IsLiked       bool                   `protobuf:"varint,3,opt,name=is_liked,json=isLiked,proto3" json:"is_liked,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RateLunchRequest) Reset() {
	*x = RateLunchRequest{}
	mi := &file_lunch_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RateLunchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLunchRequest) ProtoMessage() {}

func (x *RateLunchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lunch_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLunchRequest.ProtoReflect.Descriptor instead.
func (*RateLunchRequest) Descriptor() ([]byte, []int) {
	return file_lunch_proto_rawDescGZIP(), []int{11}
}

func (x *RateLunchRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RateLunchRequest) GetLunchId() int64 {
	if x != nil {
		return x.LunchId
	}
	return 0
}

func (x *RateLunchRequest) GetIsLiked() bool {
	if x != nil {
		return x.IsLiked
	}
	return false
}

var File_lunch_proto protoreflect.FileDescriptor

var file_lunch_proto_rawDesc = string([]byte{
	0x0a, 0x0b, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x74,
	0x6c, 0x75, 0x6e, 0x63, 0x68, 0x1a, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x67, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x69, 0x63,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x22, 0x9c, 0x02, 0x0a, 0x05, 0x4c, 0x75, 0x6e, 0x63, 0x68, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66,
	0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x14, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x50, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01,
	0x12, 0x22, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x74, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4f, 0x0a, 0x0d, 0x4c, 0x75, 0x6e, 0x63, 0x68, 0x46, 0x65, 0x65,
	0x64, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x23, 0x0a, 0x05, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x2e, 0x4c, 0x75,
	0x6e, 0x63, 0x68, 0x52, 0x05, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73,
	0x5f, 0x6c, 0x69, 0x6b, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73,
	0x4c, 0x69, 0x6b, 0x65, 0x64, 0x22, 0x26, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x55, 0x0a,
	0x0c, 0x4c, 0x75, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x22, 0x6a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4c, 0x75, 0x6e, 0x63, 0x68,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x07, 0x6c, 0x75,
	0x6e, 0x63, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x6c,
	0x75, 0x6e, 0x63, 0x68, 0x2e, 0x4c, 0x75, 0x6e, 0x63, 0x68, 0x52, 0x07, 0x6c, 0x75, 0x6e, 0x63,
	0x68, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x08, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x07, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x49, 0x64,
	0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64,
	0x22, 0x95, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x75, 0x6e, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x34, 0x0a, 0x0d, 0x4c, 0x75, 0x6e, 0x63,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x6c, 0x75, 0x6e,
	0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x6c, 0x75, 0x6e, 0x63,
	0x68, 0x2e, 0x4c, 0x75, 0x6e, 0x63, 0x68, 0x52, 0x05, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x22, 0x48,
	0x0a, 0x12, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x75, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x12, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x4c, 0x75, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x22, 0x47, 0x0a, 0x14, 0x4c, 0x75, 0x6e,
	0x63, 0x68, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2f, 0x0a, 0x07, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x2e, 0x4c, 0x75, 0x6e, 0x63,
	0x68, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x07, 0x6c, 0x75, 0x6e, 0x63, 0x68,
	0x65, 0x73, 0x22, 0x61, 0x0a, 0x10, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x75, 0x6e, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73,
	0x5f, 0x6c, 0x69, 0x6b, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73,
	0x4c, 0x69, 0x6b, 0x65, 0x64, 0x32, 0xe2, 0x06, 0x0a, 0x06, 0x54, 0x6c, 0x75, 0x6e, 0x63, 0x68,
	0x12, 0x4e, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x13,
	0x2e, 0x74, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x74, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d,
	0x12, 0x4d, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x12, 0x0c, 0x2e, 0x74, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a,
	0x0c, 0x2e, 0x74, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x20, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x32, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x12,
	0x53, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4c, 0x75, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x12, 0x14, 0x2e,
	0x74, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x2e, 0x4c, 0x75, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x2e, 0x47, 0x65, 0x74,
	0x4c, 0x75, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x75, 0x6e,
	0x63, 0x68, 0x65, 0x73, 0x12, 0x56, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x75,
	0x6e, 0x63, 0x68, 0x12, 0x1a, 0x2e, 0x74, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4c, 0x75, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x74, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x2e, 0x4c, 0x75, 0x6e, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01,
	0x2a, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x12, 0x64, 0x0a, 0x09,
	0x4a, 0x6f, 0x69, 0x6e, 0x4c, 0x75, 0x6e, 0x63, 0x68, 0x12, 0x1a, 0x2e, 0x74, 0x6c, 0x75, 0x6e,
	0x63, 0x68, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x75, 0x6e, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x74, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x2e, 0x4c,
	0x75, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x32, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x75, 0x6e,
	0x63, 0x68, 0x2f, 0x7b, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x6a, 0x6f,
	0x69, 0x6e, 0x12, 0x66, 0x0a, 0x0a, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x4c, 0x75, 0x6e, 0x63, 0x68,
	0x12, 0x1a, 0x2e, 0x74, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x75, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x74,
	0x6c, 0x75, 0x6e, 0x63, 0x68, 0x2e, 0x4c, 0x75, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x32, 0x1a,
	0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x2f, 0x7b, 0x6c, 0x75, 0x6e, 0x63, 0x68,
	0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x6c, 0x65, 0x61, 0x76, 0x65, 0x12, 0x61, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x4c, 0x75, 0x6e, 0x63, 0x68, 0x12, 0x1a, 0x2e, 0x74,
	0x6c, 0x75, 0x6e, 0x63, 0x68, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x4c, 0x75, 0x6e, 0x63,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x74, 0x6c, 0x75, 0x6e, 0x63,
	0x68, 0x2e, 0x4c, 0x75, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x75, 0x6e,
	0x63, 0x68, 0x2f, 0x7b, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x6a, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x4c, 0x75, 0x6e, 0x63, 0x68, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x12, 0x14, 0x2e, 0x74, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x2e, 0x4c, 0x75, 0x6e, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x74, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x2e,
	0x4c, 0x75, 0x6e, 0x63, 0x68, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x76,
	0x31, 0x2f, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x7d, 0x2f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x6f, 0x0a, 0x09, 0x52, 0x61, 0x74,
	0x65, 0x4c, 0x75, 0x6e, 0x63, 0x68, 0x12, 0x18, 0x2e, 0x74, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x2e,
	0x52, 0x61, 0x74, 0x65, 0x4c, 0x75, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x74, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x2e, 0x4c, 0x75, 0x6e, 0x63, 0x68, 0x46,
	0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x22, 0x31, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x3a,
	0x01, 0x2a, 0x32, 0x26, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x2f, 0x7b, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2f,
	0x7b, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x2a, 0x5a, 0x28, 0x74, 0x2d,
	0x6c, 0x75, 0x6e, 0x63, 0x68, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x3b,
	0x74, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_lunch_proto_rawDescOnce sync.Once
	file_lunch_proto_rawDescData []byte
)

func file_lunch_proto_rawDescGZIP() []byte {
	file_lunch_proto_rawDescOnce.Do(func() {
		file_lunch_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_lunch_proto_rawDesc), len(file_lunch_proto_rawDesc)))
	})
	return file_lunch_proto_rawDescData
}

var file_lunch_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_lunch_proto_goTypes = []any{
	(*User)(nil),                  // 0: tlunch.User
	(*Lunch)(nil),                 // 1: tlunch.Lunch
	(*LunchFeedback)(nil),         // 2: tlunch.LunchFeedback
	(*UserRequest)(nil),           // 3: tlunch.UserRequest
	(*LunchRequest)(nil),          // 4: tlunch.LunchRequest
	(*GetLunchesResponse)(nil),    // 5: tlunch.GetLunchesResponse
	(*CreateLunchRequest)(nil),    // 6: tlunch.CreateLunchRequest
	(*LunchResponse)(nil),         // 7: tlunch.LunchResponse
	(*ActionLunchRequest)(nil),    // 8: tlunch.ActionLunchRequest
	(*DetailLunchRequest)(nil),    // 9: tlunch.DetailLunchRequest
	(*LunchHistoryResponse)(nil),  // 10: tlunch.LunchHistoryResponse
	(*RateLunchRequest)(nil),      // 11: tlunch.RateLunchRequest
	(*timestamppb.Timestamp)(nil), // 12: google.protobuf.Timestamp
}
var file_lunch_proto_depIdxs = []int32{
	12, // 0: tlunch.Lunch.time:type_name -> google.protobuf.Timestamp
	0,  // 1: tlunch.Lunch.users:type_name -> tlunch.User
	1,  // 2: tlunch.LunchFeedback.lunch:type_name -> tlunch.Lunch
	1,  // 3: tlunch.GetLunchesResponse.lunches:type_name -> tlunch.Lunch
	12, // 4: tlunch.CreateLunchRequest.time:type_name -> google.protobuf.Timestamp
	1,  // 5: tlunch.LunchResponse.lunch:type_name -> tlunch.Lunch
	2,  // 6: tlunch.LunchHistoryResponse.lunches:type_name -> tlunch.LunchFeedback
	3,  // 7: tlunch.Tlunch.GetProfile:input_type -> tlunch.UserRequest
	0,  // 8: tlunch.Tlunch.ChangeProfile:input_type -> tlunch.User
	4,  // 9: tlunch.Tlunch.GetLunches:input_type -> tlunch.LunchRequest
	6,  // 10: tlunch.Tlunch.CreateLunch:input_type -> tlunch.CreateLunchRequest
	8,  // 11: tlunch.Tlunch.JoinLunch:input_type -> tlunch.ActionLunchRequest
	8,  // 12: tlunch.Tlunch.LeaveLunch:input_type -> tlunch.ActionLunchRequest
	9,  // 13: tlunch.Tlunch.GetDetailLunch:input_type -> tlunch.DetailLunchRequest
	4,  // 14: tlunch.Tlunch.GetLunchHistory:input_type -> tlunch.LunchRequest
	11, // 15: tlunch.Tlunch.RateLunch:input_type -> tlunch.RateLunchRequest
	0,  // 16: tlunch.Tlunch.GetProfile:output_type -> tlunch.User
	0,  // 17: tlunch.Tlunch.ChangeProfile:output_type -> tlunch.User
	5,  // 18: tlunch.Tlunch.GetLunches:output_type -> tlunch.GetLunchesResponse
	7,  // 19: tlunch.Tlunch.CreateLunch:output_type -> tlunch.LunchResponse
	7,  // 20: tlunch.Tlunch.JoinLunch:output_type -> tlunch.LunchResponse
	7,  // 21: tlunch.Tlunch.LeaveLunch:output_type -> tlunch.LunchResponse
	7,  // 22: tlunch.Tlunch.GetDetailLunch:output_type -> tlunch.LunchResponse
	10, // 23: tlunch.Tlunch.GetLunchHistory:output_type -> tlunch.LunchHistoryResponse
	2,  // 24: tlunch.Tlunch.RateLunch:output_type -> tlunch.LunchFeedback
	16, // [16:25] is the sub-list for method output_type
	7,  // [7:16] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_lunch_proto_init() }
func file_lunch_proto_init() {
	if File_lunch_proto != nil {
		return
	}
	file_lunch_proto_msgTypes[1].OneofWrappers = []any{}
	file_lunch_proto_msgTypes[5].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_lunch_proto_rawDesc), len(file_lunch_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lunch_proto_goTypes,
		DependencyIndexes: file_lunch_proto_depIdxs,
		MessageInfos:      file_lunch_proto_msgTypes,
	}.Build()
	File_lunch_proto = out.File
	file_lunch_proto_goTypes = nil
	file_lunch_proto_depIdxs = nil
}
