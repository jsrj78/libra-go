// Code generated by protoc-gen-go. DO NOT EDIT.
// source: admission_control.proto

package gowrapper

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Additional statuses that are possible from admission control in addition
// to VM statuses.
type AdmissionControlStatus int32

const (
	AdmissionControlStatus_Accepted    AdmissionControlStatus = 0
	AdmissionControlStatus_Blacklisted AdmissionControlStatus = 1
	AdmissionControlStatus_Rejected    AdmissionControlStatus = 2
)

var AdmissionControlStatus_name = map[int32]string{
	0: "Accepted",
	1: "Blacklisted",
	2: "Rejected",
}

var AdmissionControlStatus_value = map[string]int32{
	"Accepted":    0,
	"Blacklisted": 1,
	"Rejected":    2,
}

func (x AdmissionControlStatus) String() string {
	return proto.EnumName(AdmissionControlStatus_name, int32(x))
}

func (AdmissionControlStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d13d6ff9aac892b4, []int{0}
}

// -----------------------------------------------------------------------------
// ---------------- Submit transaction
// -----------------------------------------------------------------------------
// The request for transaction submission.
type SubmitTransactionRequest struct {
	// Transaction signed by wallet.
	SignedTxn            *SignedTransaction `protobuf:"bytes,1,opt,name=signed_txn,json=signedTxn,proto3" json:"signed_txn,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SubmitTransactionRequest) Reset()         { *m = SubmitTransactionRequest{} }
func (m *SubmitTransactionRequest) String() string { return proto.CompactTextString(m) }
func (*SubmitTransactionRequest) ProtoMessage()    {}
func (*SubmitTransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d13d6ff9aac892b4, []int{0}
}

func (m *SubmitTransactionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitTransactionRequest.Unmarshal(m, b)
}
func (m *SubmitTransactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitTransactionRequest.Marshal(b, m, deterministic)
}
func (m *SubmitTransactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitTransactionRequest.Merge(m, src)
}
func (m *SubmitTransactionRequest) XXX_Size() int {
	return xxx_messageInfo_SubmitTransactionRequest.Size(m)
}
func (m *SubmitTransactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitTransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitTransactionRequest proto.InternalMessageInfo

func (m *SubmitTransactionRequest) GetSignedTxn() *SignedTransaction {
	if m != nil {
		return m.SignedTxn
	}
	return nil
}

// The response for transaction submission.
//
// How does a client know if their transaction was included?
// A response from the transaction submission only means that the transaction
// was successfully added to mempool, but not that it is guaranteed to be
// included in the chain.  Each transaction should include an expiration time in
// the signed transaction.  Let's call this T0.  As a client, I submit my
// transaction to a validator. I now need to poll for the transaction I
// submitted.  I can use the query that takes my account and sequence number. If
// I receive back that the transaction is completed, I will verify the proofs to
// ensure that this is the transaction I expected.  If I receive a response that
// my transaction is not yet completed, I must check the latest timestamp in the
// ledgerInfo that I receive back from the query.  If this time is greater than
// T0, I can be certain that my transaction will never be included.  If this
// time is less than T0, I need to continue polling.
type SubmitTransactionResponse struct {
	// The status of a transaction submission can either be a VM status, or
	// some other admission control/mempool specific status e.g. Blacklisted.
	//
	// Types that are valid to be assigned to Status:
	//	*SubmitTransactionResponse_VmStatus
	//	*SubmitTransactionResponse_AcStatus
	//	*SubmitTransactionResponse_MempoolStatus
	Status isSubmitTransactionResponse_Status `protobuf_oneof:"status"`
	// Public key(id) of the validator that processed this transaction
	ValidatorId          []byte   `protobuf:"bytes,4,opt,name=validator_id,json=validatorId,proto3" json:"validator_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubmitTransactionResponse) Reset()         { *m = SubmitTransactionResponse{} }
func (m *SubmitTransactionResponse) String() string { return proto.CompactTextString(m) }
func (*SubmitTransactionResponse) ProtoMessage()    {}
func (*SubmitTransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d13d6ff9aac892b4, []int{1}
}

func (m *SubmitTransactionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitTransactionResponse.Unmarshal(m, b)
}
func (m *SubmitTransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitTransactionResponse.Marshal(b, m, deterministic)
}
func (m *SubmitTransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitTransactionResponse.Merge(m, src)
}
func (m *SubmitTransactionResponse) XXX_Size() int {
	return xxx_messageInfo_SubmitTransactionResponse.Size(m)
}
func (m *SubmitTransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitTransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitTransactionResponse proto.InternalMessageInfo

type isSubmitTransactionResponse_Status interface {
	isSubmitTransactionResponse_Status()
}

type SubmitTransactionResponse_VmStatus struct {
	VmStatus *VMStatus `protobuf:"bytes,1,opt,name=vm_status,json=vmStatus,proto3,oneof"`
}

type SubmitTransactionResponse_AcStatus struct {
	AcStatus AdmissionControlStatus `protobuf:"varint,2,opt,name=ac_status,json=acStatus,proto3,enum=admission_control.AdmissionControlStatus,oneof"`
}

type SubmitTransactionResponse_MempoolStatus struct {
	MempoolStatus MempoolAddTransactionStatus `protobuf:"varint,3,opt,name=mempool_status,json=mempoolStatus,proto3,enum=mempool.MempoolAddTransactionStatus,oneof"`
}

func (*SubmitTransactionResponse_VmStatus) isSubmitTransactionResponse_Status() {}

func (*SubmitTransactionResponse_AcStatus) isSubmitTransactionResponse_Status() {}

func (*SubmitTransactionResponse_MempoolStatus) isSubmitTransactionResponse_Status() {}

func (m *SubmitTransactionResponse) GetStatus() isSubmitTransactionResponse_Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *SubmitTransactionResponse) GetVmStatus() *VMStatus {
	if x, ok := m.GetStatus().(*SubmitTransactionResponse_VmStatus); ok {
		return x.VmStatus
	}
	return nil
}

func (m *SubmitTransactionResponse) GetAcStatus() AdmissionControlStatus {
	if x, ok := m.GetStatus().(*SubmitTransactionResponse_AcStatus); ok {
		return x.AcStatus
	}
	return AdmissionControlStatus_Accepted
}

func (m *SubmitTransactionResponse) GetMempoolStatus() MempoolAddTransactionStatus {
	if x, ok := m.GetStatus().(*SubmitTransactionResponse_MempoolStatus); ok {
		return x.MempoolStatus
	}
	return MempoolAddTransactionStatus_Valid
}

func (m *SubmitTransactionResponse) GetValidatorId() []byte {
	if m != nil {
		return m.ValidatorId
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SubmitTransactionResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SubmitTransactionResponse_VmStatus)(nil),
		(*SubmitTransactionResponse_AcStatus)(nil),
		(*SubmitTransactionResponse_MempoolStatus)(nil),
	}
}

