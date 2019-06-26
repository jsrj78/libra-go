// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events.proto

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

// An event emitted from a smart contract
type Event struct {
	AccessPath           *AccessPath `protobuf:"bytes,1,opt,name=access_path,json=accessPath,proto3" json:"access_path,omitempty"`
	SequenceNumber       uint64      `protobuf:"varint,2,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
	EventData            []byte      `protobuf:"bytes,3,opt,name=event_data,json=eventData,proto3" json:"event_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f22242cb04491f9, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetAccessPath() *AccessPath {
	if m != nil {
		return m.AccessPath
	}
	return nil
}

func (m *Event) GetSequenceNumber() uint64 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

func (m *Event) GetEventData() []byte {
	if m != nil {
		return m.EventData
	}
	return nil
}

// An event along with the proof for the event
type EventWithProof struct {
	TransactionVersion   uint64      `protobuf:"varint,1,opt,name=transaction_version,json=transactionVersion,proto3" json:"transaction_version,omitempty"`
	EventIndex           uint64      `protobuf:"varint,2,opt,name=event_index,json=eventIndex,proto3" json:"event_index,omitempty"`
	Event                *Event      `protobuf:"bytes,3,opt,name=event,proto3" json:"event,omitempty"`
	Proof                *EventProof `protobuf:"bytes,4,opt,name=proof,proto3" json:"proof,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *EventWithProof) Reset()         { *m = EventWithProof{} }
func (m *EventWithProof) String() string { return proto.CompactTextString(m) }
func (*EventWithProof) ProtoMessage()    {}
func (*EventWithProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f22242cb04491f9, []int{1}
}

func (m *EventWithProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventWithProof.Unmarshal(m, b)
}
func (m *EventWithProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventWithProof.Marshal(b, m, deterministic)
}
func (m *EventWithProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventWithProof.Merge(m, src)
}
func (m *EventWithProof) XXX_Size() int {
	return xxx_messageInfo_EventWithProof.Size(m)
}
func (m *EventWithProof) XXX_DiscardUnknown() {
	xxx_messageInfo_EventWithProof.DiscardUnknown(m)
}

var xxx_messageInfo_EventWithProof proto.InternalMessageInfo

func (m *EventWithProof) GetTransactionVersion() uint64 {
	if m != nil {
		return m.TransactionVersion
	}
	return 0
}

func (m *EventWithProof) GetEventIndex() uint64 {
	if m != nil {
		return m.EventIndex
	}
	return 0
}

func (m *EventWithProof) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *EventWithProof) GetProof() *EventProof {
	if m != nil {
		return m.Proof
	}
	return nil
}

// A list of events.
type EventsList struct {
	Events               []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventsList) Reset()         { *m = EventsList{} }
func (m *EventsList) String() string { return proto.CompactTextString(m) }
func (*EventsList) ProtoMessage()    {}
func (*EventsList) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f22242cb04491f9, []int{2}
}

func (m *EventsList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventsList.Unmarshal(m, b)
}
func (m *EventsList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventsList.Marshal(b, m, deterministic)
}
func (m *EventsList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventsList.Merge(m, src)
}
func (m *EventsList) XXX_Size() int {
	return xxx_messageInfo_EventsList.Size(m)
}
func (m *EventsList) XXX_DiscardUnknown() {
	xxx_messageInfo_EventsList.DiscardUnknown(m)
}

var xxx_messageInfo_EventsList proto.InternalMessageInfo

func (m *EventsList) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

// A list of EventList's, each representing all events for a transaction.
type EventsForVersions struct {
	EventsForVersion     []*EventsList `protobuf:"bytes,1,rep,name=events_for_version,json=eventsForVersion,proto3" json:"events_for_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *EventsForVersions) Reset()         { *m = EventsForVersions{} }
func (m *EventsForVersions) String() string { return proto.CompactTextString(m) }
func (*EventsForVersions) ProtoMessage()    {}
func (*EventsForVersions) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f22242cb04491f9, []int{3}
}

func (m *EventsForVersions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventsForVersions.Unmarshal(m, b)
}
func (m *EventsForVersions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventsForVersions.Marshal(b, m, deterministic)
}
func (m *EventsForVersions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventsForVersions.Merge(m, src)
}
func (m *EventsForVersions) XXX_Size() int {
	return xxx_messageInfo_EventsForVersions.Size(m)
}
func (m *EventsForVersions) XXX_DiscardUnknown() {
	xxx_messageInfo_EventsForVersions.DiscardUnknown(m)
}

var xxx_messageInfo_EventsForVersions proto.InternalMessageInfo

func (m *EventsForVersions) GetEventsForVersion() []*EventsList {
	if m != nil {
		return m.EventsForVersion
	}
	return nil
}

func init() {
	proto.RegisterType((*Event)(nil), "types.Event")
	proto.RegisterType((*EventWithProof)(nil), "types.EventWithProof")
	proto.RegisterType((*EventsList)(nil), "types.EventsList")
	proto.RegisterType((*EventsForVersions)(nil), "types.EventsForVersions")
}

