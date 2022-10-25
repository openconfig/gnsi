// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: github.com/openconfig/gnsi/authz/authz.proto

package authz

import (
	context "context"
	_ "github.com/openconfig/gnoi/types"
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

type ProbeResponse_Action int32

const (
	ProbeResponse_ACTION_UNSPECIFIED ProbeResponse_Action = 0
	ProbeResponse_ACTION_DENY        ProbeResponse_Action = 1
	ProbeResponse_ACTION_PERMIT      ProbeResponse_Action = 2
)

// Enum value maps for ProbeResponse_Action.
var (
	ProbeResponse_Action_name = map[int32]string{
		0: "ACTION_UNSPECIFIED",
		1: "ACTION_DENY",
		2: "ACTION_PERMIT",
	}
	ProbeResponse_Action_value = map[string]int32{
		"ACTION_UNSPECIFIED": 0,
		"ACTION_DENY":        1,
		"ACTION_PERMIT":      2,
	}
)

func (x ProbeResponse_Action) Enum() *ProbeResponse_Action {
	p := new(ProbeResponse_Action)
	*p = x
	return p
}

func (x ProbeResponse_Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProbeResponse_Action) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_openconfig_gnsi_authz_authz_proto_enumTypes[0].Descriptor()
}

func (ProbeResponse_Action) Type() protoreflect.EnumType {
	return &file_github_com_openconfig_gnsi_authz_authz_proto_enumTypes[0]
}

func (x ProbeResponse_Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProbeResponse_Action.Descriptor instead.
func (ProbeResponse_Action) EnumDescriptor() ([]byte, []int) {
	return file_github_com_openconfig_gnsi_authz_authz_proto_rawDescGZIP(), []int{6, 0}
}

type RotateAuthzRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to RotateRequest:
	//	*RotateAuthzRequest_UploadRequest
	//	*RotateAuthzRequest_FinalizeRotation
	RotateRequest  isRotateAuthzRequest_RotateRequest `protobuf_oneof:"rotate_request"`
	ForceOverwrite bool                               `protobuf:"varint,3,opt,name=force_overwrite,json=forceOverwrite,proto3" json:"force_overwrite,omitempty"`
}

func (x *RotateAuthzRequest) Reset() {
	*x = RotateAuthzRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_openconfig_gnsi_authz_authz_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RotateAuthzRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RotateAuthzRequest) ProtoMessage() {}

func (x *RotateAuthzRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_openconfig_gnsi_authz_authz_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RotateAuthzRequest.ProtoReflect.Descriptor instead.
func (*RotateAuthzRequest) Descriptor() ([]byte, []int) {
	return file_github_com_openconfig_gnsi_authz_authz_proto_rawDescGZIP(), []int{0}
}

func (m *RotateAuthzRequest) GetRotateRequest() isRotateAuthzRequest_RotateRequest {
	if m != nil {
		return m.RotateRequest
	}
	return nil
}

func (x *RotateAuthzRequest) GetUploadRequest() *UploadRequest {
	if x, ok := x.GetRotateRequest().(*RotateAuthzRequest_UploadRequest); ok {
		return x.UploadRequest
	}
	return nil
}

func (x *RotateAuthzRequest) GetFinalizeRotation() *FinalizeRequest {
	if x, ok := x.GetRotateRequest().(*RotateAuthzRequest_FinalizeRotation); ok {
		return x.FinalizeRotation
	}
	return nil
}

func (x *RotateAuthzRequest) GetForceOverwrite() bool {
	if x != nil {
		return x.ForceOverwrite
	}
	return false
}

type isRotateAuthzRequest_RotateRequest interface {
	isRotateAuthzRequest_RotateRequest()
}

type RotateAuthzRequest_UploadRequest struct {
	UploadRequest *UploadRequest `protobuf:"bytes,1,opt,name=upload_request,json=uploadRequest,proto3,oneof"`
}

type RotateAuthzRequest_FinalizeRotation struct {
	FinalizeRotation *FinalizeRequest `protobuf:"bytes,2,opt,name=finalize_rotation,json=finalizeRotation,proto3,oneof"`
}

func (*RotateAuthzRequest_UploadRequest) isRotateAuthzRequest_RotateRequest() {}

func (*RotateAuthzRequest_FinalizeRotation) isRotateAuthzRequest_RotateRequest() {}

type RotateAuthzResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to RotateResponse:
	//	*RotateAuthzResponse_UploadResponse
	RotateResponse isRotateAuthzResponse_RotateResponse `protobuf_oneof:"rotate_response"`
}

