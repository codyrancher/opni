// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0-devel
// 	protoc        v1.0.0
// source: github.com/rancher/opni/internal/codegen/cli/cli.proto

package cli

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ClientDependencyInjectionStrategy int32

const (
	ClientDependencyInjectionStrategy_InjectIntoContext ClientDependencyInjectionStrategy = 0
	ClientDependencyInjectionStrategy_InjectAsArgument  ClientDependencyInjectionStrategy = 1
)

// Enum value maps for ClientDependencyInjectionStrategy.
var (
	ClientDependencyInjectionStrategy_name = map[int32]string{
		0: "InjectIntoContext",
		1: "InjectAsArgument",
	}
	ClientDependencyInjectionStrategy_value = map[string]int32{
		"InjectIntoContext": 0,
		"InjectAsArgument":  1,
	}
)

func (x ClientDependencyInjectionStrategy) Enum() *ClientDependencyInjectionStrategy {
	p := new(ClientDependencyInjectionStrategy)
	*p = x
	return p
}

func (x ClientDependencyInjectionStrategy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClientDependencyInjectionStrategy) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_rancher_opni_internal_codegen_cli_cli_proto_enumTypes[0].Descriptor()
}

func (ClientDependencyInjectionStrategy) Type() protoreflect.EnumType {
	return &file_github_com_rancher_opni_internal_codegen_cli_cli_proto_enumTypes[0]
}

func (x ClientDependencyInjectionStrategy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClientDependencyInjectionStrategy.Descriptor instead.
func (ClientDependencyInjectionStrategy) EnumDescriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_internal_codegen_cli_cli_proto_rawDescGZIP(), []int{0}
}

type GeneratorOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Generate                  bool                              `protobuf:"varint,1,opt,name=generate,proto3" json:"generate,omitempty"`
	GenerateDeepcopy          bool                              `protobuf:"varint,2,opt,name=generate_deepcopy,json=generateDeepcopy,proto3" json:"generate_deepcopy,omitempty"`
	ClientDependencyInjection ClientDependencyInjectionStrategy `protobuf:"varint,3,opt,name=client_dependency_injection,json=clientDependencyInjection,proto3,enum=cli.ClientDependencyInjectionStrategy" json:"client_dependency_injection,omitempty"`
}

func (x *GeneratorOptions) Reset() {
	*x = GeneratorOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_internal_codegen_cli_cli_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneratorOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneratorOptions) ProtoMessage() {}

func (x *GeneratorOptions) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_internal_codegen_cli_cli_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneratorOptions.ProtoReflect.Descriptor instead.
func (*GeneratorOptions) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_internal_codegen_cli_cli_proto_rawDescGZIP(), []int{0}
}

func (x *GeneratorOptions) GetGenerate() bool {
	if x != nil {
		return x.Generate
	}
	return false
}

func (x *GeneratorOptions) GetGenerateDeepcopy() bool {
	if x != nil {
		return x.GenerateDeepcopy
	}
	return false
}

func (x *GeneratorOptions) GetClientDependencyInjection() ClientDependencyInjectionStrategy {
	if x != nil {
		return x.ClientDependencyInjection
	}
	return ClientDependencyInjectionStrategy_InjectIntoContext
}

type FlagOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Default      string `protobuf:"bytes,1,opt,name=default,proto3" json:"default,omitempty"`
	Env          string `protobuf:"bytes,2,opt,name=env,proto3" json:"env,omitempty"`
	Secret       bool   `protobuf:"varint,3,opt,name=secret,proto3" json:"secret,omitempty"`
	TypeOverride string `protobuf:"bytes,4,opt,name=type_override,json=typeOverride,proto3" json:"type_override,omitempty"`
	Skip         bool   `protobuf:"varint,5,opt,name=skip,proto3" json:"skip,omitempty"`
}

func (x *FlagOptions) Reset() {
	*x = FlagOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_internal_codegen_cli_cli_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlagOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlagOptions) ProtoMessage() {}

func (x *FlagOptions) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_internal_codegen_cli_cli_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlagOptions.ProtoReflect.Descriptor instead.
func (*FlagOptions) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_internal_codegen_cli_cli_proto_rawDescGZIP(), []int{1}
}

