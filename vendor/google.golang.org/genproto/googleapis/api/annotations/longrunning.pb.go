// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/longrunning.proto

package annotations // import "google.golang.org/genproto/googleapis/api/annotations"

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

// A message representing types returned by a longrunning operation.
type OperationData struct {
	// Required. The message name of the primary return type for this
	// long-running operation.
	// This type will be used to deserialize the LRO's response.
	//
	// If the response is in a different package from the rpc, a fully-qualified
	// message name must be used (e.g. "google.protobuf.Empty").
	ResponseType string `protobuf:"bytes,1,opt,name=response_type,json=responseType,proto3" json:"response_type,omitempty"`
	// Required. The metadata message name for this long-running operation.
	MetadataType         string   `protobuf:"bytes,2,opt,name=metadata_type,json=metadataType,proto3" json:"metadata_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperationData) Reset()         { *m = OperationData{} }
func (m *OperationData) String() string { return proto.CompactTextString(m) }
func (*OperationData) ProtoMessage()    {}
func (*OperationData) Descriptor() ([]byte, []int) {
	return fileDescriptor_longrunning_b4ee146853a0b4df, []int{0}
}
func (m *OperationData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperationData.Unmarshal(m, b)
}
func (m *OperationData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperationData.Marshal(b, m, deterministic)
}
func (dst *OperationData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperationData.Merge(dst, src)
}
func (m *OperationData) XXX_Size() int {
	return xxx_messageInfo_OperationData.Size(m)
}
func (m *OperationData) XXX_DiscardUnknown() {
	xxx_messageInfo_OperationData.DiscardUnknown(m)
}

var xxx_messageInfo_OperationData proto.InternalMessageInfo

func (m *OperationData) GetResponseType() string {
	if m != nil {
		return m.ResponseType
	}
	return ""
}

func (m *OperationData) GetMetadataType() string {
	if m != nil {
		return m.MetadataType
	}
	return ""
}

func init() {
	proto.RegisterType((*OperationData)(nil), "google.api.OperationData")
}

func init() {
	proto.RegisterFile("google/api/longrunning.proto", fileDescriptor_longrunning_b4ee146853a0b4df)
}

var fileDescriptor_longrunning_b4ee146853a0b4df = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x31, 0x4b, 0xc5, 0x40,
	0x0c, 0xc7, 0x69, 0x11, 0xc1, 0xc3, 0x8a, 0x74, 0x72, 0x70, 0x10, 0x5d, 0x9c, 0xee, 0x06, 0x47,
	0xa7, 0x16, 0x41, 0x04, 0xc1, 0x22, 0x2e, 0xba, 0x48, 0xd4, 0x10, 0x0e, 0xda, 0x24, 0xdc, 0x9d,
	0x43, 0xbf, 0xce, 0xfb, 0xa4, 0x8f, 0xde, 0xbd, 0xd2, 0xb7, 0x85, 0xff, 0xff, 0x17, 0xf2, 0x8b,
	0xb9, 0x26, 0x11, 0x1a, 0xd1, 0x81, 0x7a, 0x37, 0x0a, 0x53, 0xf8, 0x67, 0xf6, 0x4c, 0x56, 0x83,
	0x24, 0x69, 0x4d, 0x69, 0x2d, 0xa8, 0xbf, 0xfd, 0x34, 0xcd, 0x9b, 0x62, 0x80, 0xe4, 0x85, 0x9f,
	0x20, 0x41, 0x7b, 0x67, 0x9a, 0x80, 0x51, 0x85, 0x23, 0x7e, 0xa7, 0x59, 0xf1, 0xaa, 0xba, 0xa9,
	0xee, 0xcf, 0xde, 0xcf, 0xd7, 0xf0, 0x63, 0x56, 0x5c, 0xa0, 0x09, 0x13, 0xfc, 0x41, 0x82, 0x02,
	0xd5, 0x05, 0x5a, 0xc3, 0x05, 0xea, 0xd9, 0x5c, 0xfc, 0xca, 0x64, 0xb7, 0x63, 0xfd, 0xe5, 0xeb,
	0xe6, 0x32, 0x2c, 0x2a, 0x43, 0xf5, 0xd5, 0x1d, 0x7a, 0x92, 0x11, 0x98, 0xac, 0x04, 0x72, 0x84,
	0x9c, 0x45, 0x5d, 0xa9, 0x40, 0x7d, 0xcc, 0x9f, 0x00, 0xb3, 0xa4, 0x6c, 0x1a, 0x1f, 0x8f, 0xe6,
	0x5d, 0x7d, 0xf2, 0xdc, 0x0d, 0x2f, 0x3f, 0xa7, 0x79, 0xe9, 0x61, 0x1f, 0x00, 0x00, 0xff, 0xff,
	0xc6, 0x08, 0x8f, 0xbd, 0xfd, 0x00, 0x00, 0x00,
}