func (x *RotateAuthzResponse) Reset() {
	*x = RotateAuthzResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_openconfig_gnsi_authz_authz_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RotateAuthzResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RotateAuthzResponse) ProtoMessage() {}

func (x *RotateAuthzResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_openconfig_gnsi_authz_authz_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RotateAuthzResponse.ProtoReflect.Descriptor instead.
func (*RotateAuthzResponse) Descriptor() ([]byte, []int) {
	return file_github_com_openconfig_gnsi_authz_authz_proto_rawDescGZIP(), []int{1}
}

func (m *RotateAuthzResponse) GetRotateResponse() isRotateAuthzResponse_RotateResponse {
	if m != nil {
		return m.RotateResponse
	}
	return nil
}

func (x *RotateAuthzResponse) GetUploadResponse() *UploadResponse {
	if x, ok := x.GetRotateResponse().(*RotateAuthzResponse_UploadResponse); ok {
		return x.UploadResponse
	}
	return nil
}

type isRotateAuthzResponse_RotateResponse interface {
	isRotateAuthzResponse_RotateResponse()
}

type RotateAuthzResponse_UploadResponse struct {
	UploadResponse *UploadResponse `protobuf:"bytes,1,opt,name=upload_response,json=uploadResponse,proto3,oneof"`
}

func (*RotateAuthzResponse_UploadResponse) isRotateAuthzResponse_RotateResponse() {}

type FinalizeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FinalizeRequest) Reset() {
	*x = FinalizeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_openconfig_gnsi_authz_authz_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinalizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinalizeRequest) ProtoMessage() {}

func (x *FinalizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_openconfig_gnsi_authz_authz_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinalizeRequest.ProtoReflect.Descriptor instead.
func (*FinalizeRequest) Descriptor() ([]byte, []int) {
	return file_github_com_openconfig_gnsi_authz_authz_proto_rawDescGZIP(), []int{2}
}

type UploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version   string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	CreatedOn uint64 `protobuf:"varint,2,opt,name=created_on,json=createdOn,proto3" json:"created_on,omitempty"`
	Policy    string `protobuf:"bytes,3,opt,name=policy,proto3" json:"policy,omitempty"`
}

func (x *UploadRequest) Reset() {
	*x = UploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_openconfig_gnsi_authz_authz_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRequest) ProtoMessage() {}

func (x *UploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_openconfig_gnsi_authz_authz_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadRequest.ProtoReflect.Descriptor instead.
func (*UploadRequest) Descriptor() ([]byte, []int) {
	return file_github_com_openconfig_gnsi_authz_authz_proto_rawDescGZIP(), []int{3}
}

func (x *UploadRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *UploadRequest) GetCreatedOn() uint64 {
	if x != nil {
		return x.CreatedOn
	}
	return 0
}

func (x *UploadRequest) GetPolicy() string {
	if x != nil {
		return x.Policy
	}
	return ""
}

type UploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UploadResponse) Reset() {
	*x = UploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_openconfig_gnsi_authz_authz_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadResponse) ProtoMessage() {}

func (x *UploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_openconfig_gnsi_authz_authz_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadResponse.ProtoReflect.Descriptor instead.
func (*UploadResponse) Descriptor() ([]byte, []int) {
	return file_github_com_openconfig_gnsi_authz_authz_proto_rawDescGZIP(), []int{4}
}

type ProbeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Rpc  string `protobuf:"bytes,2,opt,name=rpc,proto3" json:"rpc,omitempty"`
}

func (x *ProbeRequest) Reset() {
	*x = ProbeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_openconfig_gnsi_authz_authz_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProbeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProbeRequest) ProtoMessage() {}

