// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: define/rpc.proto

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

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - 生成 token - 请求
type MakeToken_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecretKey string `protobuf:"bytes,1,opt,name=secretKey,proto3" json:"secretKey,omitempty"`  // 密钥
	ExpSecond int64  `protobuf:"varint,2,opt,name=expSecond,proto3" json:"expSecond,omitempty"` // 过期时长（秒）
	Group     string `protobuf:"bytes,3,opt,name=group,proto3" json:"group,omitempty"`          // 分组
	AccountID uint64 `protobuf:"varint,4,opt,name=accountID,proto3" json:"accountID,omitempty"` // 账户 ID
}

func (x *MakeToken_Request) Reset() {
	*x = MakeToken_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_define_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MakeToken_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MakeToken_Request) ProtoMessage() {}

func (x *MakeToken_Request) ProtoReflect() protoreflect.Message {
	mi := &file_define_rpc_proto_msgTypes[0]
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
	return file_define_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *MakeToken_Request) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

func (x *MakeToken_Request) GetExpSecond() int64 {
	if x != nil {
		return x.ExpSecond
	}
	return 0
}

func (x *MakeToken_Request) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *MakeToken_Request) GetAccountID() uint64 {
	if x != nil {
		return x.AccountID
	}
	return 0
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - 生成 token - 响应
type MakeToken_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *MakeToken_Response) Reset() {
	*x = MakeToken_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_define_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MakeToken_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MakeToken_Response) ProtoMessage() {}

func (x *MakeToken_Response) ProtoReflect() protoreflect.Message {
	mi := &file_define_rpc_proto_msgTypes[1]
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
	return file_define_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *MakeToken_Response) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - 校验 token - 请求
type CheckToken_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecretKey string `protobuf:"bytes,1,opt,name=secretKey,proto3" json:"secretKey,omitempty"` // 密钥
	Grp       string `protobuf:"bytes,2,opt,name=grp,proto3" json:"grp,omitempty"`             // 分组
	Rai       string `protobuf:"bytes,3,opt,name=rai,proto3" json:"rai,omitempty"`             // 加密的账户 ID
	Iat       string `protobuf:"bytes,4,opt,name=iat,proto3" json:"iat,omitempty"`             // 分发时间
	Exp       string `protobuf:"bytes,5,opt,name=exp,proto3" json:"exp,omitempty"`             // 过期时间
}

func (x *CheckToken_Request) Reset() {
	*x = CheckToken_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_define_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckToken_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckToken_Request) ProtoMessage() {}

func (x *CheckToken_Request) ProtoReflect() protoreflect.Message {
	mi := &file_define_rpc_proto_msgTypes[2]
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
	return file_define_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *CheckToken_Request) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

func (x *CheckToken_Request) GetGrp() string {
	if x != nil {
		return x.Grp
	}
	return ""
}

func (x *CheckToken_Request) GetRai() string {
	if x != nil {
		return x.Rai
	}
	return ""
}

func (x *CheckToken_Request) GetIat() string {
	if x != nil {
		return x.Iat
	}
	return ""
}

func (x *CheckToken_Request) GetExp() string {
	if x != nil {
		return x.Exp
	}
	return ""
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - 校验 token - 响应
type CheckToken_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewToken  string `protobuf:"bytes,1,opt,name=newToken,proto3" json:"newToken,omitempty"`
	AccountID uint64 `protobuf:"varint,2,opt,name=accountID,proto3" json:"accountID,omitempty"` // 账户 ID
}

func (x *CheckToken_Response) Reset() {
	*x = CheckToken_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_define_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckToken_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckToken_Response) ProtoMessage() {}

func (x *CheckToken_Response) ProtoReflect() protoreflect.Message {
	mi := &file_define_rpc_proto_msgTypes[3]
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
	return file_define_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *CheckToken_Response) GetNewToken() string {
	if x != nil {
		return x.NewToken
	}
	return ""
}

func (x *CheckToken_Response) GetAccountID() uint64 {
	if x != nil {
		return x.AccountID
	}
	return 0
}

var File_define_rpc_proto protoreflect.FileDescriptor

var file_define_rpc_proto_rawDesc = []byte{
	0x0a, 0x10, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x6a, 0x77, 0x74, 0x78, 0x22, 0x83, 0x01, 0x0a, 0x11, 0x4d, 0x61, 0x6b,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09,
	0x65, 0x78, 0x70, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x65, 0x78, 0x70, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x2a,
	0x0a, 0x12, 0x4d, 0x61, 0x6b, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x7a, 0x0a, 0x12, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x67, 0x72, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x67, 0x72, 0x70,
	0x12, 0x10, 0x0a, 0x03, 0x72, 0x61, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72,
	0x61, 0x69, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x69, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x65, 0x78, 0x70, 0x22, 0x4f, 0x0a, 0x13, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6e, 0x65, 0x77, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6e, 0x65, 0x77, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x32, 0x89, 0x01, 0x0a, 0x04, 0x6a, 0x77, 0x74, 0x78,
	0x12, 0x3e, 0x0a, 0x09, 0x4d, 0x61, 0x6b, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x17, 0x2e,
	0x6a, 0x77, 0x74, 0x78, 0x2e, 0x4d, 0x61, 0x6b, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6a, 0x77, 0x74, 0x78, 0x2e, 0x4d, 0x61,
	0x6b, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x41, 0x0a, 0x0a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18,
	0x2e, 0x6a, 0x77, 0x74, 0x78, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6a, 0x77, 0x74, 0x78, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x6a, 0x77, 0x74, 0x78, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_define_rpc_proto_rawDescOnce sync.Once
	file_define_rpc_proto_rawDescData = file_define_rpc_proto_rawDesc
)

func file_define_rpc_proto_rawDescGZIP() []byte {
	file_define_rpc_proto_rawDescOnce.Do(func() {
		file_define_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_define_rpc_proto_rawDescData)
	})
	return file_define_rpc_proto_rawDescData
}

var file_define_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_define_rpc_proto_goTypes = []interface{}{
	(*MakeToken_Request)(nil),   // 0: jwtx.MakeToken_Request
	(*MakeToken_Response)(nil),  // 1: jwtx.MakeToken_Response
	(*CheckToken_Request)(nil),  // 2: jwtx.CheckToken_Request
	(*CheckToken_Response)(nil), // 3: jwtx.CheckToken_Response
}
var file_define_rpc_proto_depIdxs = []int32{
	0, // 0: jwtx.jwtx.MakeToken:input_type -> jwtx.MakeToken_Request
	2, // 1: jwtx.jwtx.CheckToken:input_type -> jwtx.CheckToken_Request
	1, // 2: jwtx.jwtx.MakeToken:output_type -> jwtx.MakeToken_Response
	3, // 3: jwtx.jwtx.CheckToken:output_type -> jwtx.CheckToken_Response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_define_rpc_proto_init() }
func file_define_rpc_proto_init() {
	if File_define_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_define_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_define_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_define_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_define_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_define_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_define_rpc_proto_goTypes,
		DependencyIndexes: file_define_rpc_proto_depIdxs,
		MessageInfos:      file_define_rpc_proto_msgTypes,
	}.Build()
	File_define_rpc_proto = out.File
	file_define_rpc_proto_rawDesc = nil
	file_define_rpc_proto_goTypes = nil
	file_define_rpc_proto_depIdxs = nil
}