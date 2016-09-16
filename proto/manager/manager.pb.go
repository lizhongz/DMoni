// Code generated by protoc-gen-go.
// source: manager.proto
// DO NOT EDIT!

/*
Package manager is a generated protocol buffer package.

It is generated from these files:
	manager.proto

It has these top-level messages:
	NodeInfo
	AppIndex
	NDReply
*/
package manager

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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
	// IP address
	Ip string `protobuf:"bytes,2,opt,name=Ip,json=ip" json:"Ip,omitempty"`
	// IP address port
	Port int32 `protobuf:"varint,3,opt,name=Port,json=port" json:"Port,omitempty"`
	// The node's heartbeat value
	Heartbeat int32 `protobuf:"varint,4,opt,name=Heartbeat,json=heartbeat" json:"Heartbeat,omitempty"`
}

func (m *NodeInfo) Reset()                    { *m = NodeInfo{} }
func (m *NodeInfo) String() string            { return proto.CompactTextString(m) }
func (*NodeInfo) ProtoMessage()               {}
func (*NodeInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Application index
type AppIndex struct {
	Id string `protobuf:"bytes,1,opt,name=Id,json=id" json:"Id,omitempty"`
}

func (m *AppIndex) Reset()                    { *m = AppIndex{} }
func (m *AppIndex) String() string            { return proto.CompactTextString(m) }
func (*AppIndex) ProtoMessage()               {}
func (*AppIndex) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// NotifyDone's reply message
type NDReply struct {
}

func (m *NDReply) Reset()                    { *m = NDReply{} }
func (m *NDReply) String() string            { return proto.CompactTextString(m) }
func (*NDReply) ProtoMessage()               {}
func (*NDReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*NodeInfo)(nil), "manager.NodeInfo")
	proto.RegisterType((*AppIndex)(nil), "manager.AppIndex")
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
	NotifyDone(ctx context.Context, in *AppIndex, opts ...grpc.CallOption) (*NDReply, error)
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

func (c *managerClient) NotifyDone(ctx context.Context, in *AppIndex, opts ...grpc.CallOption) (*NDReply, error) {
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
	NotifyDone(context.Context, *AppIndex) (*NDReply, error)
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
	in := new(AppIndex)
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
		return srv.(ManagerServer).NotifyDone(ctx, req.(*AppIndex))
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
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x90, 0xc1, 0x4b, 0x80, 0x30,
	0x14, 0xc6, 0xd5, 0x34, 0xdd, 0x83, 0xa2, 0xde, 0x69, 0x48, 0x07, 0xd9, 0xc9, 0x93, 0x41, 0xfe,
	0x05, 0x81, 0x07, 0x3d, 0x24, 0x61, 0xd7, 0x2e, 0x93, 0xcd, 0x12, 0x6a, 0x6f, 0x8c, 0x1d, 0xf2,
	0xbf, 0x0f, 0xd4, 0x51, 0xd4, 0xf1, 0xf7, 0x8d, 0x6f, 0xdf, 0x8f, 0x07, 0x57, 0x9f, 0xd2, 0xc8,
	0x37, 0xed, 0x1a, 0xeb, 0xc8, 0x13, 0xe6, 0x27, 0x8a, 0x57, 0x28, 0x46, 0x52, 0x7a, 0x30, 0x0b,
	0xe1, 0x35, 0x24, 0x83, 0xe2, 0x71, 0x15, 0xd7, 0x6c, 0x4a, 0x56, 0xb5, 0xb3, 0xe5, 0xc9, 0xc9,
	0x16, 0x11, 0xd2, 0x67, 0x72, 0x9e, 0x5f, 0x54, 0x71, 0x9d, 0x4d, 0xa9, 0x25, 0xe7, 0xf1, 0x0e,
	0x58, 0xaf, 0xa5, 0xf3, 0xb3, 0x96, 0x9e, 0xa7, 0xfb, 0x03, 0x7b, 0x0f, 0x81, 0x28, 0xa1, 0x78,
	0xb4, 0x76, 0x30, 0x4a, 0x7f, 0xfd, 0xfd, 0x5d, 0x30, 0xc8, 0xc7, 0x6e, 0xd2, 0xf6, 0x63, 0x7b,
	0x20, 0xc8, 0x9f, 0x0e, 0x1f, 0xbc, 0x87, 0xec, 0x45, 0x6e, 0xfd, 0x8a, 0xb7, 0x4d, 0x30, 0x0e,
	0x7e, 0xe5, 0xff, 0x48, 0x44, 0xd8, 0x02, 0x8c, 0xe4, 0xd7, 0x65, 0xeb, 0xc8, 0xe8, 0x5f, 0xad,
	0xb0, 0x5b, 0xde, 0xfc, 0xb4, 0x8e, 0x39, 0x11, 0xcd, 0x97, 0xfb, 0x15, 0xda, 0xef, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xc0, 0x59, 0x44, 0xcb, 0x16, 0x01, 0x00, 0x00,
}
