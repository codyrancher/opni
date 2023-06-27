// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v1.0.0
// source: github.com/rancher/opni/plugins/alerting/apis/alertops/alertops.proto

package alertops

import (
	v1 "github.com/rancher/opni/pkg/apis/alerting/v1"
	_ "github.com/rancher/opni/pkg/apis/storage/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type InstallState int32

const (
	InstallState_InstallUnknown  InstallState = 0
	InstallState_NotInstalled    InstallState = 1
	InstallState_InstallUpdating InstallState = 2
	InstallState_Installed       InstallState = 3
	InstallState_Uninstalling    InstallState = 4
)

// Enum value maps for InstallState.
var (
	InstallState_name = map[int32]string{
		0: "InstallUnknown",
		1: "NotInstalled",
		2: "InstallUpdating",
		3: "Installed",
		4: "Uninstalling",
	}
	InstallState_value = map[string]int32{
		"InstallUnknown":  0,
		"NotInstalled":    1,
		"InstallUpdating": 2,
		"Installed":       3,
		"Uninstalling":    4,
	}
)

func (x InstallState) Enum() *InstallState {
	p := new(InstallState)
	*p = x
	return p
}

func (x InstallState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InstallState) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_enumTypes[0].Descriptor()
}

func (InstallState) Type() protoreflect.EnumType {
	return &file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_enumTypes[0]
}

func (x InstallState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InstallState.Descriptor instead.
func (InstallState) EnumDescriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_rawDescGZIP(), []int{0}
}

type ConnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LifecycleUuid string `protobuf:"bytes,1,opt,name=lifecycleUuid,proto3" json:"lifecycleUuid,omitempty"`
}

func (x *ConnectRequest) Reset() {
	*x = ConnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectRequest) ProtoMessage() {}

func (x *ConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectRequest.ProtoReflect.Descriptor instead.
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectRequest) GetLifecycleUuid() string {
	if x != nil {
		return x.LifecycleUuid
	}
	return ""
}

type SyncRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*v1.PutConfigRequest `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *SyncRequest) Reset() {
	*x = SyncRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncRequest) ProtoMessage() {}

func (x *SyncRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncRequest.ProtoReflect.Descriptor instead.
func (*SyncRequest) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_rawDescGZIP(), []int{1}
}

func (x *SyncRequest) GetItems() []*v1.PutConfigRequest {
	if x != nil {
		return x.Items
	}
	return nil
}

type UninstallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeleteData bool `protobuf:"varint,1,opt,name=deleteData,proto3" json:"deleteData,omitempty"`
}

func (x *UninstallRequest) Reset() {
	*x = UninstallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UninstallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UninstallRequest) ProtoMessage() {}

func (x *UninstallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UninstallRequest.ProtoReflect.Descriptor instead.
func (*UninstallRequest) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_rawDescGZIP(), []int{2}
}

func (x *UninstallRequest) GetDeleteData() bool {
	if x != nil {
		return x.DeleteData
	}
	return false
}

type InstallStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State    InstallState      `protobuf:"varint,1,opt,name=state,proto3,enum=alerting.ops.InstallState" json:"state,omitempty"`
	Version  string            `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Metadata map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *InstallStatus) Reset() {
	*x = InstallStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstallStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstallStatus) ProtoMessage() {}

func (x *InstallStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstallStatus.ProtoReflect.Descriptor instead.
func (*InstallStatus) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_rawDescGZIP(), []int{3}
}

func (x *InstallStatus) GetState() InstallState {
	if x != nil {
		return x.State
	}
	return InstallState_InstallUnknown
}

func (x *InstallStatus) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *InstallStatus) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type ClusterConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// number of replicas for the opni-alerting (odd-number for HA)
	NumReplicas int32 `protobuf:"varint,2,opt,name=numReplicas,proto3" json:"numReplicas,omitempty"`
	// Maximum time to wait for cluster
	// connections to settle before
	// evaluating notifications.
	ClusterSettleTimeout string `protobuf:"bytes,3,opt,name=clusterSettleTimeout,proto3" json:"clusterSettleTimeout,omitempty"`
	// Interval for gossip state syncs.
	// Setting this interval lower
	// (more frequent) will increase
	// convergence speeds across larger
	// clusters at the expense of
	// increased bandwidth usage.
	ClusterPushPullInterval string `protobuf:"bytes,4,opt,name=clusterPushPullInterval,proto3" json:"clusterPushPullInterval,omitempty"`
	// Interval between sending gossip
	// messages. By lowering this
	// value (more frequent) gossip
	// messages are propagated across
	// the cluster more quickly at the
	// expense of increased bandwidth.
	ClusterGossipInterval string             `protobuf:"bytes,5,opt,name=clusterGossipInterval,proto3" json:"clusterGossipInterval,omitempty"`
	ResourceLimits        *ResourceLimitSpec `protobuf:"bytes,6,opt,name=resourceLimits,proto3" json:"resourceLimits,omitempty"`
}

