// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: broker/message.proto

package broker

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Message represents a message emitted by the Mainflux adapters layer.
type Message struct {
	Channel              string               `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	Subtopic             string               `protobuf:"bytes,2,opt,name=subtopic,proto3" json:"subtopic,omitempty"`
	Publisher            string               `protobuf:"bytes,3,opt,name=publisher,proto3" json:"publisher,omitempty"`
	Protocol             string               `protobuf:"bytes,4,opt,name=protocol,proto3" json:"protocol,omitempty"`
	ContentType          string               `protobuf:"bytes,5,opt,name=contentType,proto3" json:"contentType,omitempty"`
	Payload              []byte               `protobuf:"bytes,6,opt,name=payload,proto3" json:"payload,omitempty"`
	Created              *timestamp.Timestamp `protobuf:"bytes,7,opt,name=created,proto3" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_6357da820a7eacc2, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *Message) GetSubtopic() string {
	if m != nil {
		return m.Subtopic
	}
	return ""
}

func (m *Message) GetPublisher() string {
	if m != nil {
		return m.Publisher
	}
	return ""
}

func (m *Message) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *Message) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *Message) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Message) GetCreated() *timestamp.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "broker.Message")
}

func init() { proto.RegisterFile("broker/message.proto", fileDescriptor_6357da820a7eacc2) }

var fileDescriptor_6357da820a7eacc2 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xcd, 0x4e, 0xc4, 0x20,
	0x14, 0x85, 0x83, 0x3f, 0xad, 0xc3, 0xb8, 0x22, 0x2e, 0x6e, 0x1a, 0x13, 0x1b, 0x57, 0x5d, 0xd1,
	0x44, 0x7d, 0x0d, 0x37, 0xcd, 0xbc, 0x00, 0x30, 0xd7, 0x4e, 0x23, 0xe5, 0x12, 0xa0, 0x8b, 0x79,
	0x66, 0x5f, 0xc2, 0x0c, 0x04, 0x9d, 0xe5, 0xf9, 0xbe, 0x7b, 0x02, 0x87, 0x3f, 0xe9, 0x40, 0xdf,
	0x18, 0xc6, 0x15, 0x63, 0x54, 0x33, 0x4a, 0x1f, 0x28, 0x91, 0x68, 0x0a, 0xed, 0x5e, 0x66, 0xa2,
	0xd9, 0xe2, 0x98, 0xa9, 0xde, 0xbe, 0xc6, 0xb4, 0xac, 0x18, 0x93, 0x5a, 0x7d, 0x39, 0x7c, 0xfd,
	0x61, 0xbc, 0xfd, 0x2c, 0x55, 0x01, 0xbc, 0x35, 0x27, 0xe5, 0x1c, 0x5a, 0x60, 0x3d, 0x1b, 0x76,
	0x53, 0x8d, 0xa2, 0xe3, 0x0f, 0x71, 0xd3, 0x89, 0xfc, 0x62, 0xe0, 0x26, 0xab, 0xbf, 0x2c, 0x9e,
	0xf9, 0xce, 0x6f, 0xda, 0x2e, 0xf1, 0x84, 0x01, 0x6e, 0xb3, 0xfc, 0x07, 0x97, 0x66, 0x7e, 0xc8,
	0x90, 0x85, 0xbb, 0xd2, 0xac, 0x59, 0xf4, 0x7c, 0x6f, 0xc8, 0x25, 0x74, 0xe9, 0x70, 0xf6, 0x08,
	0xf7, 0x59, 0x5f, 0xa3, 0xcb, 0x8f, 0xbc, 0x3a, 0x5b, 0x52, 0x47, 0x68, 0x7a, 0x36, 0x3c, 0x4e,
	0x35, 0x8a, 0x0f, 0xde, 0x9a, 0x80, 0x2a, 0xe1, 0x11, 0xda, 0x9e, 0x0d, 0xfb, 0xb7, 0x4e, 0x96,
	0xa9, 0xb2, 0x4e, 0x95, 0x87, 0x3a, 0x75, 0xaa, 0xa7, 0xba, 0xc9, 0xf2, 0xfd, 0x37, 0x00, 0x00,
	0xff, 0xff, 0x39, 0x63, 0xf9, 0x79, 0x35, 0x01, 0x00, 0x00,
}
