// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: prism/auth/v1/auth_service.proto

package authv1

import (
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

// Informasi minimal yang dibutuhkan auth-service untuk membuat token.
type UserInfoForToken struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	RoleName      string                 `protobuf:"bytes,3,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserInfoForToken) Reset() {
	*x = UserInfoForToken{}
	mi := &file_prism_auth_v1_auth_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInfoForToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoForToken) ProtoMessage() {}

func (x *UserInfoForToken) ProtoReflect() protoreflect.Message {
	mi := &file_prism_auth_v1_auth_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoForToken.ProtoReflect.Descriptor instead.
func (*UserInfoForToken) Descriptor() ([]byte, []int) {
	return file_prism_auth_v1_auth_service_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfoForToken) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserInfoForToken) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserInfoForToken) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

type GenerateImpersonationTokenRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TargetUser    *UserInfoForToken      `protobuf:"bytes,1,opt,name=target_user,json=targetUser,proto3" json:"target_user,omitempty"` // Info user yang akan ditiru
	ActorId       string                 `protobuf:"bytes,2,opt,name=actor_id,json=actorId,proto3" json:"actor_id,omitempty"`          // ID admin yang melakukan impersonasi
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenerateImpersonationTokenRequest) Reset() {
	*x = GenerateImpersonationTokenRequest{}
	mi := &file_prism_auth_v1_auth_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateImpersonationTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateImpersonationTokenRequest) ProtoMessage() {}

func (x *GenerateImpersonationTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_prism_auth_v1_auth_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateImpersonationTokenRequest.ProtoReflect.Descriptor instead.
func (*GenerateImpersonationTokenRequest) Descriptor() ([]byte, []int) {
	return file_prism_auth_v1_auth_service_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateImpersonationTokenRequest) GetTargetUser() *UserInfoForToken {
	if x != nil {
		return x.TargetUser
	}
	return nil
}

func (x *GenerateImpersonationTokenRequest) GetActorId() string {
	if x != nil {
		return x.ActorId
	}
	return ""
}

type GenerateImpersonationTokenResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	ExpiresAt     *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenerateImpersonationTokenResponse) Reset() {
	*x = GenerateImpersonationTokenResponse{}
	mi := &file_prism_auth_v1_auth_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateImpersonationTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateImpersonationTokenResponse) ProtoMessage() {}

func (x *GenerateImpersonationTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_prism_auth_v1_auth_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateImpersonationTokenResponse.ProtoReflect.Descriptor instead.
func (*GenerateImpersonationTokenResponse) Descriptor() ([]byte, []int) {
	return file_prism_auth_v1_auth_service_proto_rawDescGZIP(), []int{2}
}

func (x *GenerateImpersonationTokenResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *GenerateImpersonationTokenResponse) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

var File_prism_auth_v1_auth_service_proto protoreflect.FileDescriptor

const file_prism_auth_v1_auth_service_proto_rawDesc = "" +
	"\n" +
	" prism/auth/v1/auth_service.proto\x12\rprism.auth.v1\x1a\x1fgoogle/protobuf/timestamp.proto\"U\n" +
	"\x10UserInfoForToken\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12\x1b\n" +
	"\trole_name\x18\x03 \x01(\tR\broleName\"\x80\x01\n" +
	"!GenerateImpersonationTokenRequest\x12@\n" +
	"\vtarget_user\x18\x01 \x01(\v2\x1f.prism.auth.v1.UserInfoForTokenR\n" +
	"targetUser\x12\x19\n" +
	"\bactor_id\x18\x02 \x01(\tR\aactorId\"\x82\x01\n" +
	"\"GenerateImpersonationTokenResponse\x12!\n" +
	"\faccess_token\x18\x01 \x01(\tR\vaccessToken\x129\n" +
	"\n" +
	"expires_at\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\texpiresAt2\x91\x01\n" +
	"\vAuthService\x12\x81\x01\n" +
	"\x1aGenerateImpersonationToken\x120.prism.auth.v1.GenerateImpersonationTokenRequest\x1a1.prism.auth.v1.GenerateImpersonationTokenResponseBTZRgithub.com/Lumina-Enterprise-Solutions/prism-protobufs/gen/go/prism/auth/v1;authv1b\x06proto3"

var (
	file_prism_auth_v1_auth_service_proto_rawDescOnce sync.Once
	file_prism_auth_v1_auth_service_proto_rawDescData []byte
)

func file_prism_auth_v1_auth_service_proto_rawDescGZIP() []byte {
	file_prism_auth_v1_auth_service_proto_rawDescOnce.Do(func() {
		file_prism_auth_v1_auth_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_prism_auth_v1_auth_service_proto_rawDesc), len(file_prism_auth_v1_auth_service_proto_rawDesc)))
	})
	return file_prism_auth_v1_auth_service_proto_rawDescData
}

var file_prism_auth_v1_auth_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_prism_auth_v1_auth_service_proto_goTypes = []any{
	(*UserInfoForToken)(nil),                   // 0: prism.auth.v1.UserInfoForToken
	(*GenerateImpersonationTokenRequest)(nil),  // 1: prism.auth.v1.GenerateImpersonationTokenRequest
	(*GenerateImpersonationTokenResponse)(nil), // 2: prism.auth.v1.GenerateImpersonationTokenResponse
	(*timestamppb.Timestamp)(nil),              // 3: google.protobuf.Timestamp
}
var file_prism_auth_v1_auth_service_proto_depIdxs = []int32{
	0, // 0: prism.auth.v1.GenerateImpersonationTokenRequest.target_user:type_name -> prism.auth.v1.UserInfoForToken
	3, // 1: prism.auth.v1.GenerateImpersonationTokenResponse.expires_at:type_name -> google.protobuf.Timestamp
	1, // 2: prism.auth.v1.AuthService.GenerateImpersonationToken:input_type -> prism.auth.v1.GenerateImpersonationTokenRequest
	2, // 3: prism.auth.v1.AuthService.GenerateImpersonationToken:output_type -> prism.auth.v1.GenerateImpersonationTokenResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_prism_auth_v1_auth_service_proto_init() }
func file_prism_auth_v1_auth_service_proto_init() {
	if File_prism_auth_v1_auth_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_prism_auth_v1_auth_service_proto_rawDesc), len(file_prism_auth_v1_auth_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_prism_auth_v1_auth_service_proto_goTypes,
		DependencyIndexes: file_prism_auth_v1_auth_service_proto_depIdxs,
		MessageInfos:      file_prism_auth_v1_auth_service_proto_msgTypes,
	}.Build()
	File_prism_auth_v1_auth_service_proto = out.File
	file_prism_auth_v1_auth_service_proto_goTypes = nil
	file_prism_auth_v1_auth_service_proto_depIdxs = nil
}
