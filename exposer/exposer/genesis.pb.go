// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        (unknown)
// source: exposer/genesis.proto

package exposer

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/anypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GenesisState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GenesisTime     *timestamppb.Timestamp        `protobuf:"bytes,1,opt,name=genesis_time,proto3" json:"genesis_time,omitempty"`
	ChainId         string                        `protobuf:"bytes,2,opt,name=chain_id,proto3" json:"chain_id,omitempty"`
	InitialHeight   string                        `protobuf:"bytes,3,opt,name=initial_height,proto3" json:"initial_height,omitempty"`
	ConsensusParams *GenesisState_ConsensusParams `protobuf:"bytes,4,opt,name=consensus_params,proto3" json:"consensus_params,omitempty"`
	AppHash         string                        `protobuf:"bytes,5,opt,name=app_hash,proto3" json:"app_hash,omitempty"`
	AppState        *structpb.Struct              `protobuf:"bytes,6,opt,name=app_state,proto3" json:"app_state,omitempty"`
}

func (x *GenesisState) Reset() {
	*x = GenesisState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exposer_genesis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenesisState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenesisState) ProtoMessage() {}

func (x *GenesisState) ProtoReflect() protoreflect.Message {
	mi := &file_exposer_genesis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenesisState.ProtoReflect.Descriptor instead.
func (*GenesisState) Descriptor() ([]byte, []int) {
	return file_exposer_genesis_proto_rawDescGZIP(), []int{0}
}

func (x *GenesisState) GetGenesisTime() *timestamppb.Timestamp {
	if x != nil {
		return x.GenesisTime
	}
	return nil
}

func (x *GenesisState) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

func (x *GenesisState) GetInitialHeight() string {
	if x != nil {
		return x.InitialHeight
	}
	return ""
}

func (x *GenesisState) GetConsensusParams() *GenesisState_ConsensusParams {
	if x != nil {
		return x.ConsensusParams
	}
	return nil
}

func (x *GenesisState) GetAppHash() string {
	if x != nil {
		return x.AppHash
	}
	return ""
}

func (x *GenesisState) GetAppState() *structpb.Struct {
	if x != nil {
		return x.AppState
	}
	return nil
}

type GenesisState_Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxBytes   string `protobuf:"bytes,1,opt,name=max_bytes,proto3" json:"max_bytes,omitempty"`
	MaxGas     string `protobuf:"bytes,2,opt,name=max_gas,proto3" json:"max_gas,omitempty"`
	TimeIotaMs string `protobuf:"bytes,3,opt,name=time_iota_ms,proto3" json:"time_iota_ms,omitempty"`
}

func (x *GenesisState_Block) Reset() {
	*x = GenesisState_Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exposer_genesis_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenesisState_Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenesisState_Block) ProtoMessage() {}

func (x *GenesisState_Block) ProtoReflect() protoreflect.Message {
	mi := &file_exposer_genesis_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenesisState_Block.ProtoReflect.Descriptor instead.
func (*GenesisState_Block) Descriptor() ([]byte, []int) {
	return file_exposer_genesis_proto_rawDescGZIP(), []int{0, 0}
}

func (x *GenesisState_Block) GetMaxBytes() string {
	if x != nil {
		return x.MaxBytes
	}
	return ""
}

func (x *GenesisState_Block) GetMaxGas() string {
	if x != nil {
		return x.MaxGas
	}
	return ""
}

func (x *GenesisState_Block) GetTimeIotaMs() string {
	if x != nil {
		return x.TimeIotaMs
	}
	return ""
}

type GenesisState_Evidence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxAgeNumBlocks string `protobuf:"bytes,1,opt,name=max_age_num_blocks,proto3" json:"max_age_num_blocks,omitempty"`
	MaxAgeDuration  string `protobuf:"bytes,2,opt,name=max_age_duration,proto3" json:"max_age_duration,omitempty"`
	MaxBytes        string `protobuf:"bytes,3,opt,name=max_bytes,proto3" json:"max_bytes,omitempty"`
}

func (x *GenesisState_Evidence) Reset() {
	*x = GenesisState_Evidence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exposer_genesis_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenesisState_Evidence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenesisState_Evidence) ProtoMessage() {}

