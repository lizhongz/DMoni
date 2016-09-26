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

type NDRequest struct {
	AppId string `protobuf:"bytes,1,opt,name=AppId,json=appId" json:"AppId,omitempty"`
	// Start time
	StartTime *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=StartTime,json=startTime" json:"StartTime,omitempty"`
	// End time
	EndTime *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=EndTime,json=endTime" json:"EndTime,omitempty"`
	// output from stdout
	Stdout string `protobuf:"bytes,4,opt,name=Stdout,json=stdout" json:"Stdout,omitempty"`
	// output from stderr
	Stderr string `protobuf:"bytes,5,opt,name=Stderr,json=stderr" json:"Stderr,omitempty"`
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
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x90, 0x41, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0x9b, 0xb6, 0x69, 0xfe, 0x3b, 0x7f, 0x14, 0x1d, 0x44, 0x42, 0x10, 0x2c, 0x39, 0xf5,
	0x94, 0x42, 0xed, 0xc1, 0xab, 0x50, 0xa1, 0x3d, 0x58, 0x24, 0xed, 0xd1, 0xcb, 0x96, 0x4c, 0x6b,
	0xa0, 0xc9, 0xae, 0x9b, 0xc9, 0x21, 0x1f, 0xcf, 0x6f, 0x26, 0xdd, 0x64, 0x2b, 0xe2, 0xc1, 0xe3,
	0xfb, 0xcd, 0x1b, 0xe6, 0xcd, 0x83, 0x8b, 0x42, 0x96, 0xf2, 0x40, 0x26, 0xd1, 0x46, 0xb1, 0xc2,
	0xa0, 0x93, 0xd1, 0xfd, 0x41, 0xa9, 0xc3, 0x91, 0xa6, 0x16, 0xef, 0xea, 0xfd, 0x94, 0xf3, 0x82,
	0x2a, 0x96, 0x85, 0x6e, 0x9d, 0xf1, 0x1b, 0xfc, 0x5b, 0xab, 0x8c, 0x56, 0xe5, 0x5e, 0xe1, 0x25,
	0xf4, 0x57, 0x59, 0xe8, 0x8d, 0xbd, 0x89, 0x48, 0xfb, 0x79, 0x66, 0xb5, 0x0e, 0xfb, 0x9d, 0xd6,
	0x88, 0x30, 0x7c, 0x55, 0x86, 0xc3, 0xc1, 0xd8, 0x9b, 0xf8, 0xe9, 0x50, 0x2b, 0xc3, 0x78, 0x07,
	0x62, 0x49, 0xd2, 0xf0, 0x8e, 0x24, 0x87, 0x43, 0x3b, 0x10, 0xef, 0x0e, 0xc4, 0x9f, 0x1e, 0x88,
	0xf5, 0x22, 0xa5, 0x8f, 0x9a, 0x2a, 0xc6, 0x1b, 0xf0, 0x9f, 0xb4, 0x3e, 0x9f, 0xf0, 0xe5, 0x49,
	0xe0, 0x23, 0x88, 0x0d, 0x4b, 0xc3, 0xdb, 0xbc, 0x20, 0x7b, 0xec, 0xff, 0x2c, 0x4a, 0xda, 0xd8,
	0x89, 0x8b, 0x9d, 0x6c, 0x5d, 0xec, 0x54, 0x54, 0xce, 0x8c, 0x73, 0x08, 0x9e, 0xcb, 0xcc, 0xee,
	0x0d, 0xfe, 0xdc, 0x0b, 0xa8, 0xb5, 0xe2, 0x2d, 0x8c, 0x36, 0x9c, 0xa9, 0xba, 0x8d, 0x2b, 0xd2,
	0x51, 0x65, 0x55, 0xc7, 0xc9, 0x98, 0xd0, 0x3f, 0x73, 0x32, 0x26, 0x16, 0x10, 0x9c, 0x5e, 0xd0,
	0xc7, 0x66, 0xa6, 0x21, 0x78, 0x69, 0x8b, 0xc5, 0x29, 0xf8, 0x1b, 0xd9, 0x2c, 0x73, 0xbc, 0x4e,
	0x5c, 0xf5, 0xae, 0xc7, 0xe8, 0x37, 0x8a, 0x7b, 0x38, 0x07, 0x58, 0x2b, 0xce, 0xf7, 0xcd, 0x42,
	0x95, 0x84, 0xf8, 0x6d, 0x71, 0xf5, 0x44, 0x57, 0x3f, 0x98, 0x3e, 0x36, 0x71, 0x6f, 0x37, 0xb2,
	0x9f, 0x3c, 0x7c, 0x05, 0x00, 0x00, 0xff, 0xff, 0x70, 0x2b, 0xd3, 0xe4, 0xe0, 0x01, 0x00, 0x00,
}
