// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/user/v1/user_repo.proto

package v1

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

type AddUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *AddUserReq) Reset() {
	*x = AddUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_v1_user_repo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserReq) ProtoMessage() {}

func (x *AddUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_v1_user_repo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserReq.ProtoReflect.Descriptor instead.
func (*AddUserReq) Descriptor() ([]byte, []int) {
	return file_api_user_v1_user_repo_proto_rawDescGZIP(), []int{0}
}

func (x *AddUserReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddUserReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AddUserReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AddUserRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddUserRes) Reset() {
	*x = AddUserRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_v1_user_repo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserRes) ProtoMessage() {}

func (x *AddUserRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_v1_user_repo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserRes.ProtoReflect.Descriptor instead.
func (*AddUserRes) Descriptor() ([]byte, []int) {
	return file_api_user_v1_user_repo_proto_rawDescGZIP(), []int{1}
}

func (x *AddUserRes) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UserIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserIdReq) Reset() {
	*x = UserIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_v1_user_repo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserIdReq) ProtoMessage() {}

func (x *UserIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_v1_user_repo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserIdReq.ProtoReflect.Descriptor instead.
func (*UserIdReq) Descriptor() ([]byte, []int) {
	return file_api_user_v1_user_repo_proto_rawDescGZIP(), []int{2}
}

func (x *UserIdReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UserRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Avatar   string `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Password string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	Salt     string `protobuf:"bytes,6,opt,name=salt,proto3" json:"salt,omitempty"`
}

func (x *UserRes) Reset() {
	*x = UserRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_v1_user_repo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRes) ProtoMessage() {}

func (x *UserRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_v1_user_repo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRes.ProtoReflect.Descriptor instead.
func (*UserRes) Descriptor() ([]byte, []int) {
	return file_api_user_v1_user_repo_proto_rawDescGZIP(), []int{3}
}

func (x *UserRes) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserRes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserRes) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserRes) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *UserRes) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserRes) GetSalt() string {
	if x != nil {
		return x.Salt
	}
	return ""
}

type UserListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *UserListReq) Reset() {
	*x = UserListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_v1_user_repo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListReq) ProtoMessage() {}

func (x *UserListReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_v1_user_repo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListReq.ProtoReflect.Descriptor instead.
func (*UserListReq) Descriptor() ([]byte, []int) {
	return file_api_user_v1_user_repo_proto_rawDescGZIP(), []int{4}
}

func (x *UserListReq) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type UserListRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserList []*UserRes `protobuf:"bytes,1,rep,name=userList,proto3" json:"userList,omitempty"`
}

func (x *UserListRes) Reset() {
	*x = UserListRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_v1_user_repo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserListRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListRes) ProtoMessage() {}

func (x *UserListRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_v1_user_repo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListRes.ProtoReflect.Descriptor instead.
func (*UserListRes) Descriptor() ([]byte, []int) {
	return file_api_user_v1_user_repo_proto_rawDescGZIP(), []int{5}
}

func (x *UserListRes) GetUserList() []*UserRes {
	if x != nil {
		return x.UserList
	}
	return nil
}

type ConditionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *ConditionReq) Reset() {
	*x = ConditionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_v1_user_repo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConditionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConditionReq) ProtoMessage() {}

func (x *ConditionReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_v1_user_repo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConditionReq.ProtoReflect.Descriptor instead.
func (*ConditionReq) Descriptor() ([]byte, []int) {
	return file_api_user_v1_user_repo_proto_rawDescGZIP(), []int{6}
}

func (x *ConditionReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_api_user_v1_user_repo_proto protoreflect.FileDescriptor

var file_api_user_v1_user_repo_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76,
	0x31, 0x22, 0x52, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x1c, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x1b, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x8b, 0x01, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x61,
	0x6c, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x61, 0x6c, 0x74, 0x22, 0x1f,
	0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a,
	0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22,
	0x36, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x12, 0x27,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x24, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x32, 0xce, 0x01,
	0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x2b, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x2b,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0d, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0f, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a,
	0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x06,
	0x5a, 0x04, 0x2e, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_user_v1_user_repo_proto_rawDescOnce sync.Once
	file_api_user_v1_user_repo_proto_rawDescData = file_api_user_v1_user_repo_proto_rawDesc
)

func file_api_user_v1_user_repo_proto_rawDescGZIP() []byte {
	file_api_user_v1_user_repo_proto_rawDescOnce.Do(func() {
		file_api_user_v1_user_repo_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_user_v1_user_repo_proto_rawDescData)
	})
	return file_api_user_v1_user_repo_proto_rawDescData
}

var file_api_user_v1_user_repo_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_user_v1_user_repo_proto_goTypes = []interface{}{
	(*AddUserReq)(nil),   // 0: v1.AddUserReq
	(*AddUserRes)(nil),   // 1: v1.AddUserRes
	(*UserIdReq)(nil),    // 2: v1.UserIdReq
	(*UserRes)(nil),      // 3: v1.UserRes
	(*UserListReq)(nil),  // 4: v1.UserListReq
	(*UserListRes)(nil),  // 5: v1.UserListRes
	(*ConditionReq)(nil), // 6: v1.ConditionReq
}
var file_api_user_v1_user_repo_proto_depIdxs = []int32{
	3, // 0: v1.UserListRes.userList:type_name -> v1.UserRes
	0, // 1: v1.UserRepoService.AddUser:input_type -> v1.AddUserReq
	2, // 2: v1.UserRepoService.GetUserById:input_type -> v1.UserIdReq
	4, // 3: v1.UserRepoService.UserList:input_type -> v1.UserListReq
	6, // 4: v1.UserRepoService.GetByCondition:input_type -> v1.ConditionReq
	1, // 5: v1.UserRepoService.AddUser:output_type -> v1.AddUserRes
	3, // 6: v1.UserRepoService.GetUserById:output_type -> v1.UserRes
	5, // 7: v1.UserRepoService.UserList:output_type -> v1.UserListRes
	3, // 8: v1.UserRepoService.GetByCondition:output_type -> v1.UserRes
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_user_v1_user_repo_proto_init() }
func file_api_user_v1_user_repo_proto_init() {
	if File_api_user_v1_user_repo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_user_v1_user_repo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserReq); i {
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
		file_api_user_v1_user_repo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserRes); i {
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
		file_api_user_v1_user_repo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserIdReq); i {
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
		file_api_user_v1_user_repo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRes); i {
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
		file_api_user_v1_user_repo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserListReq); i {
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
		file_api_user_v1_user_repo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserListRes); i {
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
		file_api_user_v1_user_repo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConditionReq); i {
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
			RawDescriptor: file_api_user_v1_user_repo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_user_v1_user_repo_proto_goTypes,
		DependencyIndexes: file_api_user_v1_user_repo_proto_depIdxs,
		MessageInfos:      file_api_user_v1_user_repo_proto_msgTypes,
	}.Build()
	File_api_user_v1_user_repo_proto = out.File
	file_api_user_v1_user_repo_proto_rawDesc = nil
	file_api_user_v1_user_repo_proto_goTypes = nil
	file_api_user_v1_user_repo_proto_depIdxs = nil
}