func (x *ClusterConfiguration) Reset() {
	*x = ClusterConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterConfiguration) ProtoMessage() {}

func (x *ClusterConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterConfiguration.ProtoReflect.Descriptor instead.
func (*ClusterConfiguration) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_rawDescGZIP(), []int{4}
}

func (x *ClusterConfiguration) GetNumReplicas() int32 {
	if x != nil {
		return x.NumReplicas
	}
	return 0
}

func (x *ClusterConfiguration) GetClusterSettleTimeout() string {
	if x != nil {
		return x.ClusterSettleTimeout
	}
	return ""
}

func (x *ClusterConfiguration) GetClusterPushPullInterval() string {
	if x != nil {
		return x.ClusterPushPullInterval
	}
	return ""
}

func (x *ClusterConfiguration) GetClusterGossipInterval() string {
	if x != nil {
		return x.ClusterGossipInterval
	}
	return ""
}

func (x *ClusterConfiguration) GetResourceLimits() *ResourceLimitSpec {
	if x != nil {
		return x.ResourceLimits
	}
	return nil
}

type ResourceLimitSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Storage resource limit for alerting volume
	Storage string `protobuf:"bytes,1,opt,name=storage,proto3" json:"storage,omitempty"`
	// CPU resource limit per replica
	Cpu string `protobuf:"bytes,2,opt,name=cpu,proto3" json:"cpu,omitempty"`
	// Memory resource limit per replica
	Memory string `protobuf:"bytes,3,opt,name=memory,proto3" json:"memory,omitempty"`
}

func (x *ResourceLimitSpec) Reset() {
	*x = ResourceLimitSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceLimitSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceLimitSpec) ProtoMessage() {}

func (x *ResourceLimitSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceLimitSpec.ProtoReflect.Descriptor instead.
func (*ResourceLimitSpec) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_rawDescGZIP(), []int{5}
}

func (x *ResourceLimitSpec) GetStorage() string {
	if x != nil {
		return x.Storage
	}
	return ""
}

func (x *ResourceLimitSpec) GetCpu() string {
	if x != nil {
		return x.Cpu
	}
	return ""
}

func (x *ResourceLimitSpec) GetMemory() string {
	if x != nil {
		return x.Memory
	}
	return ""
}

var File_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto protoreflect.FileDescriptor

var file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_rawDesc = []byte{
	0x0a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x6f, 0x70, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x6f, 0x70,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x6f, 0x70, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65,
	0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69, 0x64, 0x65,
	0x63, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x6c, 0x69, 0x66,
	0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x55, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x55, 0x75, 0x69, 0x64, 0x22,
	0x3f, 0x0a, 0x0b, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x22, 0x32, 0x0a, 0x10, 0x55, 0x6e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x22, 0xdf, 0x01, 0x0a, 0x0d, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x6f, 0x70, 0x73, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x6f, 0x70, 0x73, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa5, 0x02, 0x0a, 0x14, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x20, 0x0a, 0x0b, 0x6e, 0x75, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6e, 0x75, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x73, 0x12, 0x32, 0x0a, 0x14, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74,
	0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x14, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x38, 0x0a, 0x17, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x50, 0x75, 0x73, 0x68, 0x50, 0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50,
	0x75, 0x73, 0x68, 0x50, 0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12,
	0x34, 0x0a, 0x15, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x47, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x6f, 0x70, 0x73, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x22, 0x57,
	0x0a, 0x11, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x53,
	0x70, 0x65, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x63, 0x70, 0x75, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x70, 0x75, 0x12,
	0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2a, 0x6a, 0x0a, 0x0c, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x4e,
	0x6f, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x10, 0x01, 0x12, 0x13, 0x0a,
	0x0f, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x10,
	0x03, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x6e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x10, 0x04, 0x32, 0xf3, 0x03, 0x0a, 0x0d, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x6d, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x22, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x6f, 0x70, 0x73, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x16, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x65, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x22, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x6f, 0x70, 0x73, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x3a, 0x01, 0x2a, 0x22,
	0x0a, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x12, 0x58, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x6f, 0x70, 0x73, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x12, 0x07, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x52, 0x0a, 0x0e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x22,
	0x08, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x12, 0x5e, 0x0a, 0x10, 0x55, 0x6e, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1e, 0x2e,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x6f, 0x70, 0x73, 0x2e, 0x55, 0x6e, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x22, 0x0a, 0x2f,
	0x75, 0x6e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x32, 0x66, 0x0a, 0x10, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x72, 0x12, 0x52, 0x0a,
	0x13, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x79,
	0x6e, 0x63, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x6f, 0x70, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x6f, 0x70,
	0x73, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x00, 0x30,
	0x01, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x6f, 0x70, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_rawDescOnce sync.Once
	file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_rawDescData = file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_rawDesc
)

