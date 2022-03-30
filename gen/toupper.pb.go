// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: toupper.proto

package gen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// The request message
type UpperRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpperRequest) Reset() {
	*x = UpperRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_toupper_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpperRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpperRequest) ProtoMessage() {}

func (x *UpperRequest) ProtoReflect() protoreflect.Message {
	mi := &file_toupper_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpperRequest.ProtoReflect.Descriptor instead.
func (*UpperRequest) Descriptor() ([]byte, []int) {
	return file_toupper_proto_rawDescGZIP(), []int{0}
}

func (x *UpperRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The response message
type UpperResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UpperResponse) Reset() {
	*x = UpperResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_toupper_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpperResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpperResponse) ProtoMessage() {}

func (x *UpperResponse) ProtoReflect() protoreflect.Message {
	mi := &file_toupper_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpperResponse.ProtoReflect.Descriptor instead.
func (*UpperResponse) Descriptor() ([]byte, []int) {
	return file_toupper_proto_rawDescGZIP(), []int{1}
}

func (x *UpperResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// query user info request
type QueryUserInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Base *Base  `protobuf:"bytes,2,opt,name=base,proto3" json:"base,omitempty"`
}

func (x *QueryUserInfoReq) Reset() {
	*x = QueryUserInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_toupper_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryUserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryUserInfoReq) ProtoMessage() {}

func (x *QueryUserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_toupper_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryUserInfoReq.ProtoReflect.Descriptor instead.
func (*QueryUserInfoReq) Descriptor() ([]byte, []int) {
	return file_toupper_proto_rawDescGZIP(), []int{2}
}

func (x *QueryUserInfoReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *QueryUserInfoReq) GetBase() *Base {
	if x != nil {
		return x.Base
	}
	return nil
}

type QueryUserInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInfo *UserInfo `protobuf:"bytes,1,opt,name=userInfo,proto3" json:"userInfo,omitempty"`
	BaseResp *BaseResp `protobuf:"bytes,2,opt,name=baseResp,proto3" json:"baseResp,omitempty"`
}

func (x *QueryUserInfoResp) Reset() {
	*x = QueryUserInfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_toupper_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryUserInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryUserInfoResp) ProtoMessage() {}

func (x *QueryUserInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_toupper_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryUserInfoResp.ProtoReflect.Descriptor instead.
func (*QueryUserInfoResp) Descriptor() ([]byte, []int) {
	return file_toupper_proto_rawDescGZIP(), []int{3}
}

func (x *QueryUserInfoResp) GetUserInfo() *UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

func (x *QueryUserInfoResp) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Gender int32  `protobuf:"varint,3,opt,name=gender,proto3" json:"gender,omitempty"`
	Hobby  string `protobuf:"bytes,4,opt,name=hobby,proto3" json:"hobby,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_toupper_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_toupper_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_toupper_proto_rawDescGZIP(), []int{4}
}

func (x *UserInfo) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UserInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserInfo) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *UserInfo) GetHobby() string {
	if x != nil {
		return x.Hobby
	}
	return ""
}

type Base struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Caller string `protobuf:"bytes,1,opt,name=caller,proto3" json:"caller,omitempty"`
	LogID  string `protobuf:"bytes,2,opt,name=logID,proto3" json:"logID,omitempty"`
}

func (x *Base) Reset() {
	*x = Base{}
	if protoimpl.UnsafeEnabled {
		mi := &file_toupper_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Base) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Base) ProtoMessage() {}

func (x *Base) ProtoReflect() protoreflect.Message {
	mi := &file_toupper_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Base.ProtoReflect.Descriptor instead.
func (*Base) Descriptor() ([]byte, []int) {
	return file_toupper_proto_rawDescGZIP(), []int{5}
}

func (x *Base) GetCaller() string {
	if x != nil {
		return x.Caller
	}
	return ""
}

func (x *Base) GetLogID() string {
	if x != nil {
		return x.LogID
	}
	return ""
}

type BaseResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusMessage string `protobuf:"bytes,1,opt,name=status_message,json=statusMessage,proto3" json:"status_message,omitempty"`
	StatusCode    int32  `protobuf:"varint,2,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
}

func (x *BaseResp) Reset() {
	*x = BaseResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_toupper_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResp) ProtoMessage() {}

func (x *BaseResp) ProtoReflect() protoreflect.Message {
	mi := &file_toupper_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResp.ProtoReflect.Descriptor instead.
func (*BaseResp) Descriptor() ([]byte, []int) {
	return file_toupper_proto_rawDescGZIP(), []int{6}
}

func (x *BaseResp) GetStatusMessage() string {
	if x != nil {
		return x.StatusMessage
	}
	return ""
}

func (x *BaseResp) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

var File_toupper_proto protoreflect.FileDescriptor

var file_toupper_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x74, 0x6f, 0x75, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x0c, 0x55, 0x70, 0x70, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a, 0x0d, 0x55, 0x70,
	0x70, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x47, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a,
	0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x22, 0x6d,
	0x0a, 0x11, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x2b, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x2b, 0x0a, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x5c, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x6f, 0x62, 0x62, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x68, 0x6f, 0x62, 0x62, 0x79, 0x22, 0x34, 0x0a, 0x04, 0x42,
	0x61, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x6f, 0x67, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x49,
	0x44, 0x22, 0x52, 0x0a, 0x08, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x25, 0x0a,
	0x0e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x85, 0x01, 0x0a, 0x07, 0x54, 0x6f, 0x55, 0x70, 0x70, 0x65,
	0x72, 0x12, 0x34, 0x0a, 0x05, 0x55, 0x70, 0x70, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x55, 0x70, 0x70, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x70, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0d, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x06, 0x5a,
	0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_toupper_proto_rawDescOnce sync.Once
	file_toupper_proto_rawDescData = file_toupper_proto_rawDesc
)

func file_toupper_proto_rawDescGZIP() []byte {
	file_toupper_proto_rawDescOnce.Do(func() {
		file_toupper_proto_rawDescData = protoimpl.X.CompressGZIP(file_toupper_proto_rawDescData)
	})
	return file_toupper_proto_rawDescData
}

var file_toupper_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_toupper_proto_goTypes = []interface{}{
	(*UpperRequest)(nil),      // 0: proto.UpperRequest
	(*UpperResponse)(nil),     // 1: proto.UpperResponse
	(*QueryUserInfoReq)(nil),  // 2: proto.QueryUserInfoReq
	(*QueryUserInfoResp)(nil), // 3: proto.QueryUserInfoResp
	(*UserInfo)(nil),          // 4: proto.UserInfo
	(*Base)(nil),              // 5: proto.Base
	(*BaseResp)(nil),          // 6: proto.BaseResp
}
var file_toupper_proto_depIdxs = []int32{
	5, // 0: proto.QueryUserInfoReq.base:type_name -> proto.Base
	4, // 1: proto.QueryUserInfoResp.userInfo:type_name -> proto.UserInfo
	6, // 2: proto.QueryUserInfoResp.baseResp:type_name -> proto.BaseResp
	0, // 3: proto.ToUpper.Upper:input_type -> proto.UpperRequest
	2, // 4: proto.ToUpper.QueryUserInfo:input_type -> proto.QueryUserInfoReq
	1, // 5: proto.ToUpper.Upper:output_type -> proto.UpperResponse
	3, // 6: proto.ToUpper.QueryUserInfo:output_type -> proto.QueryUserInfoResp
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_toupper_proto_init() }
func file_toupper_proto_init() {
	if File_toupper_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_toupper_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpperRequest); i {
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
		file_toupper_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpperResponse); i {
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
		file_toupper_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryUserInfoReq); i {
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
		file_toupper_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryUserInfoResp); i {
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
		file_toupper_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_toupper_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Base); i {
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
		file_toupper_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseResp); i {
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
			RawDescriptor: file_toupper_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_toupper_proto_goTypes,
		DependencyIndexes: file_toupper_proto_depIdxs,
		MessageInfos:      file_toupper_proto_msgTypes,
	}.Build()
	File_toupper_proto = out.File
	file_toupper_proto_rawDesc = nil
	file_toupper_proto_goTypes = nil
	file_toupper_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ToUpperClient is the client API for ToUpper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ToUpperClient interface {
	// Sends a greeting
	Upper(ctx context.Context, in *UpperRequest, opts ...grpc.CallOption) (*UpperResponse, error)
	QueryUserInfo(ctx context.Context, in *QueryUserInfoReq, opts ...grpc.CallOption) (*QueryUserInfoResp, error)
}

type toUpperClient struct {
	cc grpc.ClientConnInterface
}

func NewToUpperClient(cc grpc.ClientConnInterface) ToUpperClient {
	return &toUpperClient{cc}
}

func (c *toUpperClient) Upper(ctx context.Context, in *UpperRequest, opts ...grpc.CallOption) (*UpperResponse, error) {
	out := new(UpperResponse)
	err := c.cc.Invoke(ctx, "/proto.ToUpper/Upper", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toUpperClient) QueryUserInfo(ctx context.Context, in *QueryUserInfoReq, opts ...grpc.CallOption) (*QueryUserInfoResp, error) {
	out := new(QueryUserInfoResp)
	err := c.cc.Invoke(ctx, "/proto.ToUpper/QueryUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ToUpperServer is the server API for ToUpper service.
type ToUpperServer interface {
	// Sends a greeting
	Upper(context.Context, *UpperRequest) (*UpperResponse, error)
	QueryUserInfo(context.Context, *QueryUserInfoReq) (*QueryUserInfoResp, error)
}

// UnimplementedToUpperServer can be embedded to have forward compatible implementations.
type UnimplementedToUpperServer struct {
}

func (*UnimplementedToUpperServer) Upper(context.Context, *UpperRequest) (*UpperResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upper not implemented")
}
func (*UnimplementedToUpperServer) QueryUserInfo(context.Context, *QueryUserInfoReq) (*QueryUserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryUserInfo not implemented")
}

func RegisterToUpperServer(s *grpc.Server, srv ToUpperServer) {
	s.RegisterService(&_ToUpper_serviceDesc, srv)
}

func _ToUpper_Upper_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpperRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToUpperServer).Upper(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ToUpper/Upper",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToUpperServer).Upper(ctx, req.(*UpperRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToUpper_QueryUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToUpperServer).QueryUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ToUpper/QueryUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToUpperServer).QueryUserInfo(ctx, req.(*QueryUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ToUpper_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ToUpper",
	HandlerType: (*ToUpperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Upper",
			Handler:    _ToUpper_Upper_Handler,
		},
		{
			MethodName: "QueryUserInfo",
			Handler:    _ToUpper_QueryUserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "toupper.proto",
}