func init() {
	proto.RegisterEnum("admission_control.AdmissionControlStatus", AdmissionControlStatus_name, AdmissionControlStatus_value)
	proto.RegisterType((*SubmitTransactionRequest)(nil), "admission_control.SubmitTransactionRequest")
	proto.RegisterType((*SubmitTransactionResponse)(nil), "admission_control.SubmitTransactionResponse")
}

func init() { proto.RegisterFile("admission_control.proto", fileDescriptor_d13d6ff9aac892b4) }

var fileDescriptor_d13d6ff9aac892b4 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x8f, 0x12, 0x4d,
	0x10, 0x87, 0x19, 0xde, 0x37, 0x1b, 0x68, 0x70, 0x59, 0x3a, 0x44, 0x47, 0x4e, 0x88, 0x1e, 0x50,
	0x37, 0x4d, 0x82, 0x07, 0xcf, 0x60, 0x4c, 0xd6, 0x64, 0xb9, 0x00, 0x7a, 0xf0, 0x32, 0x69, 0xba,
	0xcb, 0xd9, 0xd6, 0xe9, 0x3f, 0x76, 0x17, 0xb3, 0xeb, 0xc7, 0xf0, 0x53, 0xfa, 0x35, 0x0c, 0x4c,
	0xb3, 0xbb, 0x0a, 0x1a, 0x4f, 0x33, 0xf3, 0xab, 0x67, 0x9e, 0xea, 0xa9, 0x1a, 0xf2, 0x88, 0x4b,
	0xad, 0x42, 0x50, 0xd6, 0x64, 0xc2, 0x1a, 0xf4, 0xb6, 0x60, 0xce, 0x5b, 0xb4, 0xb4, 0x7b, 0x50,
	0xe8, 0xf7, 0x72, 0xc0, 0xec, 0x5a, 0xe1, 0x55, 0xe6, 0xbc, 0xb5, 0x9f, 0x2a, 0xb0, 0xdf, 0xd3,
	0xa0, 0x9d, 0xb5, 0x45, 0x16, 0x90, 0xe3, 0x26, 0xc4, 0xb4, 0x8b, 0x9e, 0x9b, 0xc0, 0x05, 0x2a,
	0x6b, 0x62, 0xd4, 0x29, 0x75, 0x06, 0xde, 0x5b, 0x1f, 0x99, 0xe1, 0x92, 0xa4, 0xcb, 0xcd, 0x5a,
	0x2b, 0x5c, 0xdd, 0xb1, 0x0b, 0xf8, 0xba, 0x81, 0x80, 0xf4, 0x35, 0x21, 0x41, 0xe5, 0x06, 0x64,
	0x86, 0x37, 0x26, 0x4d, 0x06, 0xc9, 0xa8, 0x35, 0x49, 0x19, 0x7e, 0x73, 0x10, 0xd8, 0x72, 0x57,
	0xb8, 0xff, 0x52, 0xb3, 0x62, 0x57, 0x37, 0x66, 0xf8, 0xbd, 0x4e, 0x1e, 0x1f, 0xb1, 0x06, 0x67,
	0x4d, 0x00, 0xca, 0x48, 0xb3, 0xd4, 0xf1, 0xa4, 0xd1, 0xda, 0x89, 0xd6, 0x0f, 0xf3, 0xe5, 0x2e,
	0xbe, 0xa8, 0x2d, 0x1a, 0xa5, 0xae, 0xee, 0xe9, 0x05, 0x69, 0x72, 0xb1, 0xe7, 0xeb, 0x83, 0x64,
	0x74, 0x3a, 0x79, 0xce, 0x0e, 0x47, 0x36, 0xdd, 0x27, 0x6f, 0xaa, 0xe0, 0xce, 0xc4, 0x45, 0x34,
	0xcd, 0xc9, 0xe9, 0xaf, 0x83, 0x4a, 0xff, 0xdb, 0xe9, 0x9e, 0xb1, 0x18, 0xb3, 0x79, 0x75, 0x9d,
	0xca, 0xfb, 0x9f, 0x76, 0x6b, 0x7a, 0x10, 0xb1, 0xa8, 0x7b, 0x42, 0xda, 0x25, 0x2f, 0x94, 0xe4,
	0x68, 0x7d, 0xa6, 0x64, 0xfa, 0xff, 0x20, 0x19, 0xb5, 0x17, 0xad, 0xdb, 0xec, 0x9d, 0x9c, 0x35,
	0xc8, 0x49, 0xd5, 0xe9, 0xc5, 0x5b, 0xf2, 0xf0, 0xf8, 0x09, 0x69, 0x9b, 0x34, 0xa6, 0x42, 0x80,
	0x43, 0x90, 0x67, 0x35, 0xda, 0x21, 0xad, 0x59, 0xc1, 0xc5, 0x97, 0x42, 0x85, 0x6d, 0x90, 0x6c,
	0xcb, 0x0b, 0xf8, 0x0c, 0x62, 0xfb, 0x54, 0x9f, 0xfc, 0x48, 0xc8, 0xd9, 0xef, 0x1e, 0xea, 0x48,
	0xf7, 0x60, 0xdc, 0xf4, 0xe5, 0x91, 0x19, 0xfd, 0x69, 0xd5, 0xfd, 0xf3, 0x7f, 0x83, 0xab, 0x0d,
	0x0e, 0x6b, 0x94, 0x93, 0xde, 0x7b, 0x27, 0x39, 0xc2, 0xca, 0x5e, 0x72, 0x84, 0x80, 0x97, 0x20,
	0x73, 0xf0, 0x74, 0x18, 0x17, 0x79, 0xac, 0xb8, 0xef, 0xf5, 0xf4, 0xaf, 0xcc, 0xbe, 0xc5, 0x8c,
	0x7d, 0x3c, 0xcf, 0x15, 0x5e, 0x6d, 0xd6, 0x4c, 0x58, 0x3d, 0x16, 0x56, 0x82, 0xe6, 0x25, 0x78,
	0x25, 0xc6, 0x85, 0x5a, 0x7b, 0x2e, 0x0a, 0x05, 0x06, 0xc7, 0xb9, 0xbd, 0xf6, 0xdc, 0x39, 0xf0,
	0xeb, 0x93, 0xdd, 0x0f, 0xfd, 0xea, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x18, 0x31, 0xf4,
	0x4e, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AdmissionControlClient is the client API for AdmissionControl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdmissionControlClient interface {
	// Public API to submit transaction to a validator.
	SubmitTransaction(ctx context.Context, in *SubmitTransactionRequest, opts ...grpc.CallOption) (*SubmitTransactionResponse, error)
	// This API is used to update the client to the latest ledger version and
	// optionally also request 1..n other pieces of data.  This allows for batch
	// queries.  All queries return proofs that a client should check to validate
	// the data. Note that if a client only wishes to update to the latest
	// LedgerInfo and receive the proof of this latest version, they can simply
	// omit the requested_items (or pass an empty list)
	UpdateToLatestLedger(ctx context.Context, in *UpdateToLatestLedgerRequest, opts ...grpc.CallOption) (*UpdateToLatestLedgerResponse, error)
}

