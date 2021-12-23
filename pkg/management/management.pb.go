// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	ragù          v0.1.0
// source: pkg/management/management.proto

package management

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateBootstrapTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TTL *durationpb.Duration `protobuf:"bytes,1,opt,name=TTL,proto3" json:"TTL,omitempty"`
}

func (x *CreateBootstrapTokenRequest) Reset() {
	*x = CreateBootstrapTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_management_management_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBootstrapTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBootstrapTokenRequest) ProtoMessage() {}

func (x *CreateBootstrapTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_management_management_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBootstrapTokenRequest.ProtoReflect.Descriptor instead.
func (*CreateBootstrapTokenRequest) Descriptor() ([]byte, []int) {
	return file_pkg_management_management_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBootstrapTokenRequest) GetTTL() *durationpb.Duration {
	if x != nil {
		return x.TTL
	}
	return nil
}

type RevokeBootstrapTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenID string `protobuf:"bytes,1,opt,name=TokenID,proto3" json:"TokenID,omitempty"`
}

func (x *RevokeBootstrapTokenRequest) Reset() {
	*x = RevokeBootstrapTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_management_management_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevokeBootstrapTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokeBootstrapTokenRequest) ProtoMessage() {}

func (x *RevokeBootstrapTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_management_management_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokeBootstrapTokenRequest.ProtoReflect.Descriptor instead.
func (*RevokeBootstrapTokenRequest) Descriptor() ([]byte, []int) {
	return file_pkg_management_management_proto_rawDescGZIP(), []int{1}
}

func (x *RevokeBootstrapTokenRequest) GetTokenID() string {
	if x != nil {
		return x.TokenID
	}
	return ""
}

type BootstrapToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenID []byte `protobuf:"bytes,1,opt,name=TokenID,proto3" json:"TokenID,omitempty"`
	Secret  []byte `protobuf:"bytes,2,opt,name=Secret,proto3" json:"Secret,omitempty"`
	LeaseID int64  `protobuf:"varint,3,opt,name=LeaseID,proto3" json:"LeaseID,omitempty"`
	TTL     int64  `protobuf:"varint,4,opt,name=TTL,proto3" json:"TTL,omitempty"`
}

func (x *BootstrapToken) Reset() {
	*x = BootstrapToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_management_management_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BootstrapToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BootstrapToken) ProtoMessage() {}

func (x *BootstrapToken) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_management_management_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BootstrapToken.ProtoReflect.Descriptor instead.
func (*BootstrapToken) Descriptor() ([]byte, []int) {
	return file_pkg_management_management_proto_rawDescGZIP(), []int{2}
}

func (x *BootstrapToken) GetTokenID() []byte {
	if x != nil {
		return x.TokenID
	}
	return nil
}

func (x *BootstrapToken) GetSecret() []byte {
	if x != nil {
		return x.Secret
	}
	return nil
}

func (x *BootstrapToken) GetLeaseID() int64 {
	if x != nil {
		return x.LeaseID
	}
	return 0
}

func (x *BootstrapToken) GetTTL() int64 {
	if x != nil {
		return x.TTL
	}
	return 0
}

type ListBootstrapTokensRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListBootstrapTokensRequest) Reset() {
	*x = ListBootstrapTokensRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_management_management_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBootstrapTokensRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBootstrapTokensRequest) ProtoMessage() {}

func (x *ListBootstrapTokensRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_management_management_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBootstrapTokensRequest.ProtoReflect.Descriptor instead.
func (*ListBootstrapTokensRequest) Descriptor() ([]byte, []int) {
	return file_pkg_management_management_proto_rawDescGZIP(), []int{3}
}

type ListBootstrapTokensResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tokens []*BootstrapToken `protobuf:"bytes,1,rep,name=Tokens,proto3" json:"Tokens,omitempty"`
}

