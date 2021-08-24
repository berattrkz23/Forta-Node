// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: batch.proto

package protocol

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

type SignedAlertBatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data      *AlertBatch `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Signature *Signature  `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *SignedAlertBatch) Reset() {
	*x = SignedAlertBatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_batch_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedAlertBatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedAlertBatch) ProtoMessage() {}

func (x *SignedAlertBatch) ProtoReflect() protoreflect.Message {
	mi := &file_batch_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedAlertBatch.ProtoReflect.Descriptor instead.
func (*SignedAlertBatch) Descriptor() ([]byte, []int) {
	return file_batch_proto_rawDescGZIP(), []int{0}
}

func (x *SignedAlertBatch) GetData() *AlertBatch {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SignedAlertBatch) GetSignature() *Signature {
	if x != nil {
		return x.Signature
	}
	return nil
}

type AlertBatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChainId     uint64           `protobuf:"varint,1,opt,name=chainId,proto3" json:"chainId,omitempty"`
	BlockStart  uint64           `protobuf:"varint,2,opt,name=blockStart,proto3" json:"blockStart,omitempty"`
	BlockEnd    uint64           `protobuf:"varint,3,opt,name=blockEnd,proto3" json:"blockEnd,omitempty"`
	AlertCount  uint32           `protobuf:"varint,4,opt,name=alertCount,proto3" json:"alertCount,omitempty"`
	MaxSeverity Finding_Severity `protobuf:"varint,5,opt,name=maxSeverity,proto3,enum=network.forta.Finding_Severity" json:"maxSeverity,omitempty"`
	Results     []*BlockResults  `protobuf:"bytes,6,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *AlertBatch) Reset() {
	*x = AlertBatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_batch_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertBatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertBatch) ProtoMessage() {}

func (x *AlertBatch) ProtoReflect() protoreflect.Message {
	mi := &file_batch_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertBatch.ProtoReflect.Descriptor instead.
func (*AlertBatch) Descriptor() ([]byte, []int) {
	return file_batch_proto_rawDescGZIP(), []int{1}
}

func (x *AlertBatch) GetChainId() uint64 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

func (x *AlertBatch) GetBlockStart() uint64 {
	if x != nil {
		return x.BlockStart
	}
	return 0
}

func (x *AlertBatch) GetBlockEnd() uint64 {
	if x != nil {
		return x.BlockEnd
	}
	return 0
}

func (x *AlertBatch) GetAlertCount() uint32 {
	if x != nil {
		return x.AlertCount
	}
	return 0
}

func (x *AlertBatch) GetMaxSeverity() Finding_Severity {
	if x != nil {
		return x.MaxSeverity
	}
	return Finding_UNKNOWN
}

func (x *AlertBatch) GetResults() []*BlockResults {
	if x != nil {
		return x.Results
	}
	return nil
}

type BlockResults struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Block        *Block                `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	Results      []*AgentAlerts        `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	Transactions []*TransactionResults `protobuf:"bytes,3,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *BlockResults) Reset() {
	*x = BlockResults{}
	if protoimpl.UnsafeEnabled {
		mi := &file_batch_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockResults) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockResults) ProtoMessage() {}

func (x *BlockResults) ProtoReflect() protoreflect.Message {
	mi := &file_batch_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockResults.ProtoReflect.Descriptor instead.
func (*BlockResults) Descriptor() ([]byte, []int) {
	return file_batch_proto_rawDescGZIP(), []int{2}
}

func (x *BlockResults) GetBlock() *Block {
	if x != nil {
		return x.Block
	}
	return nil
}

func (x *BlockResults) GetResults() []*AgentAlerts {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *BlockResults) GetTransactions() []*TransactionResults {
	if x != nil {
		return x.Transactions
	}
	return nil
}

type TransactionResults struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *TransactionEvent `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Results     []*AgentAlerts    `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *TransactionResults) Reset() {
	*x = TransactionResults{}
	if protoimpl.UnsafeEnabled {
		mi := &file_batch_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionResults) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionResults) ProtoMessage() {}

func (x *TransactionResults) ProtoReflect() protoreflect.Message {
	mi := &file_batch_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionResults.ProtoReflect.Descriptor instead.
func (*TransactionResults) Descriptor() ([]byte, []int) {
	return file_batch_proto_rawDescGZIP(), []int{3}
}

func (x *TransactionResults) GetTransaction() *TransactionEvent {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *TransactionResults) GetResults() []*AgentAlerts {
	if x != nil {
		return x.Results
	}
	return nil
}

type AgentAlerts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Agent  *AgentInfo     `protobuf:"bytes,1,opt,name=agent,proto3" json:"agent,omitempty"`
	Alerts []*SignedAlert `protobuf:"bytes,2,rep,name=alerts,proto3" json:"alerts,omitempty"`
}

func (x *AgentAlerts) Reset() {
	*x = AgentAlerts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_batch_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentAlerts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentAlerts) ProtoMessage() {}

func (x *AgentAlerts) ProtoReflect() protoreflect.Message {
	mi := &file_batch_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentAlerts.ProtoReflect.Descriptor instead.
func (*AgentAlerts) Descriptor() ([]byte, []int) {
	return file_batch_proto_rawDescGZIP(), []int{4}
}

func (x *AgentAlerts) GetAgent() *AgentInfo {
	if x != nil {
		return x.Agent
	}
	return nil
}

func (x *AgentAlerts) GetAlerts() []*SignedAlert {
	if x != nil {
		return x.Alerts
	}
	return nil
}

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockHash   string `protobuf:"bytes,1,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	BlockNumber uint64 `protobuf:"varint,2,opt,name=blockNumber,proto3" json:"blockNumber,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_batch_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_batch_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_batch_proto_rawDescGZIP(), []int{5}
}