func (x *GenesisState_Evidence) ProtoReflect() protoreflect.Message {
	mi := &file_exposer_genesis_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenesisState_Evidence.ProtoReflect.Descriptor instead.
func (*GenesisState_Evidence) Descriptor() ([]byte, []int) {
	return file_exposer_genesis_proto_rawDescGZIP(), []int{0, 1}
}

func (x *GenesisState_Evidence) GetMaxAgeNumBlocks() string {
	if x != nil {
		return x.MaxAgeNumBlocks
	}
	return ""
}

func (x *GenesisState_Evidence) GetMaxAgeDuration() string {
	if x != nil {
		return x.MaxAgeDuration
	}
	return ""
}

func (x *GenesisState_Evidence) GetMaxBytes() string {
	if x != nil {
		return x.MaxBytes
	}
	return ""
}

type GenesisState_Validator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PubKeyTypes []string `protobuf:"bytes,1,rep,name=pub_key_types,proto3" json:"pub_key_types,omitempty"`
}

func (x *GenesisState_Validator) Reset() {
	*x = GenesisState_Validator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exposer_genesis_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenesisState_Validator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenesisState_Validator) ProtoMessage() {}

func (x *GenesisState_Validator) ProtoReflect() protoreflect.Message {
	mi := &file_exposer_genesis_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenesisState_Validator.ProtoReflect.Descriptor instead.
func (*GenesisState_Validator) Descriptor() ([]byte, []int) {
	return file_exposer_genesis_proto_rawDescGZIP(), []int{0, 2}
}

func (x *GenesisState_Validator) GetPubKeyTypes() []string {
	if x != nil {
		return x.PubKeyTypes
	}
	return nil
}

type GenesisState_Version struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GenesisState_Version) Reset() {
	*x = GenesisState_Version{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exposer_genesis_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenesisState_Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenesisState_Version) ProtoMessage() {}

func (x *GenesisState_Version) ProtoReflect() protoreflect.Message {
	mi := &file_exposer_genesis_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenesisState_Version.ProtoReflect.Descriptor instead.
func (*GenesisState_Version) Descriptor() ([]byte, []int) {
	return file_exposer_genesis_proto_rawDescGZIP(), []int{0, 3}
}

type GenesisState_ConsensusParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Block     *GenesisState_Block     `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	Evidence  *GenesisState_Evidence  `protobuf:"bytes,2,opt,name=evidence,proto3" json:"evidence,omitempty"`
	Validator *GenesisState_Validator `protobuf:"bytes,3,opt,name=validator,proto3" json:"validator,omitempty"`
	Version   *GenesisState_Version   `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *GenesisState_ConsensusParams) Reset() {
	*x = GenesisState_ConsensusParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exposer_genesis_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenesisState_ConsensusParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenesisState_ConsensusParams) ProtoMessage() {}

func (x *GenesisState_ConsensusParams) ProtoReflect() protoreflect.Message {
	mi := &file_exposer_genesis_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenesisState_ConsensusParams.ProtoReflect.Descriptor instead.
func (*GenesisState_ConsensusParams) Descriptor() ([]byte, []int) {
	return file_exposer_genesis_proto_rawDescGZIP(), []int{0, 4}
}

func (x *GenesisState_ConsensusParams) GetBlock() *GenesisState_Block {
	if x != nil {
		return x.Block
	}
	return nil
}

func (x *GenesisState_ConsensusParams) GetEvidence() *GenesisState_Evidence {
	if x != nil {
		return x.Evidence
	}
	return nil
}

func (x *GenesisState_ConsensusParams) GetValidator() *GenesisState_Validator {
	if x != nil {
		return x.Validator
	}
	return nil
}

func (x *GenesisState_ConsensusParams) GetVersion() *GenesisState_Version {
	if x != nil {
		return x.Version
	}
	return nil
}

var File_exposer_genesis_proto protoreflect.FileDescriptor

var file_exposer_genesis_proto_rawDesc = []byte{
	0x0a, 0x15, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb7, 0x06, 0x0a, 0x0c, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x12, 0x26,
	0x0a, 0x0e, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x4a, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e,
	0x73, 0x75, 0x73, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e,
	0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x52, 0x10, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x5f, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x12, 0x35,
	0x0a, 0x09, 0x61, 0x70, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x09, 0x61, 0x70, 0x70, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x1a, 0x63, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1c,
	0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x61, 0x78, 0x5f, 0x67, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x61, 0x78, 0x5f, 0x67, 0x61, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x69,
	0x6f, 0x74, 0x61, 0x5f, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x69, 0x6f, 0x74, 0x61, 0x5f, 0x6d, 0x73, 0x1a, 0x84, 0x01, 0x0a, 0x08, 0x45,
	0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x6d, 0x61, 0x78, 0x5f, 0x61,
	0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x12, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d,
	0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x6d, 0x61, 0x78, 0x5f, 0x61,
	0x67, 0x65, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x1a, 0x31, 0x0a, 0x09, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x24,
	0x0a, 0x0d, 0x70, 0x75, 0x62, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x75, 0x62, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x1a, 0x09, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x1a,
	0xd9, 0x01, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x5f, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x12, 0x29, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12,
	0x32, 0x0a, 0x08, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x2e, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x08, 0x65, 0x76, 0x69, 0x64, 0x65,
	0x6e, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x52,
	0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x2f, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x47, 0x65,
	0x6e, 0x65, 0x73, 0x69, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x37, 0x42, 0x0c, 0x47,
	0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x25, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6e, 0x6d, 0x6f, 0x6c, 0x31,
	0x36, 0x39, 0x36, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70, 0x2f, 0x65, 0x78, 0x70,
	0x6f, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_exposer_genesis_proto_rawDescOnce sync.Once
	file_exposer_genesis_proto_rawDescData = file_exposer_genesis_proto_rawDesc
)

func file_exposer_genesis_proto_rawDescGZIP() []byte {
	file_exposer_genesis_proto_rawDescOnce.Do(func() {
		file_exposer_genesis_proto_rawDescData = protoimpl.X.CompressGZIP(file_exposer_genesis_proto_rawDescData)
	})
	return file_exposer_genesis_proto_rawDescData
}

var file_exposer_genesis_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_exposer_genesis_proto_goTypes = []interface{}{
	(*GenesisState)(nil),                 // 0: GenesisState
	(*GenesisState_Block)(nil),           // 1: GenesisState.Block
	(*GenesisState_Evidence)(nil),        // 2: GenesisState.Evidence
	(*GenesisState_Validator)(nil),       // 3: GenesisState.Validator
	(*GenesisState_Version)(nil),         // 4: GenesisState.Version
	(*GenesisState_ConsensusParams)(nil), // 5: GenesisState.Consensus_params
	(*timestamppb.Timestamp)(nil),        // 6: google.protobuf.Timestamp
	(*structpb.Struct)(nil),              // 7: google.protobuf.Struct
}
var file_exposer_genesis_proto_depIdxs = []int32{
	6, // 0: GenesisState.genesis_time:type_name -> google.protobuf.Timestamp
	5, // 1: GenesisState.consensus_params:type_name -> GenesisState.Consensus_params
	7, // 2: GenesisState.app_state:type_name -> google.protobuf.Struct
	1, // 3: GenesisState.Consensus_params.block:type_name -> GenesisState.Block
	2, // 4: GenesisState.Consensus_params.evidence:type_name -> GenesisState.Evidence
	3, // 5: GenesisState.Consensus_params.validator:type_name -> GenesisState.Validator
	4, // 6: GenesisState.Consensus_params.version:type_name -> GenesisState.Version
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_exposer_genesis_proto_init() }
func file_exposer_genesis_proto_init() {
	if File_exposer_genesis_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_exposer_genesis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenesisState); i {
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
		file_exposer_genesis_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenesisState_Block); i {
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
		file_exposer_genesis_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenesisState_Evidence); i {
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
		file_exposer_genesis_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenesisState_Validator); i {
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
		file_exposer_genesis_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenesisState_Version); i {
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
		file_exposer_genesis_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenesisState_ConsensusParams); i {
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
			RawDescriptor: file_exposer_genesis_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_exposer_genesis_proto_goTypes,
		DependencyIndexes: file_exposer_genesis_proto_depIdxs,
		MessageInfos:      file_exposer_genesis_proto_msgTypes,
	}.Build()
	File_exposer_genesis_proto = out.File
	file_exposer_genesis_proto_rawDesc = nil
	file_exposer_genesis_proto_goTypes = nil
	file_exposer_genesis_proto_depIdxs = nil
}
