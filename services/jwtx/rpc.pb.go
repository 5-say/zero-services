// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: rpc.proto

package jwtx

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

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - 生成 token（登录）
// - 请求
type MakeToken_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessSecret    string `protobuf:"bytes,1,opt,name=accessSecret,proto3" json:"accessSecret,omitempty"`       // 密钥
	AccessExpire    int64  `protobuf:"varint,2,opt,name=accessExpire,proto3" json:"accessExpire,omitempty"`      // token 过期时间（秒）
	Group           string `protobuf:"bytes,3,opt,name=group,proto3" json:"group,omitempty"`                     // token 分组
	RandomAccountID string `protobuf:"bytes,4,opt,name=randomAccountID,proto3" json:"randomAccountID,omitempty"` // 加密的账户 ID
	RequestIP       string `protobuf:"bytes,5,opt,name=requestIP,proto3" json:"requestIP,omitempty"`             // 发起请求的 IP 地址
}

func (x *MakeToken_Request) Reset() {
	*x = MakeToken_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MakeToken_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MakeToken_Request) ProtoMessage() {}

func (x *MakeToken_Request) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MakeToken_Request.ProtoReflect.Descriptor instead.
func (*MakeToken_Request) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *MakeToken_Request) GetAccessSecret() string {
	if x != nil {
		return x.AccessSecret
	}
	return ""
}

func (x *MakeToken_Request) GetAccessExpire() int64 {
	if x != nil {
		return x.AccessExpire
	}
	return 0
}

func (x *MakeToken_Request) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *MakeToken_Request) GetRandomAccountID() string {
	if x != nil {
		return x.RandomAccountID
	}
	return ""
}

func (x *MakeToken_Request) GetRequestIP() string {
	if x != nil {
		return x.RequestIP
	}
	return ""
}

// - 响应
type MakeToken_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *MakeToken_Response) Reset() {
	*x = MakeToken_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MakeToken_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MakeToken_Response) ProtoMessage() {}

func (x *MakeToken_Response) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MakeToken_Response.ProtoReflect.Descriptor instead.
func (*MakeToken_Response) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *MakeToken_Response) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - 校验 token（拓展校验、刷新 token）
// - 请求
type CheckToken_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessSecret    string `protobuf:"bytes,1,opt,name=accessSecret,proto3" json:"accessSecret,omitempty"`        // 密钥
	RefreshInterval int64  `protobuf:"varint,2,opt,name=refreshInterval,proto3" json:"refreshInterval,omitempty"` // 刷新时间间隔（秒）
	FaultTolerance  int64  `protobuf:"varint,3,opt,name=faultTolerance,proto3" json:"faultTolerance,omitempty"`   // 并发容错时间（秒）
	CheckIP         bool   `protobuf:"varint,4,opt,name=checkIP,proto3" json:"checkIP,omitempty"`                 // 是否开启 IP 一致性校验，当前 IP 必须与登录时 IP 一致
	RequestIP       string `protobuf:"bytes,5,opt,name=requestIP,proto3" json:"requestIP,omitempty"`              // 发起请求的 IP 地址
	Tid             uint64 `protobuf:"varint,6,opt,name=tid,proto3" json:"tid,omitempty"`                         // 从 token 中解析出的 token ID
	Iat             int64  `protobuf:"varint,7,opt,name=iat,proto3" json:"iat,omitempty"`                         // 从 token 中解析出的 签发时间
	Exp             int64  `protobuf:"varint,8,opt,name=exp,proto3" json:"exp,omitempty"`                         // 从 token 中解析出的 过期时间
}

func (x *CheckToken_Request) Reset() {
	*x = CheckToken_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckToken_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckToken_Request) ProtoMessage() {}

func (x *CheckToken_Request) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckToken_Request.ProtoReflect.Descriptor instead.
func (*CheckToken_Request) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *CheckToken_Request) GetAccessSecret() string {
	if x != nil {
		return x.AccessSecret
	}
	return ""
}

func (x *CheckToken_Request) GetRefreshInterval() int64 {
	if x != nil {
		return x.RefreshInterval
	}
	return 0
}

func (x *CheckToken_Request) GetFaultTolerance() int64 {
	if x != nil {
		return x.FaultTolerance
	}
	return 0
}

func (x *CheckToken_Request) GetCheckIP() bool {
	if x != nil {
		return x.CheckIP
	}
	return false
}

func (x *CheckToken_Request) GetRequestIP() string {
	if x != nil {
		return x.RequestIP
	}
	return ""
}

func (x *CheckToken_Request) GetTid() uint64 {
	if x != nil {
		return x.Tid
	}
	return 0
}

func (x *CheckToken_Request) GetIat() int64 {
	if x != nil {
		return x.Iat
	}
	return 0
}

func (x *CheckToken_Request) GetExp() int64 {
	if x != nil {
		return x.Exp
	}
	return 0
}

// - 响应
type CheckToken_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewToken        string `protobuf:"bytes,1,opt,name=newToken,proto3" json:"newToken,omitempty"`
	RandomAccountID string `protobuf:"bytes,2,opt,name=randomAccountID,proto3" json:"randomAccountID,omitempty"` // 加密的账户 ID
}

func (x *CheckToken_Response) Reset() {
	*x = CheckToken_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckToken_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckToken_Response) ProtoMessage() {}