func (x *ListBootstrapTokensResponse) Reset() {
	*x = ListBootstrapTokensResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_management_management_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBootstrapTokensResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBootstrapTokensResponse) ProtoMessage() {}

func (x *ListBootstrapTokensResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_management_management_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBootstrapTokensResponse.ProtoReflect.Descriptor instead.
func (*ListBootstrapTokensResponse) Descriptor() ([]byte, []int) {
	return file_pkg_management_management_proto_rawDescGZIP(), []int{4}
}

func (x *ListBootstrapTokensResponse) GetTokens() []*BootstrapToken {
	if x != nil {
		return x.Tokens
	}
	return nil
}

type Tenant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	OrgID     string `protobuf:"bytes,2,opt,name=OrgID,proto3" json:"OrgID,omitempty"`
	PublicKey []byte `protobuf:"bytes,3,opt,name=PublicKey,proto3" json:"PublicKey,omitempty"`
}

func (x *Tenant) Reset() {
	*x = Tenant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_management_management_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tenant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tenant) ProtoMessage() {}

func (x *Tenant) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_management_management_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tenant.ProtoReflect.Descriptor instead.
func (*Tenant) Descriptor() ([]byte, []int) {
	return file_pkg_management_management_proto_rawDescGZIP(), []int{5}
}

func (x *Tenant) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Tenant) GetOrgID() string {
	if x != nil {
		return x.OrgID
	}
	return ""
}

func (x *Tenant) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

type ListTenantsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenants []*Tenant `protobuf:"bytes,1,rep,name=Tenants,proto3" json:"Tenants,omitempty"`
}

func (x *ListTenantsResponse) Reset() {
	*x = ListTenantsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_management_management_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTenantsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTenantsResponse) ProtoMessage() {}

func (x *ListTenantsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_management_management_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTenantsResponse.ProtoReflect.Descriptor instead.
func (*ListTenantsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_management_management_proto_rawDescGZIP(), []int{6}
}

func (x *ListTenantsResponse) GetTenants() []*Tenant {
	if x != nil {
		return x.Tenants
	}
	return nil
}

var File_pkg_management_management_proto protoreflect.FileDescriptor

var file_pkg_management_management_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a, 0x1b, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x03, 0x54, 0x54, 0x4c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x00, 0x3a, 0x00, 0x22, 0x32, 0x0a, 0x1b, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x42,
	0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x11, 0x0a, 0x07, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x00, 0x3a, 0x00, 0x22, 0x59, 0x0a, 0x0e, 0x42, 0x6f, 0x6f,
	0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x11, 0x0a, 0x07, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x00, 0x12, 0x10,
	0x0a, 0x06, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x00,
	0x12, 0x11, 0x0a, 0x07, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x00, 0x12, 0x0d, 0x0a, 0x03, 0x54, 0x54, 0x4c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x42, 0x00, 0x3a, 0x00, 0x22, 0x1e, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x74,
	0x73, 0x74, 0x72, 0x61, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x3a, 0x00, 0x22, 0x4d, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x74,
	0x73, 0x74, 0x72, 0x61, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42,
	0x00, 0x3a, 0x00, 0x22, 0x3e, 0x0a, 0x06, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x0c, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x00, 0x12, 0x0f, 0x0a, 0x05, 0x4f,
	0x72, 0x67, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x00, 0x12, 0x13, 0x0a, 0x09,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x42,
	0x00, 0x3a, 0x00, 0x22, 0x3e, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x07, 0x54, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x42,
	0x00, 0x3a, 0x00, 0x32, 0x8c, 0x03, 0x0a, 0x0a, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x61, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x74,
	0x73, 0x74, 0x72, 0x61, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x27, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f,
	0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x00, 0x28, 0x00, 0x30, 0x00, 0x12, 0x5d, 0x0a, 0x14, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x42,
	0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x27, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x76, 0x6f, 0x6b,
	0x65, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x28, 0x00, 0x30, 0x00, 0x12, 0x6c, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x74,
	0x73, 0x74, 0x72, 0x61, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x26, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f,
	0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x00,
	0x30, 0x00, 0x12, 0x4c, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x00, 0x30, 0x00,
	0x1a, 0x00, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6b, 0x72, 0x61, 0x6c, 0x69, 0x63, 0x6b, 0x79, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2d, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_management_management_proto_rawDescOnce sync.Once
	file_pkg_management_management_proto_rawDescData = file_pkg_management_management_proto_rawDesc
)

