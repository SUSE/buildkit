// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.11.4
// source: github.com/moby/buildkit/session/secrets/secrets.proto

package secrets

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

type GetSecretRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string            `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Annotations map[string]string `protobuf:"bytes,2,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetSecretRequest) Reset() {
	*x = GetSecretRequest{}
	mi := &file_github_com_moby_buildkit_session_secrets_secrets_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSecretRequest) ProtoMessage() {}

func (x *GetSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_moby_buildkit_session_secrets_secrets_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSecretRequest.ProtoReflect.Descriptor instead.
func (*GetSecretRequest) Descriptor() ([]byte, []int) {
	return file_github_com_moby_buildkit_session_secrets_secrets_proto_rawDescGZIP(), []int{0}
}

func (x *GetSecretRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *GetSecretRequest) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

type GetSecretResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetSecretResponse) Reset() {
	*x = GetSecretResponse{}
	mi := &file_github_com_moby_buildkit_session_secrets_secrets_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSecretResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSecretResponse) ProtoMessage() {}

func (x *GetSecretResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_moby_buildkit_session_secrets_secrets_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSecretResponse.ProtoReflect.Descriptor instead.
func (*GetSecretResponse) Descriptor() ([]byte, []int) {
	return file_github_com_moby_buildkit_session_secrets_secrets_proto_rawDescGZIP(), []int{1}
}

func (x *GetSecretResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_github_com_moby_buildkit_session_secrets_secrets_proto protoreflect.FileDescriptor

var file_github_com_moby_buildkit_session_secrets_secrets_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x62,
	0x79, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x6b, 0x69, 0x74, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x6d, 0x6f, 0x62, 0x79, 0x2e, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x6b, 0x69, 0x74, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x22, 0xc1, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x5d, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x6d,
	0x6f, 0x62, 0x79, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x6b, 0x69, 0x74, 0x2e, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x27, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32,
	0x6f, 0x0a, 0x07, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x12, 0x64, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x2a, 0x2e, 0x6d, 0x6f, 0x62, 0x79, 0x2e, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x6b, 0x69, 0x74, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6d, 0x6f, 0x62, 0x79, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x6b, 0x69, 0x74, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d,
	0x6f, 0x62, 0x79, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x6b, 0x69, 0x74, 0x2f, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_moby_buildkit_session_secrets_secrets_proto_rawDescOnce sync.Once
	file_github_com_moby_buildkit_session_secrets_secrets_proto_rawDescData = file_github_com_moby_buildkit_session_secrets_secrets_proto_rawDesc
)

func file_github_com_moby_buildkit_session_secrets_secrets_proto_rawDescGZIP() []byte {
	file_github_com_moby_buildkit_session_secrets_secrets_proto_rawDescOnce.Do(func() {
		file_github_com_moby_buildkit_session_secrets_secrets_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_moby_buildkit_session_secrets_secrets_proto_rawDescData)
	})
	return file_github_com_moby_buildkit_session_secrets_secrets_proto_rawDescData
}

var file_github_com_moby_buildkit_session_secrets_secrets_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_moby_buildkit_session_secrets_secrets_proto_goTypes = []any{
	(*GetSecretRequest)(nil),  // 0: moby.buildkit.secrets.v1.GetSecretRequest
	(*GetSecretResponse)(nil), // 1: moby.buildkit.secrets.v1.GetSecretResponse
	nil,                       // 2: moby.buildkit.secrets.v1.GetSecretRequest.AnnotationsEntry
}
var file_github_com_moby_buildkit_session_secrets_secrets_proto_depIdxs = []int32{
	2, // 0: moby.buildkit.secrets.v1.GetSecretRequest.annotations:type_name -> moby.buildkit.secrets.v1.GetSecretRequest.AnnotationsEntry
	0, // 1: moby.buildkit.secrets.v1.Secrets.GetSecret:input_type -> moby.buildkit.secrets.v1.GetSecretRequest
	1, // 2: moby.buildkit.secrets.v1.Secrets.GetSecret:output_type -> moby.buildkit.secrets.v1.GetSecretResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_github_com_moby_buildkit_session_secrets_secrets_proto_init() }
func file_github_com_moby_buildkit_session_secrets_secrets_proto_init() {
	if File_github_com_moby_buildkit_session_secrets_secrets_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_moby_buildkit_session_secrets_secrets_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_moby_buildkit_session_secrets_secrets_proto_goTypes,
		DependencyIndexes: file_github_com_moby_buildkit_session_secrets_secrets_proto_depIdxs,
		MessageInfos:      file_github_com_moby_buildkit_session_secrets_secrets_proto_msgTypes,
	}.Build()
	File_github_com_moby_buildkit_session_secrets_secrets_proto = out.File
	file_github_com_moby_buildkit_session_secrets_secrets_proto_rawDesc = nil
	file_github_com_moby_buildkit_session_secrets_secrets_proto_goTypes = nil
	file_github_com_moby_buildkit_session_secrets_secrets_proto_depIdxs = nil
}