func (x *CheckToken_Response) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckToken_Response.ProtoReflect.Descriptor instead.
func (*CheckToken_Response) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *CheckToken_Response) GetNewToken() string {
	if x != nil {
		return x.NewToken
	}
	return ""
}

func (x *CheckToken_Response) GetRandomAccountID() string {
	if x != nil {
		return x.RandomAccountID
	}
	return ""
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - 移除 token（安全退出）
// - 请求
type DeleteToken_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tid uint64 `protobuf:"varint,1,opt,name=tid,proto3" json:"tid,omitempty"` // token ID
}

func (x *DeleteToken_Request) Reset() {
	*x = DeleteToken_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteToken_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteToken_Request) ProtoMessage() {}

func (x *DeleteToken_Request) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteToken_Request.ProtoReflect.Descriptor instead.
func (*DeleteToken_Request) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteToken_Request) GetTid() uint64 {
	if x != nil {
		return x.Tid
	}
	return 0
}

// - 响应
type DeleteToken_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteToken_Response) Reset() {
	*x = DeleteToken_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteToken_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteToken_Response) ProtoMessage() {}

func (x *DeleteToken_Response) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteToken_Response.ProtoReflect.Descriptor instead.
func (*DeleteToken_Response) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{5}
}

var File_rpc_proto protoreflect.FileDescriptor

var file_rpc_proto_rawDesc = []byte{
	0x0a, 0x09, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6a, 0x77, 0x74,
	0x78, 0x22, 0xb9, 0x01, 0x0a, 0x11, 0x4d, 0x61, 0x6b, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x5f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12,
	0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x50, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x50, 0x22, 0x2a, 0x0a,
	0x12, 0x4d, 0x61, 0x6b, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xf8, 0x01, 0x0a, 0x12, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x22, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x26,
	0x0a, 0x0e, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x54, 0x6f, 0x6c, 0x65, 0x72, 0x61, 0x6e, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x54, 0x6f, 0x6c,
	0x65, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49,
	0x50, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x50,
	0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x50, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x50, 0x12, 0x10,
	0x0a, 0x03, 0x74, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x74, 0x69, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x69, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x69,
	0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x65, 0x78, 0x70, 0x22, 0x5b, 0x0a, 0x13, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e,
	0x65, 0x77, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e,
	0x65, 0x77, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x61, 0x6e, 0x64, 0x6f,
	0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x44, 0x22, 0x27, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x74, 0x69, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0xcf, 0x01, 0x0a, 0x04, 0x6a, 0x77, 0x74, 0x78, 0x12, 0x3e, 0x0a, 0x09, 0x4d,
	0x61, 0x6b, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x17, 0x2e, 0x6a, 0x77, 0x74, 0x78, 0x2e,
	0x4d, 0x61, 0x6b, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x6a, 0x77, 0x74, 0x78, 0x2e, 0x4d, 0x61, 0x6b, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x2e, 0x6a, 0x77, 0x74, 0x78,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6a, 0x77, 0x74, 0x78, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44,
	0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x2e,
	0x6a, 0x77, 0x74, 0x78, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6a, 0x77, 0x74, 0x78, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x6a, 0x77, 0x74, 0x78, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_proto_rawDescOnce sync.Once
	file_rpc_proto_rawDescData = file_rpc_proto_rawDesc
)

func file_rpc_proto_rawDescGZIP() []byte {
	file_rpc_proto_rawDescOnce.Do(func() {
		file_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_proto_rawDescData)
	})
	return file_rpc_proto_rawDescData
}

var file_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_rpc_proto_goTypes = []interface{}{
	(*MakeToken_Request)(nil),    // 0: jwtx.MakeToken_Request
	(*MakeToken_Response)(nil),   // 1: jwtx.MakeToken_Response
	(*CheckToken_Request)(nil),   // 2: jwtx.CheckToken_Request
	(*CheckToken_Response)(nil),  // 3: jwtx.CheckToken_Response
	(*DeleteToken_Request)(nil),  // 4: jwtx.DeleteToken_Request
	(*DeleteToken_Response)(nil), // 5: jwtx.DeleteToken_Response
}
var file_rpc_proto_depIdxs = []int32{
	0, // 0: jwtx.jwtx.MakeToken:input_type -> jwtx.MakeToken_Request
	2, // 1: jwtx.jwtx.CheckToken:input_type -> jwtx.CheckToken_Request
	4, // 2: jwtx.jwtx.DeleteToken:input_type -> jwtx.DeleteToken_Request
	1, // 3: jwtx.jwtx.MakeToken:output_type -> jwtx.MakeToken_Response
	3, // 4: jwtx.jwtx.CheckToken:output_type -> jwtx.CheckToken_Response
	5, // 5: jwtx.jwtx.DeleteToken:output_type -> jwtx.DeleteToken_Response
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_proto_init() }
func file_rpc_proto_init() {
	if File_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MakeToken_Request); i {
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
		file_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MakeToken_Response); i {
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
		file_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckToken_Request); i {
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
		file_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckToken_Response); i {
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
		file_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteToken_Request); i {
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
		file_rpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteToken_Response); i {
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
			RawDescriptor: file_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_proto_goTypes,
		DependencyIndexes: file_rpc_proto_depIdxs,
		MessageInfos:      file_rpc_proto_msgTypes,
	}.Build()
	File_rpc_proto = out.File
	file_rpc_proto_rawDesc = nil
	file_rpc_proto_goTypes = nil
	file_rpc_proto_depIdxs = nil
}
