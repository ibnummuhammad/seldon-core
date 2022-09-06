// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.1
// source: mlops/agent/agent.proto

package agent

import (
	scheduler "github.com/seldonio/seldon-core/scheduler/apis/mlops/scheduler"
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

type ModelEventMessage_Event int32

const (
	ModelEventMessage_UNKNOWN_EVENT    ModelEventMessage_Event = 0
	ModelEventMessage_LOAD_FAIL_MEMORY ModelEventMessage_Event = 1
	ModelEventMessage_LOADED           ModelEventMessage_Event = 2
	ModelEventMessage_LOAD_FAILED      ModelEventMessage_Event = 3
	ModelEventMessage_UNLOADED         ModelEventMessage_Event = 4
	ModelEventMessage_UNLOAD_FAILED    ModelEventMessage_Event = 5
	ModelEventMessage_REMOVED          ModelEventMessage_Event = 6 // unloaded and removed from local PVC
	ModelEventMessage_REMOVE_FAILED    ModelEventMessage_Event = 7
	ModelEventMessage_SCALE_UP_REQUEST ModelEventMessage_Event = 8
	ModelEventMessage_RSYNC            ModelEventMessage_Event = 9 // Ask server for all models that need to be loaded
)

// Enum value maps for ModelEventMessage_Event.
var (
	ModelEventMessage_Event_name = map[int32]string{
		0: "UNKNOWN_EVENT",
		1: "LOAD_FAIL_MEMORY",
		2: "LOADED",
		3: "LOAD_FAILED",
		4: "UNLOADED",
		5: "UNLOAD_FAILED",
		6: "REMOVED",
		7: "REMOVE_FAILED",
		8: "SCALE_UP_REQUEST",
		9: "RSYNC",
	}
	ModelEventMessage_Event_value = map[string]int32{
		"UNKNOWN_EVENT":    0,
		"LOAD_FAIL_MEMORY": 1,
		"LOADED":           2,
		"LOAD_FAILED":      3,
		"UNLOADED":         4,
		"UNLOAD_FAILED":    5,
		"REMOVED":          6,
		"REMOVE_FAILED":    7,
		"SCALE_UP_REQUEST": 8,
		"RSYNC":            9,
	}
)

func (x ModelEventMessage_Event) Enum() *ModelEventMessage_Event {
	p := new(ModelEventMessage_Event)
	*p = x
	return p
}

func (x ModelEventMessage_Event) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ModelEventMessage_Event) Descriptor() protoreflect.EnumDescriptor {
	return file_mlops_agent_agent_proto_enumTypes[0].Descriptor()
}

func (ModelEventMessage_Event) Type() protoreflect.EnumType {
	return &file_mlops_agent_agent_proto_enumTypes[0]
}

func (x ModelEventMessage_Event) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ModelEventMessage_Event.Descriptor instead.
func (ModelEventMessage_Event) EnumDescriptor() ([]byte, []int) {
	return file_mlops_agent_agent_proto_rawDescGZIP(), []int{0, 0}
}

type ModelOperationMessage_Operation int32

const (
	ModelOperationMessage_UNKNOWN_EVENT ModelOperationMessage_Operation = 0
	ModelOperationMessage_LOAD_MODEL    ModelOperationMessage_Operation = 1
	ModelOperationMessage_UNLOAD_MODEL  ModelOperationMessage_Operation = 2
)

// Enum value maps for ModelOperationMessage_Operation.
var (
	ModelOperationMessage_Operation_name = map[int32]string{
		0: "UNKNOWN_EVENT",
		1: "LOAD_MODEL",
		2: "UNLOAD_MODEL",
	}
	ModelOperationMessage_Operation_value = map[string]int32{
		"UNKNOWN_EVENT": 0,
		"LOAD_MODEL":    1,
		"UNLOAD_MODEL":  2,
	}
)

func (x ModelOperationMessage_Operation) Enum() *ModelOperationMessage_Operation {
	p := new(ModelOperationMessage_Operation)
	*p = x
	return p
}

