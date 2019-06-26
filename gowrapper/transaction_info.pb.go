// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transaction_info.proto

package gowrapper

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// `TransactionInfo` is the object we store in the transaction accumulator. It
// consists of the transaction as well as the execution result of this
// transaction. This are later returned to the client so that a client can
// validate the tree
type TransactionInfo struct {
	// Hash of the signed transaction that is stored
	SignedTransactionHash []byte `protobuf:"bytes,1,opt,name=signed_transaction_hash,json=signedTransactionHash,proto3" json:"signed_transaction_hash,omitempty"`
	// The root hash of Sparse Merkle Tree describing the world state at the end
	// of this transaction
	StateRootHash []byte `protobuf:"bytes,2,opt,name=state_root_hash,json=stateRootHash,proto3" json:"state_root_hash,omitempty"`
	// The root hash of Merkle Accumulator storing all events emitted during this
	// transaction.
	EventRootHash []byte `protobuf:"bytes,3,opt,name=event_root_hash,json=eventRootHash,proto3" json:"event_root_hash,omitempty"`
	// The amount of gas used by this transaction.
	GasUsed              uint64   `protobuf:"varint,4,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionInfo) Reset()         { *m = TransactionInfo{} }
func (m *TransactionInfo) String() string { return proto.CompactTextString(m) }
func (*TransactionInfo) ProtoMessage()    {}
func (*TransactionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_19ed5f4c459e84f4, []int{0}
}

func (m *TransactionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionInfo.Unmarshal(m, b)
}
func (m *TransactionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionInfo.Marshal(b, m, deterministic)
}
func (m *TransactionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionInfo.Merge(m, src)
}
func (m *TransactionInfo) XXX_Size() int {
	return xxx_messageInfo_TransactionInfo.Size(m)
}
func (m *TransactionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionInfo proto.InternalMessageInfo

func (m *TransactionInfo) GetSignedTransactionHash() []byte {
	if m != nil {
		return m.SignedTransactionHash
	}
	return nil
}

func (m *TransactionInfo) GetStateRootHash() []byte {
	if m != nil {
		return m.StateRootHash
	}
	return nil
}

func (m *TransactionInfo) GetEventRootHash() []byte {
	if m != nil {
		return m.EventRootHash
	}
	return nil
}

func (m *TransactionInfo) GetGasUsed() uint64 {
	if m != nil {
		return m.GasUsed
	}
	return 0
}

func init() {
	proto.RegisterType((*TransactionInfo)(nil), "types.TransactionInfo")
}

func init() { proto.RegisterFile("transaction_info.proto", fileDescriptor_19ed5f4c459e84f4) }

var fileDescriptor_19ed5f4c459e84f4 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xcf, 0xbf, 0x4a, 0x04, 0x31,
	0x10, 0xc7, 0x71, 0xa2, 0xe7, 0x1f, 0x82, 0x72, 0xb0, 0xa0, 0x9e, 0xdd, 0x61, 0x21, 0x57, 0xc8,
	0x6e, 0x21, 0xf8, 0x00, 0x56, 0xda, 0x2e, 0xda, 0xd8, 0x84, 0xd9, 0x64, 0x2e, 0x09, 0xdc, 0x65,
	0x42, 0x66, 0x6e, 0xc5, 0xf7, 0xf2, 0x01, 0x85, 0x2c, 0xea, 0xb6, 0xbf, 0xf9, 0x7c, 0x8b, 0xd1,
	0xd7, 0x52, 0x20, 0x31, 0x58, 0x89, 0x94, 0x4c, 0x4c, 0x5b, 0x6a, 0x73, 0x21, 0xa1, 0xe6, 0x44,
	0xbe, 0x32, 0xf2, 0xdd, 0xb7, 0xd2, 0xcb, 0xb7, 0x7f, 0xf1, 0x9a, 0xb6, 0xd4, 0x3c, 0xe9, 0x1b,
	0x8e, 0x3e, 0xa1, 0x33, 0xf3, 0x36, 0x00, 0x87, 0x95, 0x5a, 0xab, 0xcd, 0x45, 0x7f, 0x35, 0x9d,
	0x67, 0xdd, 0x0b, 0x70, 0x68, 0xee, 0xf5, 0x92, 0x05, 0x04, 0x4d, 0x21, 0x92, 0xc9, 0x1f, 0x55,
	0x7f, 0x59, 0xe7, 0x9e, 0x48, 0x7e, 0x1d, 0x8e, 0x98, 0x64, 0xe6, 0x8e, 0x27, 0x57, 0xe7, 0x3f,
	0x77, 0xab, 0xcf, 0x3d, 0xb0, 0x39, 0x30, 0xba, 0xd5, 0x62, 0xad, 0x36, 0x8b, 0xfe, 0xcc, 0x03,
	0xbf, 0x33, 0xba, 0xe7, 0xf6, 0xe3, 0xc1, 0x47, 0x09, 0x87, 0xa1, 0xb5, 0xb4, 0xef, 0x2c, 0x39,
	0xdc, 0xc3, 0x88, 0x25, 0xda, 0x6e, 0x17, 0x87, 0x02, 0x76, 0x17, 0x31, 0x49, 0xe7, 0xe9, 0xb3,
	0x40, 0xce, 0x58, 0x86, 0xd3, 0xfa, 0xf4, 0xe3, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x38,
	0x29, 0xa2, 0x0e, 0x01, 0x00, 0x00,
}
