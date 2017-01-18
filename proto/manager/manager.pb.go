// Code generated by protoc-gen-go.
// source: manager.proto
// DO NOT EDIT!

/*
Package manager is a generated protocol buffer package.

It is generated from these files:
	manager.proto

It has these top-level messages:
	NodeInfo
	NDRequest
	NDReply
*/
package manager

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Cluster node inforamtion.
//
// Both manager and agents are treated as a node.
type NodeInfo struct {
	// Id of the node
	Id string `protobuf:"bytes,1,opt,name=Id,json=id" json:"Id,omitempty"`
	// Node's address
	Host string `protobuf:"bytes,2,opt,name=Host,json=host" json:"Host,omitempty"`
	// IP address port
	Port int32 `protobuf:"varint,3,opt,name=Port,json=port" json:"Port,omitempty"`
	// The node's heartbeat value
	Heartbeat int32 `protobuf:"varint,4,opt,name=Heartbeat,json=heartbeat" json:"Heartbeat,omitempty"`
}

func (m *NodeInfo) Reset()                    { *m = NodeInfo{} }
func (m *NodeInfo) String() string            { return proto.CompactTextString(m) }
func (*NodeInfo) ProtoMessage()               {}
func (*NodeInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// NotifyDone's request
type NDRequest struct {
	AppId string `protobuf:"bytes,1,opt,name=AppId,json=appId" json:"AppId,omitempty"`
	// Start time
	StartTime *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=StartTime,json=startTime" json:"StartTime,omitempty"`
	// End time
	EndTime *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=EndTime,json=endTime" json:"EndTime,omitempty"`
}

func (m *NDRequest) Reset()                    { *m = NDRequest{} }
func (m *NDRequest) String() string            { return proto.CompactTextString(m) }
func (*NDRequest) ProtoMessage()               {}
func (*NDRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *NDRequest) GetStartTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *NDRequest) GetEndTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

// NotifyDone's reply message
type NDReply struct {
}