func (x *Block) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

func (x *Block) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

var File_batch_proto protoreflect.FileDescriptor

var file_batch_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x66, 0x6f, 0x72, 0x74, 0x61, 0x1a, 0x0b, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x79, 0x0a, 0x10, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2e, 0x66, 0x6f, 0x72, 0x74, 0x61, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x36, 0x0a, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x66, 0x6f, 0x72, 0x74, 0x61, 0x2e, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x22, 0xfc, 0x01, 0x0a, 0x0a, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x45, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x45, 0x6e, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x41, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x53, 0x65, 0x76,
	0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x66, 0x6f, 0x72, 0x74, 0x61, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x52, 0x0b, 0x6d, 0x61,
	0x78, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x35, 0x0a, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x66, 0x6f, 0x72, 0x74, 0x61, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x22, 0xb7, 0x01, 0x0a, 0x0c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x12, 0x2a, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x66, 0x6f, 0x72, 0x74, 0x61,
	0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x34, 0x0a,
	0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x66, 0x6f, 0x72, 0x74, 0x61, 0x2e, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x12, 0x45, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2e, 0x66, 0x6f, 0x72, 0x74, 0x61, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x0c, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x8d, 0x01, 0x0a, 0x12, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x12, 0x41, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x2e, 0x66, 0x6f, 0x72, 0x74, 0x61, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e,
	0x66, 0x6f, 0x72, 0x74, 0x61, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x41, 0x6c, 0x65, 0x72, 0x74,
	0x73, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x71, 0x0a, 0x0b, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x12, 0x2e, 0x0a, 0x05, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2e, 0x66, 0x6f, 0x72, 0x74, 0x61, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x06, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2e, 0x66, 0x6f, 0x72, 0x74, 0x61, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x06, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x22, 0x47, 0x0a,
	0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48,
	0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x50, 0x00, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_batch_proto_rawDescOnce sync.Once
	file_batch_proto_rawDescData = file_batch_proto_rawDesc
)

func file_batch_proto_rawDescGZIP() []byte {
	file_batch_proto_rawDescOnce.Do(func() {
		file_batch_proto_rawDescData = protoimpl.X.CompressGZIP(file_batch_proto_rawDescData)
	})
	return file_batch_proto_rawDescData
}

var file_batch_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_batch_proto_goTypes = []interface{}{
	(*SignedAlertBatch)(nil),   // 0: network.forta.SignedAlertBatch
	(*AlertBatch)(nil),         // 1: network.forta.AlertBatch
	(*BlockResults)(nil),       // 2: network.forta.BlockResults
	(*TransactionResults)(nil), // 3: network.forta.TransactionResults
	(*AgentAlerts)(nil),        // 4: network.forta.AgentAlerts
	(*Block)(nil),              // 5: network.forta.Block
	(*Signature)(nil),          // 6: network.forta.Signature
	(Finding_Severity)(0),      // 7: network.forta.Finding.Severity
	(*TransactionEvent)(nil),   // 8: network.forta.TransactionEvent
	(*AgentInfo)(nil),          // 9: network.forta.AgentInfo
	(*SignedAlert)(nil),        // 10: network.forta.SignedAlert
}
var file_batch_proto_depIdxs = []int32{
	1,  // 0: network.forta.SignedAlertBatch.data:type_name -> network.forta.AlertBatch
	6,  // 1: network.forta.SignedAlertBatch.signature:type_name -> network.forta.Signature
	7,  // 2: network.forta.AlertBatch.maxSeverity:type_name -> network.forta.Finding.Severity
	2,  // 3: network.forta.AlertBatch.results:type_name -> network.forta.BlockResults
	5,  // 4: network.forta.BlockResults.block:type_name -> network.forta.Block
	4,  // 5: network.forta.BlockResults.results:type_name -> network.forta.AgentAlerts
	3,  // 6: network.forta.BlockResults.transactions:type_name -> network.forta.TransactionResults
	8,  // 7: network.forta.TransactionResults.transaction:type_name -> network.forta.TransactionEvent
	4,  // 8: network.forta.TransactionResults.results:type_name -> network.forta.AgentAlerts
	9,  // 9: network.forta.AgentAlerts.agent:type_name -> network.forta.AgentInfo
	10, // 10: network.forta.AgentAlerts.alerts:type_name -> network.forta.SignedAlert
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_batch_proto_init() }
func file_batch_proto_init() {
	if File_batch_proto != nil {
		return
	}
	file_query_proto_init()
	file_agent_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_batch_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedAlertBatch); i {
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
		file_batch_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertBatch); i {
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
		file_batch_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockResults); i {
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
		file_batch_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionResults); i {
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
		file_batch_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentAlerts); i {
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
		file_batch_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
			RawDescriptor: file_batch_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_batch_proto_goTypes,
		DependencyIndexes: file_batch_proto_depIdxs,
		MessageInfos:      file_batch_proto_msgTypes,
	}.Build()
	File_batch_proto = out.File
	file_batch_proto_rawDesc = nil
	file_batch_proto_goTypes = nil
	file_batch_proto_depIdxs = nil
}
