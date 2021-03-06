// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/RangelReale/fproto-wrap-headers/headers.proto

/*
Package gwproto is a generated protocol buffer package.

It is generated from these files:
	github.com/RangelReale/fproto-wrap-headers/headers.proto

It has these top-level messages:
	Headers
*/
package gwproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Headers struct {
	Headers map[string]*Headers_Values `protobuf:"bytes,1,rep,name=headers" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Headers) Reset()                    { *m = Headers{} }
func (m *Headers) String() string            { return proto.CompactTextString(m) }
func (*Headers) ProtoMessage()               {}
func (*Headers) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Headers) GetHeaders() map[string]*Headers_Values {
	if m != nil {
		return m.Headers
	}
	return nil
}

type Headers_Values struct {
	Values []string `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
}

func (m *Headers_Values) Reset()                    { *m = Headers_Values{} }
func (m *Headers_Values) String() string            { return proto.CompactTextString(m) }
func (*Headers_Values) ProtoMessage()               {}
func (*Headers_Values) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Headers_Values) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*Headers)(nil), "fproto_wrap_headers.Headers")
	proto.RegisterType((*Headers_Values)(nil), "fproto_wrap_headers.Headers.Values")
}

func init() {
	proto.RegisterFile("github.com/RangelReale/fproto-wrap-headers/headers.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x48, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x0f, 0x4a, 0xcc, 0x4b, 0x4f, 0xcd, 0x09, 0x4a, 0x4d,
	0xcc, 0x49, 0xd5, 0x4f, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2d, 0x2f, 0x4a, 0x2c, 0xd0, 0xcd,
	0x48, 0x4d, 0x4c, 0x49, 0x2d, 0x2a, 0xd6, 0x87, 0xd2, 0x7a, 0x60, 0x29, 0x21, 0x61, 0x88, 0x92,
	0x78, 0x90, 0x92, 0x78, 0xa8, 0x94, 0xd2, 0x45, 0x46, 0x2e, 0x76, 0x0f, 0x08, 0x5b, 0xc8, 0x99,
	0x8b, 0x1d, 0x2a, 0x2c, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x6d, 0xa4, 0xa9, 0x87, 0x45, 0x8b, 0x9e,
	0x07, 0x2a, 0xed, 0x9a, 0x57, 0x52, 0x54, 0x19, 0x04, 0xd3, 0x29, 0xa5, 0xc0, 0xc5, 0x16, 0x96,
	0x98, 0x53, 0x9a, 0x5a, 0x2c, 0x24, 0xc6, 0xc5, 0x56, 0x06, 0x66, 0x81, 0x4d, 0xe3, 0x0c, 0x82,
	0xf2, 0xa4, 0xe2, 0xb9, 0x78, 0x90, 0xb5, 0x0a, 0x09, 0x70, 0x31, 0x67, 0xa7, 0x56, 0x4a, 0x30,
	0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0x98, 0x42, 0x96, 0x5c, 0xac, 0x60, 0xb5, 0x12, 0x4c, 0x0a,
	0x8c, 0x1a, 0xdc, 0x46, 0xca, 0x78, 0x9d, 0x01, 0xb1, 0x2d, 0x08, 0xa2, 0xc3, 0x8a, 0xc9, 0x82,
	0xd1, 0xc9, 0x3a, 0xca, 0x92, 0x84, 0x40, 0x4a, 0xcf, 0x07, 0x71, 0xf5, 0xd3, 0xcb, 0xc1, 0x72,
	0x49, 0x6c, 0x60, 0xca, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x2c, 0xe1, 0x45, 0xc8, 0x68, 0x01,
	0x00, 0x00,
}