func (x ModelOperationMessage_Operation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ModelOperationMessage_Operation) Descriptor() protoreflect.EnumDescriptor {
	return file_mlops_agent_agent_proto_enumTypes[1].Descriptor()
}

func (ModelOperationMessage_Operation) Type() protoreflect.EnumType {
	return &file_mlops_agent_agent_proto_enumTypes[1]
}

func (x ModelOperationMessage_Operation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ModelOperationMessage_Operation.Descriptor instead.
func (ModelOperationMessage_Operation) EnumDescriptor() ([]byte, []int) {
	return file_mlops_agent_agent_proto_rawDescGZIP(), []int{4, 0}
}

type ModelEventMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerName           string                  `protobuf:"bytes,1,opt,name=serverName,proto3" json:"serverName,omitempty"`
	ReplicaIdx           uint32                  `protobuf:"varint,2,opt,name=replicaIdx,proto3" json:"replicaIdx,omitempty"`
	ModelName            string                  `protobuf:"bytes,3,opt,name=modelName,proto3" json:"modelName,omitempty"`
	ModelVersion         uint32                  `protobuf:"varint,4,opt,name=modelVersion,proto3" json:"modelVersion,omitempty"`
	Event                ModelEventMessage_Event `protobuf:"varint,5,opt,name=event,proto3,enum=seldon.mlops.agent.ModelEventMessage_Event" json:"event,omitempty"`
	Message              string                  `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
	AvailableMemoryBytes uint64                  `protobuf:"varint,7,opt,name=availableMemoryBytes,proto3" json:"availableMemoryBytes,omitempty"`
}

func (x *ModelEventMessage) Reset() {
	*x = ModelEventMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mlops_agent_agent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelEventMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelEventMessage) ProtoMessage() {}

func (x *ModelEventMessage) ProtoReflect() protoreflect.Message {
	mi := &file_mlops_agent_agent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelEventMessage.ProtoReflect.Descriptor instead.
func (*ModelEventMessage) Descriptor() ([]byte, []int) {
	return file_mlops_agent_agent_proto_rawDescGZIP(), []int{0}
}

func (x *ModelEventMessage) GetServerName() string {
	if x != nil {
		return x.ServerName
	}
	return ""
}

func (x *ModelEventMessage) GetReplicaIdx() uint32 {
	if x != nil {
		return x.ReplicaIdx
	}
	return 0
}

func (x *ModelEventMessage) GetModelName() string {
	if x != nil {
		return x.ModelName
	}
	return ""
}

func (x *ModelEventMessage) GetModelVersion() uint32 {
	if x != nil {
		return x.ModelVersion
	}
	return 0
}

func (x *ModelEventMessage) GetEvent() ModelEventMessage_Event {
	if x != nil {
		return x.Event
	}
	return ModelEventMessage_UNKNOWN_EVENT
}

func (x *ModelEventMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ModelEventMessage) GetAvailableMemoryBytes() uint64 {
	if x != nil {
		return x.AvailableMemoryBytes
	}
	return 0
}

type ModelEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ModelEventResponse) Reset() {
	*x = ModelEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mlops_agent_agent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelEventResponse) ProtoMessage() {}

func (x *ModelEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mlops_agent_agent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelEventResponse.ProtoReflect.Descriptor instead.
func (*ModelEventResponse) Descriptor() ([]byte, []int) {
	return file_mlops_agent_agent_proto_rawDescGZIP(), []int{1}
}

type AgentSubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerName           string          `protobuf:"bytes,1,opt,name=serverName,proto3" json:"serverName,omitempty"`
	Shared               bool            `protobuf:"varint,2,opt,name=shared,proto3" json:"shared,omitempty"`
	ReplicaIdx           uint32          `protobuf:"varint,3,opt,name=replicaIdx,proto3" json:"replicaIdx,omitempty"`
	ReplicaConfig        *ReplicaConfig  `protobuf:"bytes,4,opt,name=replicaConfig,proto3" json:"replicaConfig,omitempty"`
	LoadedModels         []*ModelVersion `protobuf:"bytes,5,rep,name=loadedModels,proto3" json:"loadedModels,omitempty"`
	AvailableMemoryBytes uint64          `protobuf:"varint,6,opt,name=availableMemoryBytes,proto3" json:"availableMemoryBytes,omitempty"`
}

func (x *AgentSubscribeRequest) Reset() {
	*x = AgentSubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mlops_agent_agent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentSubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentSubscribeRequest) ProtoMessage() {}

func (x *AgentSubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mlops_agent_agent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentSubscribeRequest.ProtoReflect.Descriptor instead.
func (*AgentSubscribeRequest) Descriptor() ([]byte, []int) {
	return file_mlops_agent_agent_proto_rawDescGZIP(), []int{2}
}

func (x *AgentSubscribeRequest) GetServerName() string {
	if x != nil {
		return x.ServerName
	}
	return ""
}

func (x *AgentSubscribeRequest) GetShared() bool {
	if x != nil {
		return x.Shared
	}
	return false
}

func (x *AgentSubscribeRequest) GetReplicaIdx() uint32 {
	if x != nil {
		return x.ReplicaIdx
	}
	return 0
}

func (x *AgentSubscribeRequest) GetReplicaConfig() *ReplicaConfig {
	if x != nil {
		return x.ReplicaConfig
	}
	return nil
}

func (x *AgentSubscribeRequest) GetLoadedModels() []*ModelVersion {
	if x != nil {
		return x.LoadedModels
	}
	return nil
}

func (x *AgentSubscribeRequest) GetAvailableMemoryBytes() uint64 {
	if x != nil {
		return x.AvailableMemoryBytes
	}
	return 0
}

type ReplicaConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InferenceSvc         string   `protobuf:"bytes,1,opt,name=inferenceSvc,proto3" json:"inferenceSvc,omitempty"`                  // inference DNS service name
	InferenceHttpPort    int32    `protobuf:"varint,2,opt,name=inferenceHttpPort,proto3" json:"inferenceHttpPort,omitempty"`       // inference HTTP port
	InferenceGrpcPort    int32    `protobuf:"varint,3,opt,name=inferenceGrpcPort,proto3" json:"inferenceGrpcPort,omitempty"`       // Inference grpc port
	MemoryBytes          uint64   `protobuf:"varint,4,opt,name=memoryBytes,proto3" json:"memoryBytes,omitempty"`                   // The memory capacity of the server replica
	Capabilities         []string `protobuf:"bytes,5,rep,name=capabilities,proto3" json:"capabilities,omitempty"`                  // The list of capabilities of the server, e.g. sklearn, pytorch, xgboost, mlflow
	OverCommitPercentage uint32   `protobuf:"varint,6,opt,name=overCommitPercentage,proto3" json:"overCommitPercentage,omitempty"` // The percentage of over commit to allow, set to 0 (%) to disable over commit
}

func (x *ReplicaConfig) Reset() {
	*x = ReplicaConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mlops_agent_agent_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplicaConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplicaConfig) ProtoMessage() {}

func (x *ReplicaConfig) ProtoReflect() protoreflect.Message {
	mi := &file_mlops_agent_agent_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplicaConfig.ProtoReflect.Descriptor instead.
func (*ReplicaConfig) Descriptor() ([]byte, []int) {
	return file_mlops_agent_agent_proto_rawDescGZIP(), []int{3}
}

func (x *ReplicaConfig) GetInferenceSvc() string {
	if x != nil {
		return x.InferenceSvc
	}
	return ""
}

func (x *ReplicaConfig) GetInferenceHttpPort() int32 {
	if x != nil {
		return x.InferenceHttpPort
	}
	return 0
}

func (x *ReplicaConfig) GetInferenceGrpcPort() int32 {
	if x != nil {
		return x.InferenceGrpcPort
	}
	return 0
}

func (x *ReplicaConfig) GetMemoryBytes() uint64 {
	if x != nil {
		return x.MemoryBytes
	}
	return 0
}

func (x *ReplicaConfig) GetCapabilities() []string {
	if x != nil {
		return x.Capabilities
	}
	return nil
}

func (x *ReplicaConfig) GetOverCommitPercentage() uint32 {
	if x != nil {
		return x.OverCommitPercentage
	}
	return 0
}

type ModelOperationMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation    ModelOperationMessage_Operation `protobuf:"varint,1,opt,name=operation,proto3,enum=seldon.mlops.agent.ModelOperationMessage_Operation" json:"operation,omitempty"`
	ModelVersion *ModelVersion                   `protobuf:"bytes,2,opt,name=modelVersion,proto3" json:"modelVersion,omitempty"`
}

func (x *ModelOperationMessage) Reset() {
	*x = ModelOperationMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mlops_agent_agent_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelOperationMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelOperationMessage) ProtoMessage() {}

func (x *ModelOperationMessage) ProtoReflect() protoreflect.Message {
	mi := &file_mlops_agent_agent_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelOperationMessage.ProtoReflect.Descriptor instead.
func (*ModelOperationMessage) Descriptor() ([]byte, []int) {
	return file_mlops_agent_agent_proto_rawDescGZIP(), []int{4}
}

func (x *ModelOperationMessage) GetOperation() ModelOperationMessage_Operation {
	if x != nil {
		return x.Operation
	}
	return ModelOperationMessage_UNKNOWN_EVENT
}

func (x *ModelOperationMessage) GetModelVersion() *ModelVersion {
	if x != nil {
		return x.ModelVersion
	}
	return nil
}

type ModelVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Model   *scheduler.Model `protobuf:"bytes,1,opt,name=model,proto3" json:"model,omitempty"`
	Version uint32           `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ModelVersion) Reset() {
	*x = ModelVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mlops_agent_agent_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelVersion) ProtoMessage() {}