func (x *ProbeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_openconfig_gnsi_authz_authz_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProbeRequest.ProtoReflect.Descriptor instead.
func (*ProbeRequest) Descriptor() ([]byte, []int) {
	return file_github_com_openconfig_gnsi_authz_authz_proto_rawDescGZIP(), []int{5}
}

func (x *ProbeRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *ProbeRequest) GetRpc() string {
	if x != nil {
		return x.Rpc
	}
	return ""
}

type ProbeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action  ProbeResponse_Action `protobuf:"varint,1,opt,name=action,proto3,enum=gnsi.authz.ProbeResponse_Action" json:"action,omitempty"`
	Version string               `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ProbeResponse) Reset() {
	*x = ProbeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_openconfig_gnsi_authz_authz_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProbeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProbeResponse) ProtoMessage() {}

func (x *ProbeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_openconfig_gnsi_authz_authz_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProbeResponse.ProtoReflect.Descriptor instead.
func (*ProbeResponse) Descriptor() ([]byte, []int) {
	return file_github_com_openconfig_gnsi_authz_authz_proto_rawDescGZIP(), []int{6}
}

func (x *ProbeResponse) GetAction() ProbeResponse_Action {
	if x != nil {
		return x.Action
	}
	return ProbeResponse_ACTION_UNSPECIFIED
}

func (x *ProbeResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_openconfig_gnsi_authz_authz_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_openconfig_gnsi_authz_authz_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_github_com_openconfig_gnsi_authz_authz_proto_rawDescGZIP(), []int{7}
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version   string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	CreatedOn uint64 `protobuf:"varint,2,opt,name=created_on,json=createdOn,proto3" json:"created_on,omitempty"`
	Policy    string `protobuf:"bytes,3,opt,name=policy,proto3" json:"policy,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_openconfig_gnsi_authz_authz_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_openconfig_gnsi_authz_authz_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_github_com_openconfig_gnsi_authz_authz_proto_rawDescGZIP(), []int{8}
}

func (x *GetResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *GetResponse) GetCreatedOn() uint64 {
	if x != nil {
		return x.CreatedOn
	}
	return 0
}

func (x *GetResponse) GetPolicy() string {
	if x != nil {
		return x.Policy
	}
	return ""
}

var File_github_com_openconfig_gnsi_authz_authz_proto protoreflect.FileDescriptor

var file_github_com_openconfig_gnsi_authz_authz_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65,
	0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x67, 0x6e, 0x73, 0x69, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x67, 0x6e, 0x73, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x1a, 0x2c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x67, 0x6e, 0x6f, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf, 0x01, 0x0a, 0x12, 0x52, 0x6f, 0x74,
	0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x42, 0x0a, 0x0e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6e, 0x73, 0x69, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x7a, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x48, 0x00, 0x52, 0x0d, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x11, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x5f,
	0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x67, 0x6e, 0x73, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x46, 0x69, 0x6e, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x10, 0x66,
	0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x27, 0x0a, 0x0f, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x4f,
	0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x72, 0x6f, 0x74, 0x61,
	0x74, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x6f, 0x0a, 0x13, 0x52, 0x6f,
	0x74, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x45, 0x0a, 0x0f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6e, 0x73,
	0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x11, 0x0a, 0x0f, 0x72, 0x6f, 0x74, 0x61,
	0x74, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x0a, 0x0f, 0x46,
	0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x60,
	0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x22, 0x10, 0x0a, 0x0e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x34, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x70, 0x63, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x70, 0x63, 0x22, 0xa9, 0x01, 0x0a, 0x0d, 0x50, 0x72, 0x6f,
	0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x67, 0x6e, 0x73,
	0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x44,
	0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x4e, 0x59, 0x10,
	0x01, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x45, 0x52, 0x4d,
	0x49, 0x54, 0x10, 0x02, 0x22, 0x0c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x5e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x32, 0xcc, 0x01, 0x0a, 0x05, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x12, 0x4d, 0x0a, 0x06,
	0x52, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x67, 0x6e, 0x73, 0x69, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x7a, 0x2e, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x67, 0x6e, 0x73, 0x69, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x7a, 0x2e, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x3c, 0x0a, 0x05, 0x50,
	0x72, 0x6f, 0x62, 0x65, 0x12, 0x18, 0x2e, 0x67, 0x6e, 0x73, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x7a, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x67, 0x6e, 0x73, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x50, 0x72, 0x6f, 0x62,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x16, 0x2e, 0x67, 0x6e, 0x73, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67, 0x6e, 0x73, 0x69, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x2a, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6f, 0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x67, 0x6e, 0x73, 0x69, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x7a, 0xd2, 0x3e, 0x05, 0x30, 0x2e, 0x31, 0x2e, 0x30, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_openconfig_gnsi_authz_authz_proto_rawDescOnce sync.Once
	file_github_com_openconfig_gnsi_authz_authz_proto_rawDescData = file_github_com_openconfig_gnsi_authz_authz_proto_rawDesc
)