func file_pkg_management_management_proto_rawDescGZIP() []byte {
	file_pkg_management_management_proto_rawDescOnce.Do(func() {
		file_pkg_management_management_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_management_management_proto_rawDescData)
	})
	return file_pkg_management_management_proto_rawDescData
}

var file_pkg_management_management_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pkg_management_management_proto_goTypes = []interface{}{
	(*CreateBootstrapTokenRequest)(nil), // 0: management.CreateBootstrapTokenRequest
	(*RevokeBootstrapTokenRequest)(nil), // 1: management.RevokeBootstrapTokenRequest
	(*BootstrapToken)(nil),              // 2: management.BootstrapToken
	(*ListBootstrapTokensRequest)(nil),  // 3: management.ListBootstrapTokensRequest
	(*ListBootstrapTokensResponse)(nil), // 4: management.ListBootstrapTokensResponse
	(*Tenant)(nil),                      // 5: management.Tenant
	(*ListTenantsResponse)(nil),         // 6: management.ListTenantsResponse
	(*durationpb.Duration)(nil),         // 7: google.protobuf.Duration
	(*emptypb.Empty)(nil),               // 8: google.protobuf.Empty
}
var file_pkg_management_management_proto_depIdxs = []int32{
	7, // 0: management.CreateBootstrapTokenRequest.TTL:type_name -> google.protobuf.Duration
	2, // 1: management.ListBootstrapTokensResponse.Tokens:type_name -> management.BootstrapToken
	5, // 2: management.ListTenantsResponse.Tenants:type_name -> management.Tenant
	0, // 3: management.Management.CreateBootstrapToken:input_type -> management.CreateBootstrapTokenRequest
	1, // 4: management.Management.RevokeBootstrapToken:input_type -> management.RevokeBootstrapTokenRequest
	3, // 5: management.Management.ListBootstrapTokens:input_type -> management.ListBootstrapTokensRequest
	8, // 6: management.Management.ListTenants:input_type -> google.protobuf.Empty
	2, // 7: management.Management.CreateBootstrapToken:output_type -> management.BootstrapToken
	8, // 8: management.Management.RevokeBootstrapToken:output_type -> google.protobuf.Empty
	4, // 9: management.Management.ListBootstrapTokens:output_type -> management.ListBootstrapTokensResponse
	6, // 10: management.Management.ListTenants:output_type -> management.ListTenantsResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_management_management_proto_init() }
func file_pkg_management_management_proto_init() {
	if File_pkg_management_management_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_management_management_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBootstrapTokenRequest); i {
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
		file_pkg_management_management_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RevokeBootstrapTokenRequest); i {
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
		file_pkg_management_management_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BootstrapToken); i {
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
		file_pkg_management_management_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBootstrapTokensRequest); i {
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
		file_pkg_management_management_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBootstrapTokensResponse); i {
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
		file_pkg_management_management_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tenant); i {
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
		file_pkg_management_management_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTenantsResponse); i {
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
			RawDescriptor: file_pkg_management_management_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_management_management_proto_goTypes,
		DependencyIndexes: file_pkg_management_management_proto_depIdxs,
		MessageInfos:      file_pkg_management_management_proto_msgTypes,
	}.Build()
	File_pkg_management_management_proto = out.File
	file_pkg_management_management_proto_rawDesc = nil
	file_pkg_management_management_proto_goTypes = nil
	file_pkg_management_management_proto_depIdxs = nil
}