func (x *ModelVersion) ProtoReflect() protoreflect.Message {
	mi := &file_mlops_agent_agent_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelVersion.ProtoReflect.Descriptor instead.
func (*ModelVersion) Descriptor() ([]byte, []int) {
	return file_mlops_agent_agent_proto_rawDescGZIP(), []int{5}
}

func (x *ModelVersion) GetModel() *scheduler.Model {
	if x != nil {
		return x.Model
	}
	return nil
}

func (x *ModelVersion) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

var File_mlops_agent_agent_proto protoreflect.FileDescriptor

var file_mlops_agent_agent_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x6c, 0x6f, 0x70, 0x73, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x73, 0x65, 0x6c, 0x64, 0x6f,
	0x6e, 0x2e, 0x6d, 0x6c, 0x6f, 0x70, 0x73, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x1a, 0x1f, 0x6d,
	0x6c, 0x6f, 0x70, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd8,
	0x03, 0x0a, 0x11, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x49,
	0x64, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x49, 0x64, 0x78, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x41, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x73, 0x65, 0x6c, 0x64, 0x6f, 0x6e, 0x2e, 0x6d,
	0x6c, 0x6f, 0x70, 0x73, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x14, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x42, 0x79, 0x74, 0x65, 0x73, 0x22, 0xaf, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x56, 0x45,
	0x4e, 0x54, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x46, 0x41, 0x49,
	0x4c, 0x5f, 0x4d, 0x45, 0x4d, 0x4f, 0x52, 0x59, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x4f,
	0x41, 0x44, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x4e, 0x4c, 0x4f, 0x41,
	0x44, 0x45, 0x44, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4c, 0x4f, 0x41, 0x44, 0x5f,
	0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x4d, 0x4f,
	0x56, 0x45, 0x44, 0x10, 0x06, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x5f,
	0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x07, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x43, 0x41, 0x4c,
	0x45, 0x5f, 0x55, 0x50, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x08, 0x12, 0x09,
	0x0a, 0x05, 0x52, 0x53, 0x59, 0x4e, 0x43, 0x10, 0x09, 0x22, 0x14, 0x0a, 0x12, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0xb2, 0x02, 0x0a, 0x15, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x49, 0x64, 0x78, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x49, 0x64,
	0x78, 0x12, 0x47, 0x0a, 0x0d, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x65, 0x6c, 0x64, 0x6f,
	0x6e, 0x2e, 0x6d, 0x6c, 0x6f, 0x70, 0x73, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0d, 0x72, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x44, 0x0a, 0x0c, 0x6c, 0x6f,
	0x61, 0x64, 0x65, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x73, 0x65, 0x6c, 0x64, 0x6f, 0x6e, 0x2e, 0x6d, 0x6c, 0x6f, 0x70, 0x73, 0x2e,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x0c, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x12, 0x32, 0x0a, 0x14, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x14,
	0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x42,
	0x79, 0x74, 0x65, 0x73, 0x22, 0x89, 0x02, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x53, 0x76, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x76, 0x63, 0x12, 0x2c, 0x0a, 0x11, 0x69, 0x6e,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x48, 0x74, 0x74, 0x70, 0x50, 0x6f, 0x72, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x48, 0x74, 0x74, 0x70, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x69, 0x6e, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x47, 0x72, 0x70, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x11, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x47, 0x72,
	0x70, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x61, 0x70, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c,
	0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x14,
	0x6f, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x6f, 0x76, 0x65, 0x72,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65,
	0x22, 0xf2, 0x01, 0x0a, 0x15, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x51, 0x0a, 0x09, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e,
	0x73, 0x65, 0x6c, 0x64, 0x6f, 0x6e, 0x2e, 0x6d, 0x6c, 0x6f, 0x70, 0x73, 0x2e, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a,
	0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x65, 0x6c, 0x64, 0x6f, 0x6e, 0x2e, 0x6d, 0x6c, 0x6f,
	0x70, 0x73, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0x40, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x56, 0x45, 0x4e,
	0x54, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x4d, 0x4f, 0x44, 0x45,
	0x4c, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x4e, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x4d, 0x4f,
	0x44, 0x45, 0x4c, 0x10, 0x02, 0x22, 0x5d, 0x0a, 0x0c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x65, 0x6c, 0x64, 0x6f, 0x6e, 0x2e, 0x6d, 0x6c,
	0x6f, 0x70, 0x73, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x32, 0xd4, 0x01, 0x0a, 0x0c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a, 0x0a, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x25, 0x2e, 0x73, 0x65, 0x6c, 0x64, 0x6f, 0x6e, 0x2e, 0x6d, 0x6c, 0x6f,
	0x70, 0x73, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x26, 0x2e, 0x73, 0x65, 0x6c,
	0x64, 0x6f, 0x6e, 0x2e, 0x6d, 0x6c, 0x6f, 0x70, 0x73, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x65, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x12, 0x29, 0x2e, 0x73, 0x65, 0x6c, 0x64, 0x6f, 0x6e, 0x2e, 0x6d, 0x6c, 0x6f, 0x70, 0x73,
	0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x73,
	0x65, 0x6c, 0x64, 0x6f, 0x6e, 0x2e, 0x6d, 0x6c, 0x6f, 0x70, 0x73, 0x2e, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x3c, 0x5a, 0x3a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x65, 0x6c, 0x64, 0x6f, 0x6e,
	0x69, 0x6f, 0x2f, 0x73, 0x65, 0x6c, 0x64, 0x6f, 0x6e, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x6c,
	0x6f, 0x70, 0x73, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_mlops_agent_agent_proto_rawDescOnce sync.Once
	file_mlops_agent_agent_proto_rawDescData = file_mlops_agent_agent_proto_rawDesc
)