func file_github_com_openconfig_gnsi_authz_authz_proto_rawDescGZIP() []byte {
	file_github_com_openconfig_gnsi_authz_authz_proto_rawDescOnce.Do(func() {
		file_github_com_openconfig_gnsi_authz_authz_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_openconfig_gnsi_authz_authz_proto_rawDescData)
	})
	return file_github_com_openconfig_gnsi_authz_authz_proto_rawDescData
}

var file_github_com_openconfig_gnsi_authz_authz_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_openconfig_gnsi_authz_authz_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_github_com_openconfig_gnsi_authz_authz_proto_goTypes = []interface{}{
	(ProbeResponse_Action)(0),   // 0: gnsi.authz.ProbeResponse.Action
	(*RotateAuthzRequest)(nil),  // 1: gnsi.authz.RotateAuthzRequest
	(*RotateAuthzResponse)(nil), // 2: gnsi.authz.RotateAuthzResponse
	(*FinalizeRequest)(nil),     // 3: gnsi.authz.FinalizeRequest
	(*UploadRequest)(nil),       // 4: gnsi.authz.UploadRequest
	(*UploadResponse)(nil),      // 5: gnsi.authz.UploadResponse
	(*ProbeRequest)(nil),        // 6: gnsi.authz.ProbeRequest
	(*ProbeResponse)(nil),       // 7: gnsi.authz.ProbeResponse
	(*GetRequest)(nil),          // 8: gnsi.authz.GetRequest
	(*GetResponse)(nil),         // 9: gnsi.authz.GetResponse
}
var file_github_com_openconfig_gnsi_authz_authz_proto_depIdxs = []int32{
	4, // 0: gnsi.authz.RotateAuthzRequest.upload_request:type_name -> gnsi.authz.UploadRequest
	3, // 1: gnsi.authz.RotateAuthzRequest.finalize_rotation:type_name -> gnsi.authz.FinalizeRequest
	5, // 2: gnsi.authz.RotateAuthzResponse.upload_response:type_name -> gnsi.authz.UploadResponse
	0, // 3: gnsi.authz.ProbeResponse.action:type_name -> gnsi.authz.ProbeResponse.Action
	1, // 4: gnsi.authz.Authz.Rotate:input_type -> gnsi.authz.RotateAuthzRequest
	6, // 5: gnsi.authz.Authz.Probe:input_type -> gnsi.authz.ProbeRequest
	8, // 6: gnsi.authz.Authz.Get:input_type -> gnsi.authz.GetRequest
	2, // 7: gnsi.authz.Authz.Rotate:output_type -> gnsi.authz.RotateAuthzResponse
	7, // 8: gnsi.authz.Authz.Probe:output_type -> gnsi.authz.ProbeResponse
	9, // 9: gnsi.authz.Authz.Get:output_type -> gnsi.authz.GetResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_github_com_openconfig_gnsi_authz_authz_proto_init() }
func file_github_com_openconfig_gnsi_authz_authz_proto_init() {
	if File_github_com_openconfig_gnsi_authz_authz_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_openconfig_gnsi_authz_authz_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RotateAuthzRequest); i {
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
		file_github_com_openconfig_gnsi_authz_authz_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RotateAuthzResponse); i {
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
		file_github_com_openconfig_gnsi_authz_authz_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinalizeRequest); i {
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
		file_github_com_openconfig_gnsi_authz_authz_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadRequest); i {
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
		file_github_com_openconfig_gnsi_authz_authz_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadResponse); i {
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
		file_github_com_openconfig_gnsi_authz_authz_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProbeRequest); i {
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
		file_github_com_openconfig_gnsi_authz_authz_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProbeResponse); i {
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
		file_github_com_openconfig_gnsi_authz_authz_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_github_com_openconfig_gnsi_authz_authz_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
	file_github_com_openconfig_gnsi_authz_authz_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*RotateAuthzRequest_UploadRequest)(nil),
		(*RotateAuthzRequest_FinalizeRotation)(nil),
	}
	file_github_com_openconfig_gnsi_authz_authz_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*RotateAuthzResponse_UploadResponse)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_openconfig_gnsi_authz_authz_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_openconfig_gnsi_authz_authz_proto_goTypes,
		DependencyIndexes: file_github_com_openconfig_gnsi_authz_authz_proto_depIdxs,
		EnumInfos:         file_github_com_openconfig_gnsi_authz_authz_proto_enumTypes,
		MessageInfos:      file_github_com_openconfig_gnsi_authz_authz_proto_msgTypes,
	}.Build()
	File_github_com_openconfig_gnsi_authz_authz_proto = out.File
	file_github_com_openconfig_gnsi_authz_authz_proto_rawDesc = nil
	file_github_com_openconfig_gnsi_authz_authz_proto_goTypes = nil
	file_github_com_openconfig_gnsi_authz_authz_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthzClient is the client API for Authz service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthzClient interface {
	Rotate(ctx context.Context, opts ...grpc.CallOption) (Authz_RotateClient, error)
	Probe(ctx context.Context, in *ProbeRequest, opts ...grpc.CallOption) (*ProbeResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type authzClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthzClient(cc grpc.ClientConnInterface) AuthzClient {
	return &authzClient{cc}
}

func (c *authzClient) Rotate(ctx context.Context, opts ...grpc.CallOption) (Authz_RotateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Authz_serviceDesc.Streams[0], "/gnsi.authz.Authz/Rotate", opts...)
	if err != nil {
		return nil, err
	}
	x := &authzRotateClient{stream}
	return x, nil
}

type Authz_RotateClient interface {
	Send(*RotateAuthzRequest) error
	Recv() (*RotateAuthzResponse, error)
	grpc.ClientStream
}

type authzRotateClient struct {
	grpc.ClientStream
}

func (x *authzRotateClient) Send(m *RotateAuthzRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *authzRotateClient) Recv() (*RotateAuthzResponse, error) {
	m := new(RotateAuthzResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *authzClient) Probe(ctx context.Context, in *ProbeRequest, opts ...grpc.CallOption) (*ProbeResponse, error) {
	out := new(ProbeResponse)
	err := c.cc.Invoke(ctx, "/gnsi.authz.Authz/Probe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authzClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/gnsi.authz.Authz/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthzServer is the server API for Authz service.
type AuthzServer interface {
	Rotate(Authz_RotateServer) error
	Probe(context.Context, *ProbeRequest) (*ProbeResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
}

// UnimplementedAuthzServer can be embedded to have forward compatible implementations.
type UnimplementedAuthzServer struct {
}

func (*UnimplementedAuthzServer) Rotate(Authz_RotateServer) error {
	return status.Errorf(codes.Unimplemented, "method Rotate not implemented")
}
func (*UnimplementedAuthzServer) Probe(context.Context, *ProbeRequest) (*ProbeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Probe not implemented")
}
func (*UnimplementedAuthzServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterAuthzServer(s *grpc.Server, srv AuthzServer) {
	s.RegisterService(&_Authz_serviceDesc, srv)
}

func _Authz_Rotate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AuthzServer).Rotate(&authzRotateServer{stream})
}

type Authz_RotateServer interface {
	Send(*RotateAuthzResponse) error
	Recv() (*RotateAuthzRequest, error)
	grpc.ServerStream
}

type authzRotateServer struct {
	grpc.ServerStream
}

func (x *authzRotateServer) Send(m *RotateAuthzResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *authzRotateServer) Recv() (*RotateAuthzRequest, error) {
	m := new(RotateAuthzRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Authz_Probe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProbeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthzServer).Probe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnsi.authz.Authz/Probe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthzServer).Probe(ctx, req.(*ProbeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authz_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthzServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnsi.authz.Authz/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthzServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authz_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gnsi.authz.Authz",
	HandlerType: (*AuthzServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Probe",
			Handler:    _Authz_Probe_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Authz_Get_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Rotate",
			Handler:       _Authz_Rotate_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "github.com/openconfig/gnsi/authz/authz.proto",
}