func file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_rawDescGZIP() []byte {
	file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_rawDescOnce.Do(func() {
		file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_rawDescData)
	})
	return file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_rawDescData
}

var file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_goTypes = []interface{}{
	(InstallState)(0),            // 0: alerting.ops.InstallState
	(*ConnectRequest)(nil),       // 1: alerting.ops.ConnectRequest
	(*SyncRequest)(nil),          // 2: alerting.ops.SyncRequest
	(*UninstallRequest)(nil),     // 3: alerting.ops.UninstallRequest
	(*InstallStatus)(nil),        // 4: alerting.ops.InstallStatus
	(*ClusterConfiguration)(nil), // 5: alerting.ops.ClusterConfiguration
	(*ResourceLimitSpec)(nil),    // 6: alerting.ops.ResourceLimitSpec
	nil,                          // 7: alerting.ops.InstallStatus.MetadataEntry
	(*v1.PutConfigRequest)(nil),  // 8: alerting.PutConfigRequest
	(*emptypb.Empty)(nil),        // 9: google.protobuf.Empty
}
var file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_depIdxs = []int32{
	8,  // 0: alerting.ops.SyncRequest.items:type_name -> alerting.PutConfigRequest
	0,  // 1: alerting.ops.InstallStatus.state:type_name -> alerting.ops.InstallState
	7,  // 2: alerting.ops.InstallStatus.metadata:type_name -> alerting.ops.InstallStatus.MetadataEntry
	6,  // 3: alerting.ops.ClusterConfiguration.resourceLimits:type_name -> alerting.ops.ResourceLimitSpec
	9,  // 4: alerting.ops.AlertingAdmin.GetClusterConfiguration:input_type -> google.protobuf.Empty
	5,  // 5: alerting.ops.AlertingAdmin.ConfigureCluster:input_type -> alerting.ops.ClusterConfiguration
	9,  // 6: alerting.ops.AlertingAdmin.GetClusterStatus:input_type -> google.protobuf.Empty
	9,  // 7: alerting.ops.AlertingAdmin.InstallCluster:input_type -> google.protobuf.Empty
	3,  // 8: alerting.ops.AlertingAdmin.UninstallCluster:input_type -> alerting.ops.UninstallRequest
	1,  // 9: alerting.ops.ConfigReconciler.ConnectRemoteSyncer:input_type -> alerting.ops.ConnectRequest
	5,  // 10: alerting.ops.AlertingAdmin.GetClusterConfiguration:output_type -> alerting.ops.ClusterConfiguration
	9,  // 11: alerting.ops.AlertingAdmin.ConfigureCluster:output_type -> google.protobuf.Empty
	4,  // 12: alerting.ops.AlertingAdmin.GetClusterStatus:output_type -> alerting.ops.InstallStatus
	9,  // 13: alerting.ops.AlertingAdmin.InstallCluster:output_type -> google.protobuf.Empty
	9,  // 14: alerting.ops.AlertingAdmin.UninstallCluster:output_type -> google.protobuf.Empty
	2,  // 15: alerting.ops.ConfigReconciler.ConnectRemoteSyncer:output_type -> alerting.ops.SyncRequest
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_init() }
func file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_init() {
	if File_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectRequest); i {
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
		file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncRequest); i {
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
		file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UninstallRequest); i {
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
		file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstallStatus); i {
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
		file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterConfiguration); i {
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
		file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceLimitSpec); i {
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
			RawDescriptor: file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_goTypes,
		DependencyIndexes: file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_depIdxs,
		EnumInfos:         file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_enumTypes,
		MessageInfos:      file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_msgTypes,
	}.Build()
	File_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto = out.File
	file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_rawDesc = nil
	file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_goTypes = nil
	file_github_com_rancher_opni_plugins_alerting_apis_alertops_alertops_proto_depIdxs = nil
}