func file_mlops_agent_agent_proto_rawDescGZIP() []byte {
	file_mlops_agent_agent_proto_rawDescOnce.Do(func() {
		file_mlops_agent_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_mlops_agent_agent_proto_rawDescData)
	})
	return file_mlops_agent_agent_proto_rawDescData
}

var file_mlops_agent_agent_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_mlops_agent_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_mlops_agent_agent_proto_goTypes = []interface{}{
	(ModelEventMessage_Event)(0),         // 0: seldon.mlops.agent.ModelEventMessage.Event
	(ModelOperationMessage_Operation)(0), // 1: seldon.mlops.agent.ModelOperationMessage.Operation
	(*ModelEventMessage)(nil),            // 2: seldon.mlops.agent.ModelEventMessage
	(*ModelEventResponse)(nil),           // 3: seldon.mlops.agent.ModelEventResponse
	(*AgentSubscribeRequest)(nil),        // 4: seldon.mlops.agent.AgentSubscribeRequest
	(*ReplicaConfig)(nil),                // 5: seldon.mlops.agent.ReplicaConfig
	(*ModelOperationMessage)(nil),        // 6: seldon.mlops.agent.ModelOperationMessage
	(*ModelVersion)(nil),                 // 7: seldon.mlops.agent.ModelVersion
	(*scheduler.Model)(nil),              // 8: seldon.mlops.scheduler.Model
}
var file_mlops_agent_agent_proto_depIdxs = []int32{
	0, // 0: seldon.mlops.agent.ModelEventMessage.event:type_name -> seldon.mlops.agent.ModelEventMessage.Event
	5, // 1: seldon.mlops.agent.AgentSubscribeRequest.replicaConfig:type_name -> seldon.mlops.agent.ReplicaConfig
	7, // 2: seldon.mlops.agent.AgentSubscribeRequest.loadedModels:type_name -> seldon.mlops.agent.ModelVersion
	1, // 3: seldon.mlops.agent.ModelOperationMessage.operation:type_name -> seldon.mlops.agent.ModelOperationMessage.Operation
	7, // 4: seldon.mlops.agent.ModelOperationMessage.modelVersion:type_name -> seldon.mlops.agent.ModelVersion
	8, // 5: seldon.mlops.agent.ModelVersion.model:type_name -> seldon.mlops.scheduler.Model
	2, // 6: seldon.mlops.agent.AgentService.AgentEvent:input_type -> seldon.mlops.agent.ModelEventMessage
	4, // 7: seldon.mlops.agent.AgentService.Subscribe:input_type -> seldon.mlops.agent.AgentSubscribeRequest
	3, // 8: seldon.mlops.agent.AgentService.AgentEvent:output_type -> seldon.mlops.agent.ModelEventResponse
	6, // 9: seldon.mlops.agent.AgentService.Subscribe:output_type -> seldon.mlops.agent.ModelOperationMessage
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_mlops_agent_agent_proto_init() }
func file_mlops_agent_agent_proto_init() {
	if File_mlops_agent_agent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mlops_agent_agent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelEventMessage); i {
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
		file_mlops_agent_agent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelEventResponse); i {
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
		file_mlops_agent_agent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentSubscribeRequest); i {
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
		file_mlops_agent_agent_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplicaConfig); i {
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
		file_mlops_agent_agent_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelOperationMessage); i {
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
		file_mlops_agent_agent_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelVersion); i {
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
			RawDescriptor: file_mlops_agent_agent_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mlops_agent_agent_proto_goTypes,
		DependencyIndexes: file_mlops_agent_agent_proto_depIdxs,
		EnumInfos:         file_mlops_agent_agent_proto_enumTypes,
		MessageInfos:      file_mlops_agent_agent_proto_msgTypes,
	}.Build()
	File_mlops_agent_agent_proto = out.File
	file_mlops_agent_agent_proto_rawDesc = nil
	file_mlops_agent_agent_proto_goTypes = nil
	file_mlops_agent_agent_proto_depIdxs = nil
}