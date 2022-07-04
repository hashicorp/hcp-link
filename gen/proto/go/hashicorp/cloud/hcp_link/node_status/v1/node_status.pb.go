// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hashicorp/cloud/hcp_link/node_status/v1/node_status.proto

package node_statusv1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

// GetNodeStatusRequest is empty for now as GetNodeStatus does not expect any
// arguments.
type GetNodeStatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNodeStatusRequest) Reset()         { *m = GetNodeStatusRequest{} }
func (m *GetNodeStatusRequest) String() string { return proto.CompactTextString(m) }
func (*GetNodeStatusRequest) ProtoMessage()    {}
func (*GetNodeStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b45ec5d664cd63e5, []int{0}
}

func (m *GetNodeStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNodeStatusRequest.Unmarshal(m, b)
}
func (m *GetNodeStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNodeStatusRequest.Marshal(b, m, deterministic)
}
func (m *GetNodeStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNodeStatusRequest.Merge(m, src)
}
func (m *GetNodeStatusRequest) XXX_Size() int {
	return xxx_messageInfo_GetNodeStatusRequest.Size(m)
}
func (m *GetNodeStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNodeStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetNodeStatusRequest proto.InternalMessageInfo

// GetNodeStatusResponse contains information about the node's current status
// the status is a combination of common status information and product specific
// status information.
type GetNodeStatusResponse struct {
	// node_id is the ID assigned to the node. It is expected to unique within the
	// link resource (e.g. within the cluster).
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// node_version is the node's version in semantic version format.
	NodeVersion string `protobuf:"bytes,2,opt,name=node_version,json=nodeVersion,proto3" json:"node_version,omitempty"`
	// node_os is the lower-case name of the operating system the client is
	// running on (e.g. linux, windows).
	NodeOs string `protobuf:"bytes,3,opt,name=node_os,json=nodeOs,proto3" json:"node_os,omitempty"`
	// node_architecture is the lower-case architecture of the client binary
	// (e.g. amd64, arm, ...).
	NodeArchitecture string `protobuf:"bytes,4,opt,name=node_architecture,json=nodeArchitecture,proto3" json:"node_architecture,omitempty"`
	// timestamp is the time the status was recorded on the node.
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// status_version is the version of the status message format. To ensure
	// that the version is not omitted by accident the initial version is 1.
	StatusVersion uint32 `protobuf:"varint,6,opt,name=status_version,json=statusVersion,proto3" json:"status_version,omitempty"`
	// status is the product specific status of the node. The link library and
	// service is agnostic to the information transmitted in this field.
	Status               *anypb.Any `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetNodeStatusResponse) Reset()         { *m = GetNodeStatusResponse{} }
func (m *GetNodeStatusResponse) String() string { return proto.CompactTextString(m) }
func (*GetNodeStatusResponse) ProtoMessage()    {}
func (*GetNodeStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b45ec5d664cd63e5, []int{1}
}

func (m *GetNodeStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNodeStatusResponse.Unmarshal(m, b)
}
func (m *GetNodeStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNodeStatusResponse.Marshal(b, m, deterministic)
}
func (m *GetNodeStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNodeStatusResponse.Merge(m, src)
}
func (m *GetNodeStatusResponse) XXX_Size() int {
	return xxx_messageInfo_GetNodeStatusResponse.Size(m)
}
func (m *GetNodeStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNodeStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetNodeStatusResponse proto.InternalMessageInfo

func (m *GetNodeStatusResponse) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *GetNodeStatusResponse) GetNodeVersion() string {
	if m != nil {
		return m.NodeVersion
	}
	return ""
}

func (m *GetNodeStatusResponse) GetNodeOs() string {
	if m != nil {
		return m.NodeOs
	}
	return ""
}

func (m *GetNodeStatusResponse) GetNodeArchitecture() string {
	if m != nil {
		return m.NodeArchitecture
	}
	return ""
}

func (m *GetNodeStatusResponse) GetTimestamp() *timestamppb.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *GetNodeStatusResponse) GetStatusVersion() uint32 {
	if m != nil {
		return m.StatusVersion
	}
	return 0
}

func (m *GetNodeStatusResponse) GetStatus() *anypb.Any {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*GetNodeStatusRequest)(nil), "hashicorp.cloud.hcp_link.node_status.v1.GetNodeStatusRequest")
	proto.RegisterType((*GetNodeStatusResponse)(nil), "hashicorp.cloud.hcp_link.node_status.v1.GetNodeStatusResponse")
}

func init() {
	proto.RegisterFile("hashicorp/cloud/hcp_link/node_status/v1/node_status.proto", fileDescriptor_b45ec5d664cd63e5)
}

var fileDescriptor_b45ec5d664cd63e5 = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xcd, 0x6e, 0xd3, 0x30,
	0x1c, 0x57, 0xc2, 0xe8, 0x34, 0x8f, 0x02, 0x8b, 0x06, 0x84, 0x5e, 0x28, 0x93, 0x26, 0x8a, 0x06,
	0xb6, 0x32, 0x2e, 0x60, 0x04, 0xd2, 0xba, 0xc3, 0x8a, 0x04, 0xa3, 0xca, 0x50, 0x0f, 0x53, 0xa4,
	0xc9, 0x75, 0x4c, 0x62, 0xd1, 0xc6, 0x21, 0x76, 0x22, 0xed, 0x25, 0x78, 0x08, 0xc4, 0x01, 0xf1,
	0x28, 0x5c, 0x79, 0x03, 0x8e, 0x3c, 0x05, 0x8a, 0xdd, 0xc6, 0xa1, 0x70, 0xa8, 0x38, 0xfa, 0xff,
	0xfb, 0xf0, 0xef, 0xff, 0x01, 0x9e, 0xa5, 0x44, 0xa6, 0x9c, 0x8a, 0x22, 0x47, 0x74, 0x26, 0xca,
	0x18, 0xa5, 0x34, 0xbf, 0x98, 0xf1, 0xec, 0x03, 0xca, 0x44, 0xcc, 0x2e, 0xa4, 0x22, 0xaa, 0x94,
	0xa8, 0x0a, 0xda, 0x4f, 0x98, 0x17, 0x42, 0x09, 0xef, 0x41, 0x23, 0x85, 0x5a, 0x0a, 0x97, 0x52,
	0xd8, 0xe6, 0x56, 0x41, 0xef, 0x6e, 0x22, 0x44, 0x32, 0x63, 0x48, 0xcb, 0xa6, 0xe5, 0x7b, 0x44,
	0xb2, 0x4b, 0xe3, 0xd1, 0xbb, 0xb7, 0x0a, 0x29, 0x3e, 0x67, 0x52, 0x91, 0x79, 0x6e, 0x08, 0x7b,
	0xb7, 0xc1, 0xee, 0x09, 0x53, 0xa7, 0x22, 0x66, 0x67, 0xda, 0x2f, 0x64, 0x1f, 0x4b, 0x26, 0xd5,
	0xde, 0x57, 0x17, 0xdc, 0x5a, 0x01, 0x64, 0x2e, 0x32, 0xc9, 0xbc, 0x3b, 0x60, 0x53, 0xff, 0xcf,
	0x63, 0xdf, 0xe9, 0x3b, 0x83, 0xad, 0xb0, 0x53, 0x3f, 0x5f, 0xc5, 0xde, 0x7d, 0x70, 0x4d, 0x03,
	0x15, 0x2b, 0x24, 0x17, 0x99, 0xef, 0x6a, 0x74, 0xbb, 0xae, 0x4d, 0x4c, 0xa9, 0xd1, 0x0a, 0xe9,
	0x5f, 0xb1, 0xda, 0xb7, 0xd2, 0x3b, 0x00, 0x3b, 0x1a, 0x20, 0x05, 0x4d, 0xb9, 0x62, 0x54, 0x95,
	0x05, 0xf3, 0x37, 0x34, 0xe5, 0x66, 0x0d, 0x1c, 0xb5, 0xea, 0xde, 0x53, 0xb0, 0xd5, 0xb4, 0xe1,
	0x5f, 0xed, 0x3b, 0x83, 0xed, 0xc3, 0x1e, 0x34, 0x8d, 0xc2, 0x65, 0xa3, 0xf0, 0xdd, 0x92, 0x11,
	0x5a, 0xb2, 0xb7, 0x0f, 0xae, 0x9b, 0xb1, 0x35, 0x21, 0x3b, 0x7d, 0x67, 0xd0, 0x0d, 0xbb, 0xa6,
	0xba, 0x8c, 0xf9, 0x08, 0x74, 0x4c, 0xc1, 0xdf, 0xd4, 0xee, 0xbb, 0x7f, 0xb9, 0x1f, 0x65, 0x97,
	0xe1, 0x82, 0x73, 0xf8, 0xc5, 0x01, 0x3b, 0x76, 0x4e, 0x67, 0xac, 0xa8, 0x38, 0x65, 0xde, 0x27,
	0x07, 0x74, 0xff, 0x18, 0xa0, 0xf7, 0x02, 0xae, 0xb9, 0x50, 0xf8, 0xaf, 0x8d, 0xf4, 0x5e, 0xfe,
	0xaf, 0xdc, 0xec, 0x6d, 0xf8, 0xc3, 0x05, 0x07, 0x54, 0xcc, 0xd7, 0x75, 0x19, 0xde, 0xb0, 0x1e,
	0xe3, 0xba, 0xed, 0xb1, 0x73, 0x7e, 0x9e, 0x70, 0x95, 0x96, 0x53, 0x48, 0xc5, 0x1c, 0xd9, 0xbb,
	0x4e, 0x69, 0xfe, 0x58, 0x5f, 0x74, 0xc2, 0x32, 0x94, 0x08, 0xb4, 0xe6, 0xc5, 0x3f, 0x6f, 0x3d,
	0xab, 0xe0, 0xb3, 0xbb, 0x31, 0x3a, 0x1e, 0x9d, 0x7e, 0x73, 0xf7, 0x47, 0x4d, 0xbe, 0x63, 0x9d,
	0x6f, 0x44, 0xf3, 0xd7, 0x75, 0x3c, 0x1b, 0x07, 0x4e, 0x82, 0xef, 0x2d, 0x5e, 0xa4, 0x79, 0xd1,
	0x82, 0x17, 0x59, 0x5e, 0x34, 0x09, 0x7e, 0xba, 0xc1, 0x5a, 0xbc, 0xe8, 0x64, 0x3c, 0x7c, 0xc3,
	0x14, 0x89, 0x89, 0x22, 0xbf, 0xdc, 0x87, 0x8d, 0x06, 0x63, 0x2d, 0xc2, 0x78, 0xa1, 0xc2, 0xd8,
	0xca, 0x30, 0x9e, 0x04, 0xd3, 0x8e, 0x3e, 0x89, 0x27, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xbd,
	0x2c, 0x70, 0x6c, 0xe8, 0x03, 0x00, 0x00,
}