func init() { proto.RegisterFile("events.proto", fileDescriptor_8f22242cb04491f9) }

var fileDescriptor_8f22242cb04491f9 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xc1, 0x6a, 0xdb, 0x40,
	0x10, 0x86, 0x51, 0x6d, 0x19, 0x3a, 0x32, 0x6e, 0xbd, 0xbd, 0x08, 0x43, 0xa9, 0x10, 0x05, 0xfb,
	0x50, 0x24, 0x50, 0x1f, 0xa0, 0x34, 0x24, 0x81, 0x40, 0x08, 0x46, 0x84, 0x04, 0x72, 0x11, 0xab,
	0xf5, 0xd8, 0x5a, 0xb0, 0x77, 0x95, 0xdd, 0xb5, 0x93, 0x9c, 0xf3, 0x40, 0x79, 0xc5, 0xe0, 0x59,
	0x25, 0x38, 0xce, 0x4d, 0xfb, 0x7f, 0x33, 0xf3, 0xff, 0x33, 0x82, 0x21, 0xee, 0x50, 0x39, 0x9b,
	0xb5, 0x46, 0x3b, 0xcd, 0x42, 0xf7, 0xd4, 0xa2, 0x9d, 0x8c, 0xb9, 0x10, 0x68, 0x6d, 0xd5, 0x72,
	0xd7, 0x78, 0x32, 0x89, 0x5a, 0xa3, 0xf5, 0xd2, 0x3f, 0xd2, 0xe7, 0x00, 0xc2, 0xb3, 0x7d, 0x1f,
	0x2b, 0x20, 0x3a, 0xa8, 0x8d, 0x83, 0x24, 0x98, 0x45, 0xc5, 0x38, 0xa3, 0x31, 0xd9, 0x7f, 0x22,
	0x73, 0xee, 0x9a, 0x12, 0xf8, 0xfb, 0x37, 0x9b, 0xc2, 0x37, 0x8b, 0xf7, 0x5b, 0x54, 0x02, 0x2b,
	0xb5, 0xdd, 0xd4, 0x68, 0xe2, 0x2f, 0x49, 0x30, 0xeb, 0x97, 0xa3, 0x37, 0xf9, 0x8a, 0x54, 0xf6,
	0x13, 0x80, 0xd2, 0x55, 0x0b, 0xee, 0x78, 0xdc, 0x4b, 0x82, 0xd9, 0xb0, 0xfc, 0x4a, 0xca, 0x29,
	0x77, 0x3c, 0x7d, 0x09, 0x60, 0x44, 0x29, 0x6e, 0xa5, 0x6b, 0xe6, 0xfb, 0x78, 0x2c, 0x87, 0x1f,
	0xce, 0x70, 0x65, 0xb9, 0x70, 0x52, 0xab, 0x6a, 0x87, 0xc6, 0x4a, 0xad, 0x28, 0x56, 0xbf, 0x64,
	0x07, 0xe8, 0xc6, 0x13, 0xf6, 0x0b, 0x22, 0x6f, 0x21, 0xd5, 0x02, 0x1f, 0xbb, 0x1c, 0xde, 0xf5,
	0x62, 0xaf, 0xb0, 0x14, 0x42, 0x7a, 0x91, 0x7d, 0x54, 0x0c, 0xbb, 0xd5, 0xc8, 0xb7, 0xf4, 0x88,
	0x4d, 0x21, 0xa4, 0xeb, 0xc4, 0xfd, 0x0f, 0xeb, 0x53, 0x0d, 0xe5, 0x2a, 0x3d, 0x4f, 0x0b, 0x00,
	0x12, 0xed, 0xa5, 0xb4, 0x8e, 0xfd, 0x86, 0x81, 0x3f, 0x7e, 0x1c, 0x24, 0xbd, 0x4f, 0xb3, 0x3b,
	0x96, 0x5e, 0xc3, 0xd8, 0xf7, 0x9c, 0x6b, 0xd3, 0xa5, 0xb6, 0xec, 0x1f, 0x30, 0x8f, 0xab, 0xa5,
	0x36, 0x07, 0x6b, 0xf6, 0x8e, 0xed, 0xc9, 0xa9, 0xfc, 0x8e, 0x47, 0x13, 0x4e, 0xb2, 0xbb, 0x3f,
	0x2b, 0xe9, 0x9a, 0x6d, 0x9d, 0x09, 0xbd, 0xc9, 0x85, 0x5e, 0xe0, 0x86, 0xef, 0xd0, 0x48, 0x91,
	0xaf, 0x65, 0x6d, 0xb8, 0x58, 0x4b, 0x54, 0x2e, 0x5f, 0xe9, 0x07, 0xc3, 0xdb, 0x16, 0x4d, 0x3d,
	0xa0, 0x1f, 0xff, 0xf7, 0x35, 0x00, 0x00, 0xff, 0xff, 0x21, 0x2d, 0x61, 0x64, 0x2f, 0x02, 0x00,
	0x00,
}