func (x *FlagOptions) GetDefault() string {
	if x != nil {
		return x.Default
	}
	return ""
}

func (x *FlagOptions) GetEnv() string {
	if x != nil {
		return x.Env
	}
	return ""
}

func (x *FlagOptions) GetSecret() bool {
	if x != nil {
		return x.Secret
	}
	return false
}

func (x *FlagOptions) GetTypeOverride() string {
	if x != nil {
		return x.TypeOverride
	}
	return ""
}

func (x *FlagOptions) GetSkip() bool {
	if x != nil {
		return x.Skip
	}
	return false
}

type FlagSetOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Default *anypb.Any `protobuf:"bytes,1,opt,name=default,proto3" json:"default,omitempty"`
}

func (x *FlagSetOptions) Reset() {
	*x = FlagSetOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_internal_codegen_cli_cli_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlagSetOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlagSetOptions) ProtoMessage() {}

func (x *FlagSetOptions) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_internal_codegen_cli_cli_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlagSetOptions.ProtoReflect.Descriptor instead.
func (*FlagSetOptions) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_internal_codegen_cli_cli_proto_rawDescGZIP(), []int{2}
}

func (x *FlagSetOptions) GetDefault() *anypb.Any {
	if x != nil {
		return x.Default
	}
	return nil
}

type CommandGroupOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Use     string `protobuf:"bytes,25601,opt,name=use,proto3" json:"use,omitempty"`
	GroupId string `protobuf:"bytes,25602,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *CommandGroupOptions) Reset() {
	*x = CommandGroupOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_internal_codegen_cli_cli_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandGroupOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandGroupOptions) ProtoMessage() {}

func (x *CommandGroupOptions) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_internal_codegen_cli_cli_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandGroupOptions.ProtoReflect.Descriptor instead.
func (*CommandGroupOptions) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_internal_codegen_cli_cli_proto_rawDescGZIP(), []int{3}
}

func (x *CommandGroupOptions) GetUse() string {
	if x != nil {
		return x.Use
	}
	return ""
}

func (x *CommandGroupOptions) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

type CommandOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Use           string   `protobuf:"bytes,25601,opt,name=use,proto3" json:"use,omitempty"`
	RequiredFlags []string `protobuf:"bytes,25603,rep,name=required_flags,json=requiredFlags,proto3" json:"required_flags,omitempty"`
}

func (x *CommandOptions) Reset() {
	*x = CommandOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_internal_codegen_cli_cli_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandOptions) ProtoMessage() {}

func (x *CommandOptions) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_internal_codegen_cli_cli_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandOptions.ProtoReflect.Descriptor instead.
func (*CommandOptions) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_internal_codegen_cli_cli_proto_rawDescGZIP(), []int{4}
}

func (x *CommandOptions) GetUse() string {
	if x != nil {
		return x.Use
	}
	return ""
}

func (x *CommandOptions) GetRequiredFlags() []string {
	if x != nil {
		return x.RequiredFlags
	}
	return nil
}

var file_github_com_rancher_opni_internal_codegen_cli_cli_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*GeneratorOptions)(nil),
		Field:         25600,
		Name:          "cli.generator",
		Tag:           "bytes,25600,opt,name=generator",
		Filename:      "github.com/rancher/opni/internal/codegen/cli/cli.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*FlagOptions)(nil),
		Field:         25601,
		Name:          "cli.flag",
		Tag:           "bytes,25601,opt,name=flag",
		Filename:      "github.com/rancher/opni/internal/codegen/cli/cli.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*FlagSetOptions)(nil),
		Field:         25602,
		Name:          "cli.flag_set",
		Tag:           "bytes,25602,opt,name=flag_set",
		Filename:      "github.com/rancher/opni/internal/codegen/cli/cli.proto",
	},
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*CommandGroupOptions)(nil),
		Field:         25600,
		Name:          "cli.command_group",
		Tag:           "bytes,25600,opt,name=command_group",
		Filename:      "github.com/rancher/opni/internal/codegen/cli/cli.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*CommandOptions)(nil),
		Field:         25600,
		Name:          "cli.command",
		Tag:           "bytes,25600,opt,name=command",
		Filename:      "github.com/rancher/opni/internal/codegen/cli/cli.proto",
	},
}

// Extension fields to descriptorpb.FileOptions.
var (
	// optional cli.GeneratorOptions generator = 25600;
	E_Generator = &file_github_com_rancher_opni_internal_codegen_cli_cli_proto_extTypes[0]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional cli.FlagOptions flag = 25601;
	E_Flag = &file_github_com_rancher_opni_internal_codegen_cli_cli_proto_extTypes[1]
	// optional cli.FlagSetOptions flag_set = 25602;
	E_FlagSet = &file_github_com_rancher_opni_internal_codegen_cli_cli_proto_extTypes[2]
)

// Extension fields to descriptorpb.ServiceOptions.
var (
	// optional cli.CommandGroupOptions command_group = 25600;
	E_CommandGroup = &file_github_com_rancher_opni_internal_codegen_cli_cli_proto_extTypes[3]
)

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional cli.CommandOptions command = 25600;
	E_Command = &file_github_com_rancher_opni_internal_codegen_cli_cli_proto_extTypes[4]
)

var File_github_com_rancher_opni_internal_codegen_cli_cli_proto protoreflect.FileDescriptor

var file_github_com_rancher_opni_internal_codegen_cli_cli_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6c, 0x69, 0x2f, 0x63,
	0x6c, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x63, 0x6c, 0x69, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x01, 0x0a, 0x10, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x65, 0x65, 0x70, 0x63, 0x6f, 0x70, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x44, 0x65, 0x65, 0x70, 0x63, 0x6f, 0x70, 0x79, 0x12, 0x66, 0x0a, 0x1b, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x69, 0x6e,
	0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e,
	0x63, 0x6c, 0x69, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64,
	0x65, 0x6e, 0x63, 0x79, 0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x19, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x70,
	0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x8a, 0x01, 0x0a, 0x0b, 0x46, 0x6c, 0x61, 0x67, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e,
	0x76, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6f, 0x76, 0x65,
	0x72, 0x72, 0x69, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x79, 0x70,
	0x65, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6b, 0x69,
	0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x73, 0x6b, 0x69, 0x70, 0x22, 0x40, 0x0a,
	0x0e, 0x46, 0x6c, 0x61, 0x67, 0x53, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x2e, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x22,
	0x46, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x03, 0x75, 0x73, 0x65, 0x18, 0x81, 0xc8,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x08, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x82, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x03, 0x75, 0x73, 0x65,
	0x18, 0x81, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x73, 0x65, 0x12, 0x27, 0x0a,
	0x0e, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18,
	0x83, 0xc8, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x2a, 0x50, 0x0a, 0x21, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x15, 0x0a, 0x11, 0x49,
	0x6e, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x74, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x73, 0x41, 0x72,
	0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0x01, 0x3a, 0x53, 0x0a, 0x09, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x80, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6c,
	0x69, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x09, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x3a, 0x45, 0x0a,
	0x04, 0x66, 0x6c, 0x61, 0x67, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x81, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63,
	0x6c, 0x69, 0x2e, 0x46, 0x6c, 0x61, 0x67, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x04,
	0x66, 0x6c, 0x61, 0x67, 0x3a, 0x4f, 0x0a, 0x08, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x73, 0x65, 0x74,
	0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x82, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x46, 0x6c,
	0x61, 0x67, 0x53, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x66, 0x6c,
	0x61, 0x67, 0x53, 0x65, 0x74, 0x3a, 0x60, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x80, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x3a, 0x4f, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x80, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6c, 0x69,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f,
	0x70, 0x6e, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x64,
	0x65, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6c, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_rancher_opni_internal_codegen_cli_cli_proto_rawDescOnce sync.Once
	file_github_com_rancher_opni_internal_codegen_cli_cli_proto_rawDescData = file_github_com_rancher_opni_internal_codegen_cli_cli_proto_rawDesc
)

func file_github_com_rancher_opni_internal_codegen_cli_cli_proto_rawDescGZIP() []byte {
	file_github_com_rancher_opni_internal_codegen_cli_cli_proto_rawDescOnce.Do(func() {
		file_github_com_rancher_opni_internal_codegen_cli_cli_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_rancher_opni_internal_codegen_cli_cli_proto_rawDescData)
	})
	return file_github_com_rancher_opni_internal_codegen_cli_cli_proto_rawDescData
}

var file_github_com_rancher_opni_internal_codegen_cli_cli_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_rancher_opni_internal_codegen_cli_cli_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_github_com_rancher_opni_internal_codegen_cli_cli_proto_goTypes = []interface{}{
	(ClientDependencyInjectionStrategy)(0), // 0: cli.ClientDependencyInjectionStrategy
	(*GeneratorOptions)(nil),               // 1: cli.GeneratorOptions
	(*FlagOptions)(nil),                    // 2: cli.FlagOptions
	(*FlagSetOptions)(nil),                 // 3: cli.FlagSetOptions
	(*CommandGroupOptions)(nil),            // 4: cli.CommandGroupOptions
	(*CommandOptions)(nil),                 // 5: cli.CommandOptions
	(*anypb.Any)(nil),                      // 6: google.protobuf.Any
	(*descriptorpb.FileOptions)(nil),       // 7: google.protobuf.FileOptions
	(*descriptorpb.FieldOptions)(nil),      // 8: google.protobuf.FieldOptions
	(*descriptorpb.ServiceOptions)(nil),    // 9: google.protobuf.ServiceOptions
	(*descriptorpb.MethodOptions)(nil),     // 10: google.protobuf.MethodOptions
}
var file_github_com_rancher_opni_internal_codegen_cli_cli_proto_depIdxs = []int32{
	0,  // 0: cli.GeneratorOptions.client_dependency_injection:type_name -> cli.ClientDependencyInjectionStrategy
	6,  // 1: cli.FlagSetOptions.default:type_name -> google.protobuf.Any
	7,  // 2: cli.generator:extendee -> google.protobuf.FileOptions
	8,  // 3: cli.flag:extendee -> google.protobuf.FieldOptions
	8,  // 4: cli.flag_set:extendee -> google.protobuf.FieldOptions
	9,  // 5: cli.command_group:extendee -> google.protobuf.ServiceOptions
	10, // 6: cli.command:extendee -> google.protobuf.MethodOptions
	1,  // 7: cli.generator:type_name -> cli.GeneratorOptions
	2,  // 8: cli.flag:type_name -> cli.FlagOptions
	3,  // 9: cli.flag_set:type_name -> cli.FlagSetOptions
	4,  // 10: cli.command_group:type_name -> cli.CommandGroupOptions
	5,  // 11: cli.command:type_name -> cli.CommandOptions
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	7,  // [7:12] is the sub-list for extension type_name
	2,  // [2:7] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_github_com_rancher_opni_internal_codegen_cli_cli_proto_init() }
func file_github_com_rancher_opni_internal_codegen_cli_cli_proto_init() {
	if File_github_com_rancher_opni_internal_codegen_cli_cli_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_rancher_opni_internal_codegen_cli_cli_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneratorOptions); i {
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
		file_github_com_rancher_opni_internal_codegen_cli_cli_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlagOptions); i {
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
		file_github_com_rancher_opni_internal_codegen_cli_cli_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlagSetOptions); i {
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
		file_github_com_rancher_opni_internal_codegen_cli_cli_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandGroupOptions); i {
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
		file_github_com_rancher_opni_internal_codegen_cli_cli_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandOptions); i {
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
			RawDescriptor: file_github_com_rancher_opni_internal_codegen_cli_cli_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 5,
			NumServices:   0,
		},
		GoTypes:           file_github_com_rancher_opni_internal_codegen_cli_cli_proto_goTypes,
		DependencyIndexes: file_github_com_rancher_opni_internal_codegen_cli_cli_proto_depIdxs,
		EnumInfos:         file_github_com_rancher_opni_internal_codegen_cli_cli_proto_enumTypes,
		MessageInfos:      file_github_com_rancher_opni_internal_codegen_cli_cli_proto_msgTypes,
		ExtensionInfos:    file_github_com_rancher_opni_internal_codegen_cli_cli_proto_extTypes,
	}.Build()
	File_github_com_rancher_opni_internal_codegen_cli_cli_proto = out.File
	file_github_com_rancher_opni_internal_codegen_cli_cli_proto_rawDesc = nil
	file_github_com_rancher_opni_internal_codegen_cli_cli_proto_goTypes = nil
	file_github_com_rancher_opni_internal_codegen_cli_cli_proto_depIdxs = nil
}