func (m *NDReply) Reset()                    { *m = NDReply{} }
func (m *NDReply) String() string            { return proto.CompactTextString(m) }
func (*NDReply) ProtoMessage()               {}
func (*NDReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*NodeInfo)(nil), "manager.NodeInfo")
	proto.RegisterType((*NDRequest)(nil), "manager.NDRequest")
	proto.RegisterType((*NDReply)(nil), "manager.NDReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Manager service

type ManagerClient interface {
	// Agent pushes its node info to manager, and master responds
	// with its own node info.
	SayHi(ctx context.Context, in *NodeInfo, opts ...grpc.CallOption) (*NodeInfo, error)
	// Agent notifies the manager when an application is finished
	// or the entry/main process does not exists.
	NotifyDone(ctx context.Context, in *NDRequest, opts ...grpc.CallOption) (*NDReply, error)
}

type managerClient struct {
	cc *grpc.ClientConn
}

func NewManagerClient(cc *grpc.ClientConn) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) SayHi(ctx context.Context, in *NodeInfo, opts ...grpc.CallOption) (*NodeInfo, error) {
	out := new(NodeInfo)
	err := grpc.Invoke(ctx, "/manager.Manager/SayHi", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) NotifyDone(ctx context.Context, in *NDRequest, opts ...grpc.CallOption) (*NDReply, error) {
	out := new(NDReply)
	err := grpc.Invoke(ctx, "/manager.Manager/NotifyDone", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Manager service

type ManagerServer interface {
	// Agent pushes its node info to manager, and master responds
	// with its own node info.
	SayHi(context.Context, *NodeInfo) (*NodeInfo, error)
	// Agent notifies the manager when an application is finished
	// or the entry/main process does not exists.
	NotifyDone(context.Context, *NDRequest) (*NDReply, error)
}

func RegisterManagerServer(s *grpc.Server, srv ManagerServer) {
	s.RegisterService(&_Manager_serviceDesc, srv)
}

func _Manager_SayHi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).SayHi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.Manager/SayHi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).SayHi(ctx, req.(*NodeInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_NotifyDone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).NotifyDone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.Manager/NotifyDone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).NotifyDone(ctx, req.(*NDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Manager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "manager.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHi",
			Handler:    _Manager_SayHi_Handler,
		},
		{
			MethodName: "NotifyDone",
			Handler:    _Manager_NotifyDone_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("manager.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x9b, 0x36, 0x31, 0xee, 0x88, 0xa2, 0x83, 0x87, 0x10, 0x04, 0x4b, 0x4e, 0x3d, 0xa5,
	0x50, 0x7b, 0xf0, 0x2a, 0x54, 0x68, 0x0f, 0x16, 0x49, 0xfb, 0x00, 0x6e, 0xc9, 0xb6, 0x0d, 0x24,
	0x99, 0x75, 0x33, 0x3d, 0xe4, 0x31, 0x7c, 0x63, 0xc9, 0xc6, 0x55, 0xc4, 0x83, 0xc7, 0xf9, 0xf6,
	0x1f, 0xe6, 0xdb, 0x1f, 0x2e, 0x2b, 0x59, 0xcb, 0x83, 0x32, 0xa9, 0x36, 0xc4, 0x84, 0xe1, 0xd7,
	0x18, 0xdf, 0x1f, 0x88, 0x0e, 0xa5, 0x9a, 0x5a, 0xbc, 0x3b, 0xed, 0xa7, 0x5c, 0x54, 0xaa, 0x61,
	0x59, 0xe9, 0x3e, 0x99, 0xbc, 0xc1, 0xf9, 0x9a, 0x72, 0xb5, 0xaa, 0xf7, 0x84, 0x57, 0x30, 0x5c,
	0xe5, 0x91, 0x37, 0xf6, 0x26, 0x22, 0x1b, 0x16, 0x39, 0x22, 0xf8, 0x4b, 0x6a, 0x38, 0x1a, 0x5a,
	0xe2, 0x1f, 0xa9, 0xe1, 0x8e, 0xbd, 0x92, 0xe1, 0x68, 0x34, 0xf6, 0x26, 0x41, 0xe6, 0x6b, 0x32,
	0x8c, 0x77, 0x20, 0x96, 0x4a, 0x1a, 0xde, 0x29, 0xc9, 0x91, 0x6f, 0x1f, 0xc4, 0xd1, 0x81, 0xe4,
	0xc3, 0x03, 0xb1, 0x5e, 0x64, 0xea, 0xfd, 0xa4, 0x1a, 0xc6, 0x5b, 0x08, 0x9e, 0xb4, 0xfe, 0x3e,
	0x13, 0xc8, 0x6e, 0xc0, 0x47, 0x10, 0x1b, 0x96, 0x86, 0xb7, 0x45, 0xa5, 0xec, 0xb9, 0x8b, 0x59,
	0x9c, 0xf6, 0xea, 0xa9, 0x53, 0x4f, 0xb7, 0x4e, 0x3d, 0x13, 0x8d, 0x0b, 0xe3, 0x1c, 0xc2, 0xe7,
	0x3a, 0xb7, 0x7b, 0xa3, 0x7f, 0xf7, 0x42, 0xd5, 0x47, 0x13, 0x01, 0x61, 0xa7, 0xa4, 0xcb, 0x76,
	0xa6, 0x21, 0x7c, 0xe9, 0xcb, 0xc2, 0x29, 0x04, 0x1b, 0xd9, 0x2e, 0x0b, 0xbc, 0x49, 0x5d, 0x9d,
	0xae, 0x9b, 0xf8, 0x2f, 0x4a, 0x06, 0x38, 0x07, 0x58, 0x13, 0x17, 0xfb, 0x76, 0x41, 0xb5, 0x42,
	0xfc, 0x89, 0xb8, 0xef, 0xc6, 0xd7, 0xbf, 0x98, 0x2e, 0xdb, 0x64, 0xb0, 0x3b, 0xb3, 0x66, 0x0f,
	0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x75, 0xfd, 0xf7, 0x25, 0xb4, 0x01, 0x00, 0x00,
}