type admissionControlClient struct {
	cc *grpc.ClientConn
}

func NewAdmissionControlClient(cc *grpc.ClientConn) AdmissionControlClient {
	return &admissionControlClient{cc}
}

func (c *admissionControlClient) SubmitTransaction(ctx context.Context, in *SubmitTransactionRequest, opts ...grpc.CallOption) (*SubmitTransactionResponse, error) {
	out := new(SubmitTransactionResponse)
	err := c.cc.Invoke(ctx, "/admission_control.AdmissionControl/SubmitTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *admissionControlClient) UpdateToLatestLedger(ctx context.Context, in *UpdateToLatestLedgerRequest, opts ...grpc.CallOption) (*UpdateToLatestLedgerResponse, error) {
	out := new(UpdateToLatestLedgerResponse)
	err := c.cc.Invoke(ctx, "/admission_control.AdmissionControl/UpdateToLatestLedger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdmissionControlServer is the server API for AdmissionControl service.
type AdmissionControlServer interface {
	// Public API to submit transaction to a validator.
	SubmitTransaction(context.Context, *SubmitTransactionRequest) (*SubmitTransactionResponse, error)
	// This API is used to update the client to the latest ledger version and
	// optionally also request 1..n other pieces of data.  This allows for batch
	// queries.  All queries return proofs that a client should check to validate
	// the data. Note that if a client only wishes to update to the latest
	// LedgerInfo and receive the proof of this latest version, they can simply
	// omit the requested_items (or pass an empty list)
	UpdateToLatestLedger(context.Context, *UpdateToLatestLedgerRequest) (*UpdateToLatestLedgerResponse, error)
}

// UnimplementedAdmissionControlServer can be embedded to have forward compatible implementations.
type UnimplementedAdmissionControlServer struct {
}

func (*UnimplementedAdmissionControlServer) SubmitTransaction(ctx context.Context, req *SubmitTransactionRequest) (*SubmitTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitTransaction not implemented")
}
func (*UnimplementedAdmissionControlServer) UpdateToLatestLedger(ctx context.Context, req *UpdateToLatestLedgerRequest) (*UpdateToLatestLedgerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateToLatestLedger not implemented")
}

func RegisterAdmissionControlServer(s *grpc.Server, srv AdmissionControlServer) {
	s.RegisterService(&_AdmissionControl_serviceDesc, srv)
}

func _AdmissionControl_SubmitTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdmissionControlServer).SubmitTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admission_control.AdmissionControl/SubmitTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdmissionControlServer).SubmitTransaction(ctx, req.(*SubmitTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdmissionControl_UpdateToLatestLedger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateToLatestLedgerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdmissionControlServer).UpdateToLatestLedger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admission_control.AdmissionControl/UpdateToLatestLedger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdmissionControlServer).UpdateToLatestLedger(ctx, req.(*UpdateToLatestLedgerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdmissionControl_serviceDesc = grpc.ServiceDesc{
	ServiceName: "admission_control.AdmissionControl",
	HandlerType: (*AdmissionControlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitTransaction",
			Handler:    _AdmissionControl_SubmitTransaction_Handler,
		},
		{
			MethodName: "UpdateToLatestLedger",
			Handler:    _AdmissionControl_UpdateToLatestLedger_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admission_control.proto",
}
