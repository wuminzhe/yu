// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.13.0
// source: block.proto

package goproto

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

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompactBlock *CompactBlock `protobuf:"bytes,1,opt,name=CompactBlock,proto3" json:"CompactBlock,omitempty"`
	Txns         []*SignedTxn  `protobuf:"bytes,2,rep,name=Txns,proto3" json:"Txns,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[0]
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
	return file_block_proto_rawDescGZIP(), []int{0}
}

func (x *Block) GetCompactBlock() *CompactBlock {
	if x != nil {
		return x.CompactBlock
	}
	return nil
}

func (x *Block) GetTxns() []*SignedTxn {
	if x != nil {
		return x.Txns
	}
	return nil
}

type CompactBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header     *Header  `protobuf:"bytes,1,opt,name=Header,proto3" json:"Header,omitempty"`
	TxnsHashes [][]byte `protobuf:"bytes,2,rep,name=TxnsHashes,proto3" json:"TxnsHashes,omitempty"`
}

func (x *CompactBlock) Reset() {
	*x = CompactBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompactBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompactBlock) ProtoMessage() {}

func (x *CompactBlock) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompactBlock.ProtoReflect.Descriptor instead.
func (*CompactBlock) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{1}
}

func (x *CompactBlock) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *CompactBlock) GetTxnsHashes() [][]byte {
	if x != nil {
		return x.TxnsHashes
	}
	return nil
}

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash       []byte      `protobuf:"bytes,1,opt,name=Hash,proto3" json:"Hash,omitempty"`
	PrevHash   []byte      `protobuf:"bytes,2,opt,name=PrevHash,proto3" json:"PrevHash,omitempty"`
	Height     uint64      `protobuf:"varint,3,opt,name=Height,proto3" json:"Height,omitempty"`
	TxnRoot    []byte      `protobuf:"bytes,4,opt,name=TxnRoot,proto3" json:"TxnRoot,omitempty"`
	StateRoot  []byte      `protobuf:"bytes,5,opt,name=StateRoot,proto3" json:"StateRoot,omitempty"`
	Timestamp  uint64      `protobuf:"varint,6,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	PeerID     string      `protobuf:"bytes,7,opt,name=PeerID,proto3" json:"PeerID,omitempty"`
	LeiLimit   uint64      `protobuf:"varint,8,opt,name=LeiLimit,proto3" json:"LeiLimit,omitempty"`
	LeiUsed    uint64      `protobuf:"varint,9,opt,name=LeiUsed,proto3" json:"LeiUsed,omitempty"`
	Pubkey     []byte      `protobuf:"bytes,10,opt,name=Pubkey,proto3" json:"Pubkey,omitempty"`
	Signature  []byte      `protobuf:"bytes,11,opt,name=Signature,proto3" json:"Signature,omitempty"`
	Validators *Validators `protobuf:"bytes,12,opt,name=Validators,proto3" json:"Validators,omitempty"`
	Proof      *Proof      `protobuf:"bytes,13,opt,name=Proof,proto3" json:"Proof,omitempty"`
	PowInfo    *PowInfo    `protobuf:"bytes,14,opt,name=PowInfo,proto3" json:"PowInfo,omitempty"`
	Extra      []byte      `protobuf:"bytes,15,opt,name=Extra,proto3" json:"Extra,omitempty"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{2}
}

func (x *Header) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *Header) GetPrevHash() []byte {
	if x != nil {
		return x.PrevHash
	}
	return nil
}

func (x *Header) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Header) GetTxnRoot() []byte {
	if x != nil {
		return x.TxnRoot
	}
	return nil
}

func (x *Header) GetStateRoot() []byte {
	if x != nil {
		return x.StateRoot
	}
	return nil
}

func (x *Header) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Header) GetPeerID() string {
	if x != nil {
		return x.PeerID
	}
	return ""
}

func (x *Header) GetLeiLimit() uint64 {
	if x != nil {
		return x.LeiLimit
	}
	return 0
}

func (x *Header) GetLeiUsed() uint64 {
	if x != nil {
		return x.LeiUsed
	}
	return 0
}

func (x *Header) GetPubkey() []byte {
	if x != nil {
		return x.Pubkey
	}
	return nil
}

func (x *Header) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *Header) GetValidators() *Validators {
	if x != nil {
		return x.Validators
	}
	return nil
}

func (x *Header) GetProof() *Proof {
	if x != nil {
		return x.Proof
	}
	return nil
}

func (x *Header) GetPowInfo() *PowInfo {
	if x != nil {
		return x.PowInfo
	}
	return nil
}

func (x *Header) GetExtra() []byte {
	if x != nil {
		return x.Extra
	}
	return nil
}

type Validators struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Validators []*Validator `protobuf:"bytes,1,rep,name=Validators,proto3" json:"Validators,omitempty"`
}

func (x *Validators) Reset() {
	*x = Validators{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Validators) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Validators) ProtoMessage() {}

func (x *Validators) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Validators.ProtoReflect.Descriptor instead.
func (*Validators) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{3}
}

func (x *Validators) GetValidators() []*Validator {
	if x != nil {
		return x.Validators
	}
	return nil
}

type Validator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PubKey        []byte `protobuf:"bytes,1,opt,name=PubKey,proto3" json:"PubKey,omitempty"`
	ProposeWeight uint64 `protobuf:"varint,2,opt,name=ProposeWeight,proto3" json:"ProposeWeight,omitempty"`
	VoteWeight    uint64 `protobuf:"varint,3,opt,name=VoteWeight,proto3" json:"VoteWeight,omitempty"`
}

func (x *Validator) Reset() {
	*x = Validator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Validator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Validator) ProtoMessage() {}

func (x *Validator) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Validator.ProtoReflect.Descriptor instead.
func (*Validator) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{4}
}

func (x *Validator) GetPubKey() []byte {
	if x != nil {
		return x.PubKey
	}
	return nil
}

func (x *Validator) GetProposeWeight() uint64 {
	if x != nil {
		return x.ProposeWeight
	}
	return 0
}

func (x *Validator) GetVoteWeight() uint64 {
	if x != nil {
		return x.VoteWeight
	}
	return 0
}

type Proof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockHash []byte `protobuf:"bytes,1,opt,name=BlockHash,proto3" json:"BlockHash,omitempty"`
	Height    uint64 `protobuf:"varint,2,opt,name=Height,proto3" json:"Height,omitempty"`
	VrfProof  []byte `protobuf:"bytes,3,opt,name=VrfProof,proto3" json:"VrfProof,omitempty"`
}

func (x *Proof) Reset() {
	*x = Proof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Proof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Proof) ProtoMessage() {}

func (x *Proof) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Proof.ProtoReflect.Descriptor instead.
func (*Proof) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{5}
}

func (x *Proof) GetBlockHash() []byte {
	if x != nil {
		return x.BlockHash
	}
	return nil
}

func (x *Proof) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Proof) GetVrfProof() []byte {
	if x != nil {
		return x.VrfProof
	}
	return nil
}

type PowInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nonce      uint64 `protobuf:"varint,1,opt,name=Nonce,proto3" json:"Nonce,omitempty"`
	Difficulty uint64 `protobuf:"varint,2,opt,name=Difficulty,proto3" json:"Difficulty,omitempty"`
}

func (x *PowInfo) Reset() {
	*x = PowInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PowInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PowInfo) ProtoMessage() {}

func (x *PowInfo) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PowInfo.ProtoReflect.Descriptor instead.
func (*PowInfo) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{6}
}

func (x *PowInfo) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *PowInfo) GetDifficulty() uint64 {
	if x != nil {
		return x.Difficulty
	}
	return 0
}

var File_block_proto protoreflect.FileDescriptor

var file_block_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x74,
	0x78, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5a, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x12, 0x31, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63,
	0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x0c, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1e, 0x0a, 0x04, 0x54, 0x78, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x78, 0x6e, 0x52, 0x04,
	0x54, 0x78, 0x6e, 0x73, 0x22, 0x4f, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1f, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x78, 0x6e, 0x73, 0x48, 0x61, 0x73,
	0x68, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0a, 0x54, 0x78, 0x6e, 0x73, 0x48,
	0x61, 0x73, 0x68, 0x65, 0x73, 0x22, 0xaf, 0x03, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x65, 0x76, 0x48, 0x61, 0x73, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x50, 0x72, 0x65, 0x76, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x16, 0x0a, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x78, 0x6e, 0x52,
	0x6f, 0x6f, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x54, 0x78, 0x6e, 0x52, 0x6f,
	0x6f, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x16,
	0x0a, 0x06, 0x50, 0x65, 0x65, 0x72, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x50, 0x65, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x65, 0x69, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x4c, 0x65, 0x69, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4c, 0x65, 0x69, 0x55, 0x73, 0x65, 0x64, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x07, 0x4c, 0x65, 0x69, 0x55, 0x73, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x50, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x50, 0x75,
	0x62, 0x6b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x2b, 0x0a, 0x0a, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x73, 0x52, 0x0a, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x12,
	0x1c, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06,
	0x2e, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x05, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x22, 0x0a,
	0x07, 0x50, 0x6f, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x50, 0x6f, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x50, 0x6f, 0x77, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x78, 0x74, 0x72, 0x61, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x45, 0x78, 0x74, 0x72, 0x61, 0x22, 0x38, 0x0a, 0x0a, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x2a, 0x0a, 0x0a, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x0a, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x22, 0x69, 0x0a, 0x09, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06,
	0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x65, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x50,
	0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x56, 0x6f, 0x74, 0x65, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0a, 0x56, 0x6f, 0x74, 0x65, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x59, 0x0a, 0x05,
	0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61,
	0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48,
	0x61, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x56,
	0x72, 0x66, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x56,
	0x72, 0x66, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x22, 0x3f, 0x0a, 0x07, 0x50, 0x6f, 0x77, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x69, 0x66, 0x66,
	0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x44, 0x69,
	0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x67, 0x6f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_block_proto_rawDescOnce sync.Once
	file_block_proto_rawDescData = file_block_proto_rawDesc
)

func file_block_proto_rawDescGZIP() []byte {
	file_block_proto_rawDescOnce.Do(func() {
		file_block_proto_rawDescData = protoimpl.X.CompressGZIP(file_block_proto_rawDescData)
	})
	return file_block_proto_rawDescData
}

var file_block_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_block_proto_goTypes = []interface{}{
	(*Block)(nil),        // 0: Block
	(*CompactBlock)(nil), // 1: CompactBlock
	(*Header)(nil),       // 2: Header
	(*Validators)(nil),   // 3: Validators
	(*Validator)(nil),    // 4: Validator
	(*Proof)(nil),        // 5: Proof
	(*PowInfo)(nil),      // 6: PowInfo
	(*SignedTxn)(nil),    // 7: SignedTxn
}
var file_block_proto_depIdxs = []int32{
	1, // 0: Block.CompactBlock:type_name -> CompactBlock
	7, // 1: Block.Txns:type_name -> SignedTxn
	2, // 2: CompactBlock.Header:type_name -> Header
	3, // 3: Header.Validators:type_name -> Validators
	5, // 4: Header.Proof:type_name -> Proof
	6, // 5: Header.PowInfo:type_name -> PowInfo
	4, // 6: Validators.Validators:type_name -> Validator
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_block_proto_init() }
func file_block_proto_init() {
	if File_block_proto != nil {
		return
	}
	file_txn_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_block_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_block_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompactBlock); i {
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
		file_block_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
		file_block_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Validators); i {
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
		file_block_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Validator); i {
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
		file_block_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Proof); i {
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
		file_block_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PowInfo); i {
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
			RawDescriptor: file_block_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_block_proto_goTypes,
		DependencyIndexes: file_block_proto_depIdxs,
		MessageInfos:      file_block_proto_msgTypes,
	}.Build()
	File_block_proto = out.File
	file_block_proto_rawDesc = nil
	file_block_proto_goTypes = nil
	file_block_proto_depIdxs = nil
}
